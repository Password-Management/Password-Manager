package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"log"
)

// GenerateRSAKeys generates a pair of RSA keys (private and public)
func GenerateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

// PrivateKeyToPEM converts a private key to PEM format
func PrivateKeyToPEM(privateKey *rsa.PrivateKey) string {
	// Convert the private key to ASN.1 DER encoded form
	privateKeyDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// Create a PEM block with the private key
	pemBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyDER,
	}

	// Encode the PEM block to a string
	privateKeyPEM := pem.EncodeToMemory(pemBlock)
	return string(privateKeyPEM)
}

// PublicKeyToPEM converts a public key to PEM format
func PublicKeyToPEM(publicKey *rsa.PublicKey) (string, error) {
	// Marshal the public key to PKIX, ASN.1 DER form
	pubASN1, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}

	// Create a PEM block with the public key
	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	return string(pubPEM), nil
}

// EncryptPassword encrypts the password using the provided RSA public key
func EncryptPassword(publicKey *rsa.PublicKey, password string) (string, error) {
	// Encrypt the password using PKCS1v15
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(password))
	if err != nil {
		return "", err
	}

	// Convert the encrypted bytes to a Base64-encoded string for storage
	encodedString := base64.StdEncoding.EncodeToString(encryptedBytes)
	return encodedString, nil
}

// DecryptPassword decrypts the encrypted password using the RSA private key
func DecryptPassword(privateKey *rsa.PrivateKey, encryptedPassword string) (string, error) {
	// First, decode the Base64 encoded string to get the encrypted bytes
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedPassword)
	if err != nil {
		log.Println("Failed to base64 decode encrypted password:", err)
		return "", err
	}

	// Decrypt the password using PKCS1v15
	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedBytes)
	if err != nil {
		log.Println("Failed to decrypt password:", err)
		return "", err
	}

	// Return the decrypted password as a string
	return string(decryptedBytes), nil
}

// PemToPublicKey converts a PEM encoded public key string back to an RSA public key
func PemToPublicKey(pemEncodedPublicKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemEncodedPublicKey))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not RSA public key")
	}

	return publicKey, nil
}

// PemToPrivateKey converts a PEM encoded private key string back to an RSA private key
func PemToPrivateKey(pemEncodedPrivateKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemEncodedPrivateKey))
	if block == nil || block.Type != "PRIVATE KEY" {
		log.Println("Failed to decode PEM block containing private key")
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println("Failed to parse private key:", err)
		return nil, err
	}

	return privateKey, nil
}
