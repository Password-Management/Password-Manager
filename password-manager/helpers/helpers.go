package helpers

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"os"
	"password-manager/models"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var customerEmail string
var smtpHost string
var smtpPort string

func Getenv() error {
	err := godotenv.Load()
	if err != nil {
		log.Print("error loading .env file" + err.Error())
		return err
	}
	return nil
}

func getEnvValue() error {
	err := Getenv()
	if err != nil {
		return err
	}
	config, err := ReadConfig("/app/config.yml")
	if err != nil {
		return errors.New("error while reading the config: " + err.Error())
	}
	customerEmail = config.Email
	smtpHost = os.Getenv("SMTP_HOST")
	smtpPort = os.Getenv("SMTP_PORT")
	return nil
}

func SendEmail(body string, subject string) error {
	viper.AutomaticEnv()
	err := getEnvValue()
	if err != nil {
		return errors.New("error while setting global values" + err.Error())
	}
	fromEmail := "passwordmanager@noreply.com"
	err = sendEmail(fromEmail, customerEmail, body, subject)
	if err != nil {
		fmt.Print("Error sending email: ", err)
		return errors.New("error while sending email: " + err.Error())
	}
	return nil
}

func sendEmail(from, to string, body string, subject string) error {
	addr := smtpHost + ":" + smtpPort
	fmt.Println("The address: ", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to connect to MailHog: %v", err)
	}
	// Create an unencrypted client
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %v", err)
	}

	if err := client.Mail(from); err != nil {
		return fmt.Errorf("failed to set sender: %v", err)
	}

	if err := client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set recipient: %v", err)
	}

	// Send the email body
	wc, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to send email data: %v", err)
	}

	message := []byte(
		"From: " + from + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" + // Add Subject here
			"\r\n" +
			body + "\n" +
			"Please do not reply to this email.\n",
	)

	_, err = wc.Write(message)
	if err != nil {
		return fmt.Errorf("failed to write email message: %v", err)
	}
	err = wc.Close()
	if err != nil {
		return fmt.Errorf("failed to close email writer: %v", err)
	}
	client.Quit()
	return nil
}

func ReadConfig(filename string) (config *models.Config, err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}
	return config, nil
}

func CreateConfig(product *models.Config) (string, error) {
	yamlData, err := yaml.Marshal(&product)
	if err != nil {
		log.Fatalf("Error marshaling to YAML: %v", err)
		return "", nil
	}

	// Write the YAML to a file
	err = os.WriteFile("config.yml", yamlData, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
		return "", err
	}

	return "YAML configuration written to product.yaml", nil
}
