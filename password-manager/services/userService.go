package services

import (
	"errors"
	"log"
	dallayer "password-manager/dalLayer"
	"password-manager/encryption"
	"password-manager/models"

	"github.com/google/uuid"
)

func UserServiceRequest() (UserService, error) {
	return &UserServiceImpl{}, nil
}

type UserServiceImpl struct {
	UserRepo     dallayer.User
	PasswordRepo dallayer.Password
}

type UserService interface {
	CreateWebsiteEntry(value *models.CreatePasswordRequest) (resp *models.SuccessResponse, err error)
	GetPassword(value *models.GetPasswordRequest) (resp *models.SuccessResponse, err error)
	ListWebsites(value *models.ListWebsiteRequest) (resp []*models.SuccessResponse, err error)
}

func (us *UserServiceImpl) SetupDalLayer() error {
	var err error
	us.UserRepo, err = dallayer.NewUserDalRequest()
	if err != nil {
		return errors.New("error while connecting to the user Dal layer: " + err.Error())
	}

	us.PasswordRepo, err = dallayer.NewPasswordDalRequest()
	if err != nil {
		return errors.New("error while connecting to the master Dal layer: " + err.Error())
	}

	return nil
}

func (us *UserServiceImpl) CreateWebsiteEntry(value *models.CreatePasswordRequest) (resp *models.SuccessResponse, err error) {
	err = us.SetupDalLayer()
	if err != nil {
		return nil, errors.New("error while setting up the dal connection")
	}
	userInfo, err := us.UserRepo.FindById(uuid.MustParse(value.UserId), uuid.MustParse(value.MasterId))
	if err != nil {
		log.Println("the error = ", err)
		return nil, errors.New("error while finding the user information: " + err.Error())
	}
	publicKey, err := encryption.PemToPublicKey(userInfo.PublicKey)
	if err != nil {
		return nil, errors.New("error while convertng pem to public key: " + err.Error())
	}
	encryptedPassword, err := encryption.EncryptPassword(publicKey, value.Password)
	if err != nil {
		return nil, errors.New("error while encrypting your password: " + err.Error())
	}
	err = us.PasswordRepo.Create(&models.DbPassword{
		UserId:      uuid.MustParse(value.UserId),
		WebisteName: value.WebisteName,
		Password:    encryptedPassword,
	})
	return &models.SuccessResponse{
		Message: "Entry for webiste " + value.WebisteName + " is added successfully",
	}, nil
}

func (us *UserServiceImpl) GetPassword(value *models.GetPasswordRequest) (resp *models.SuccessResponse, err error) {
	err = us.SetupDalLayer()
	if err != nil {
		return nil, errors.New("error while setting up the dal connection")
	}
	userInfo, err := us.UserRepo.FindById(uuid.MustParse(value.UserId), uuid.MustParse(value.MasterId))
	if err != nil {
		log.Println("the error = ", err)
		return nil, errors.New("error while finding the user information: " + err.Error())
	}
	privateKey, err := encryption.PemToPrivateKey(userInfo.PrivateKey)
	if err != nil {
		return nil, errors.New("error while convertng pem to private key: " + err.Error())
	}
	PasswordInfo, err := us.PasswordRepo.FindWebsiteName(value.WebisteName, uuid.MustParse(value.UserId))
	if err != nil {
		return nil, errors.New("error while getting the password information as per userinfo: " + err.Error())
	}
	log.Println("the private key: ", userInfo.PrivateKey)
	decryptedPassword, err := encryption.DecryptPassword(privateKey, PasswordInfo.Password)
	if err != nil {
		return nil, errors.New("error while decrypting the password: " + err.Error())
	}
	return &models.SuccessResponse{
		Message: "The password for your webiste " + value.WebisteName + " is " + decryptedPassword,
	}, nil
}

func (us *UserServiceImpl) ListWebsites(value *models.ListWebsiteRequest) (resp []*models.SuccessResponse, err error) {
	err = us.SetupDalLayer()
	if err != nil {
		return nil, errors.New("error while setting up the dal connection")
	}
	PasswordInfo, err := us.PasswordRepo.FindAll(uuid.MustParse(value.UserId))
	if err != nil {
		return nil, errors.New("error while fetching the password info: "+ err.Error())
	}
	for _, info := range PasswordInfo {
		var dummyresponse models.SuccessResponse
		dummyresponse.Message = info.WebisteName
		resp = append(resp, &dummyresponse)
	}
	return resp, nil
}