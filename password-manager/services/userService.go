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
	MasterRepo   dallayer.Master
}

type UserService interface {
	CreateWebsiteEntry(value *models.CreatePasswordRequest, userId string, masterId string) (resp *models.SuccessResponse, err error)
	GetPassword(value *models.GetPasswordRequest, userId string, masterId string) (resp *models.SuccessResponse, err error)
	ListWebsites(userId string, masterId string) (resp []*models.ListWebsiteResponse, err error)
	GetUserInfo(userId string, masterId string) (resp *models.GetUserInfoResponse, err error)
	DeletePassword(websiteName string, masterId string, userId string) (*models.DeleteWebsiteResponse, error)
	UpdatePassKey(value *models.UserPassKeyUpdateRequest, userId string, masterId string) (response *models.SuccessResponse, err error)
	VerifySpecialKey(userId string, masterId string, specialKey string) (response *models.SuccessResponse, err error)
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

	us.MasterRepo, err = dallayer.NewMasterDalRequest()
	if err != nil {
		return errors.New("error while connecting to the master Dal layer from service: " + err.Error())
	}

	return nil
}

func (us *UserServiceImpl) CreateWebsiteEntry(value *models.CreatePasswordRequest, userId string, masterId string) (resp *models.SuccessResponse, err error) {
	err = us.SetupDalLayer()
	if err != nil {
		return nil, errors.New("error while setting up the dal connection")
	}
	userInfo, err := us.UserRepo.FindById(uuid.MustParse(userId), uuid.MustParse(masterId))
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
		UserId:      uuid.MustParse(userId),
		WebisteName: value.WebisteName,
		UserName:    value.UserName,
		Password:    encryptedPassword,
	})
	if err != nil {
		return nil, errors.New("error while creating the password entry: " + err.Error())
	}
	return &models.SuccessResponse{
		Message: "Entry for webiste " + value.WebisteName + " is added successfully",
	}, nil
}

func (us *UserServiceImpl) GetPassword(value *models.GetPasswordRequest, userId string, masterId string) (resp *models.SuccessResponse, err error) {
	err = us.SetupDalLayer()
	if err != nil {
		return nil, errors.New("error while setting up the dal connection")
	}
	userInfo, err := us.UserRepo.FindById(uuid.MustParse(userId), uuid.MustParse(masterId))
	if err != nil {
		log.Println("the error = ", err)
		return nil, errors.New("error while finding the user information: " + err.Error())
	}
	privateKey, err := encryption.PemToPrivateKey(userInfo.PrivateKey)
	if err != nil {
		return nil, errors.New("error while convertng pem to private key: " + err.Error())
	}
	PasswordInfo, err := us.PasswordRepo.FindWebsiteName(value.WebisteName, uuid.MustParse(userId))
	if err != nil {
		return nil, errors.New("error while getting the password information as per userinfo: " + err.Error())
	}
	log.Println(userInfo.PrivateKey)
	log.Println("the password info = ", PasswordInfo)
	decryptedPassword, err := encryption.DecryptPassword(privateKey, PasswordInfo.Password)
	if err != nil {
		return nil, errors.New("error while decrypting the password: " + err.Error())
	}
	log.Println("the decrypt password = ", decryptedPassword)
	return &models.SuccessResponse{
		Message: decryptedPassword,
	}, nil
}

func (us *UserServiceImpl) ListWebsites(userId string, masterId string) (resp []*models.ListWebsiteResponse, err error) {
	err = us.SetupDalLayer()
	if err != nil {
		return nil, errors.New("error while setting up the dal connection")
	}

	PasswordInfo, err := us.PasswordRepo.FindAll(uuid.MustParse(userId))
	if err != nil {
		return nil, errors.New("error while fetching the password info: " + err.Error())
	}
	for _, info := range PasswordInfo {
		var dummyresponse models.ListWebsiteResponse
		dummyresponse.WebsiteName = info.WebisteName
		dummyresponse.UserName = info.UserName
		resp = append(resp, &dummyresponse)
	}
	return resp, nil
}

func (us *UserServiceImpl) GetUserInfo(userId string, masterId string) (resp *models.GetUserInfoResponse, err error) {
	log.Println("From the service level = ", userId)
	err = us.SetupDalLayer()
	if err != nil {
		return nil, errors.New("error while setting up the dal connection")
	}
	userInfo, err := us.UserRepo.FindById(uuid.MustParse(userId), uuid.MustParse(masterId))
	if err != nil {
		log.Println("the error = ", err)
		return nil, errors.New("error while finding the user information: " + err.Error())
	}
	passwordDeails, err := us.PasswordRepo.FindAll(uuid.MustParse(userId))
	if err != nil {
		return nil, errors.New("error while finding the users website: " + err.Error())
	}
	var webisites []string
	for _, i := range passwordDeails {
		webisites = append(webisites, i.WebisteName)
	}
	return &models.GetUserInfoResponse{
		Email:        userInfo.Email,
		Name:         userInfo.Name,
		WebsiteNames: webisites,
	}, nil
}

func (us *UserServiceImpl) DeletePassword(websiteName string, masterId string, userId string) (*models.DeleteWebsiteResponse, error) {
	err := us.SetupDalLayer()
	if err != nil {
		return nil, errors.New("error while setting up the dal connection for deleting the website entry")
	}
	// Getting the algorithm currently on the system
	masterInfo, err := us.MasterRepo.FindBy(&models.DBMaster{
		Id: uuid.MustParse(masterId),
	})
	if err != nil {
		return nil, err
	}
	algo := masterInfo.Algorithm
	flag := false
	if algo == "RSA" {
		_, err := us.UserRepo.FindByRSA(&models.DBRSAUser{
			UserId: uuid.MustParse(userId),
		})
		if err != nil {
			return nil, err
		}
		flag = true
	} else {
		_, err := us.UserRepo.FindByASA(&models.DBASAUser{
			UserId: uuid.MustParse(userId),
		})
		if err != nil {
			return nil, err
		}
		flag = true
	}
	if !flag {
		return nil, errors.New("user not found")
	}
	//Find if the website exist and delete it
	websiteDetails, err := us.PasswordRepo.FindWebsiteName(websiteName, uuid.MustParse(userId))
	if err != nil {
		return nil, errors.New("website not found: " + err.Error())
	}
	err = us.PasswordRepo.DeletePassword(websiteDetails.WebisteName)
	if err != nil {
		return nil, errors.New("unable to delete the password entry for this website: " + err.Error())
	}
	return &models.DeleteWebsiteResponse{
		Response: "Entry for websiteName " + websiteName + " deleted successfully",
	}, nil
}

func (us *UserServiceImpl) UpdatePassKey(value *models.UserPassKeyUpdateRequest, userId string, masterId string) (response *models.SuccessResponse, err error) {
	err = us.SetupDalLayer()
	if err != nil {
		return nil, errors.New("error while setting up the dal connection for deleting the website entry")
	}
	log.Println("the request Body: ", value)
	masterInfo, err := us.MasterRepo.FindBy(&models.DBMaster{
		Id: uuid.MustParse(masterId),
	})
	if err != nil {
		return nil, err
	}
	algo := masterInfo.Algorithm
	var message string
	if algo == "RSA" {
		userDetails, err := us.UserRepo.FindByRSA(&models.DBRSAUser{
			UserId: uuid.MustParse(userId),
		})
		if err != nil {
			return nil, errors.New("error while finding the userDetails: " + err.Error())
		}
		if value.Type == "key" {
			userDetails.SpecialKey = value.Value
			message = "Special Key is updated for RSA type user."
		} else if value.Type == "pass" {
			userDetails.Password = value.Value
			message = "Password is updated for RSA type user."
		}
	} else {
		userDetails, err := us.UserRepo.FindByASA(&models.DBASAUser{
			UserId: uuid.MustParse(userId),
		})
		if err != nil {
			return nil, errors.New("error while finding the userDetails: " + err.Error())
		}
		if value.Type == "key" {
			userDetails.SpecialKey = value.Value
			message = "Special Key is updated for ASA type user."
		} else if value.Type == "pass" {
			userDetails.Password = value.Value
			message = "Password is updated for ASA type user."
		}
	}
	return &models.SuccessResponse{
		Message: message,
	}, nil
}

func (us *UserServiceImpl) VerifySpecialKey(userId string, masterId string, specialKey string) (response *models.SuccessResponse, err error) {
	err = us.SetupDalLayer()
	if err != nil {
		return nil, errors.New("error while setting up the dal connection for checking the website Key")
	}
	masterInfo, err := us.MasterRepo.FindBy(&models.DBMaster{
		Id: uuid.MustParse(masterId),
	})
	if err != nil {
		return nil, err
	}
	algo := masterInfo.Algorithm
	if algo == "RSA" {
		userDetails, err := us.UserRepo.FindByRSA(&models.DBRSAUser{
			UserId: uuid.MustParse(userId),
		})
		if err != nil {
			return nil, errors.New("error while finding the userDetails: " + err.Error())
		}
		if userDetails.SpecialKey == specialKey {
			return &models.SuccessResponse{
				Message: "Success",
			}, nil
		}
	} else {
		userDetails, err := us.UserRepo.FindByASA(&models.DBASAUser{
			UserId: uuid.MustParse(userId),
		})
		if err != nil {
			return nil, errors.New("error while finding the userDetails: " + err.Error())
		}
		if userDetails.SpecialKey == specialKey {
			return &models.SuccessResponse{
				Message: "Success",
			}, nil
		}
	}
	return &models.SuccessResponse{
		Message: "Failure",
	}, nil
}
