package services

import (
	"errors"
	"os"
	dallayer "password-manager/dalLayer"
	"password-manager/helpers"
	"password-manager/models"

	"github.com/google/uuid"
)

func LoginServiceRequest() (LoginService, error) {
	return &LoginServiceImpl{}, nil
}

type LoginServiceImpl struct {
	MasterRepo dallayer.Master
	UserRepo   dallayer.User
	LoginRepo  dallayer.Login
}

type LoginService interface {
	LoginMaster(value *models.MasterLoginRequest) (string, error)
	LoginUser(value *models.UserLoginRequest) (string, error)
	Logout(userID uuid.UUID) (string, error)
}

func (ls *LoginServiceImpl) SetupDalLayer() error {
	var err error
	ls.MasterRepo, err = dallayer.NewMasterDalRequest()
	if err != nil {
		return errors.New("error while connecting to the master Dal layer from service: " + err.Error())
	}
	ls.UserRepo, err = dallayer.NewUserDalRequest()
	if err != nil {
		return errors.New("error while connecting to the user Dal layer from service: " + err.Error())
	}
	ls.LoginRepo, err = dallayer.NewLoginDalRequest()
	if err != nil {
		return errors.New("error while connecting to the login Dal layer from service: " + err.Error())
	}
	return nil
}

func (ls *LoginServiceImpl) LoginMaster(value *models.MasterLoginRequest) (string, error) {
	err := ls.SetupDalLayer()
	if err != nil {
		return "", err
	}
	masterDetails, err := ls.MasterRepo.FindAll()
	if err != nil {
		return "", nil
	}
	if masterDetails[0].SpecialKey != value.SpecialKey {
		return "", errors.New("The special key provided was not found or is incorrect.")
	}
	err = ls.LoginRepo.Create(&models.DBLogin{
		UserId:   masterDetails[0].Id,
		IsLogin:  true,
		IsMaster: true,
	})
	if err != nil {
		return "", errors.New("error while creating a login entry in the database: " + err.Error())
	}
	return "Login of master successfull for user " + masterDetails[0].Name + " .", nil
}

func (ls *LoginServiceImpl) LoginUser(value *models.UserLoginRequest) (string, error) {
	err := ls.SetupDalLayer()
	if err != nil {
		return "", err
	}
	err = helpers.Getenv()
	if err != nil {
		return "", errors.New("error while connecting to env file: " + err.Error())
	}
	config := os.Getenv("ALGORITHM")
	var Name string
	var userID uuid.UUID
	if config == "RSA" {
		userDetails, userErr := ls.UserRepo.FindByRSA(&models.DBRSAUser{
			Email: value.Email,
		})
		if userErr != nil {
			return "", errors.New("error while finding user from database: " + userErr.Error())
		}
		if userDetails.Password != value.Password {
			return "", errors.New("Password is incorrect")
		}
		Name = userDetails.Name
		userID = userDetails.Id
	} else if config  == "ASA" {
		userDetails, userErr := ls.UserRepo.FindByASA(&models.DBASAUser{
			Email: value.Email,
		})
		if userErr != nil {
			return "", errors.New("error while finding user from database: " + userErr.Error())
		}
		if userDetails.Password != value.Password {
			return "", errors.New("Password is incorrect")
		}
		Name = userDetails.Name
		userID = userDetails.Id
	}
	err = ls.LoginRepo.Create(&models.DBLogin{
		UserId:   userID,
		IsLogin:  true,
		IsMaster: false,
	})
	if err != nil {
		return "", errors.New("error while creating a login entry in the database: " + err.Error())
	}
	return "User logged in successfully: " + Name, nil
}

func (ls *LoginServiceImpl) Logout(userID uuid.UUID) (string, error) {
	err := ls.SetupDalLayer()
	if err != nil {
		return "", err
	}
	err = ls.LoginRepo.Logout(userID)
	if err != nil {
		return "", err
	}
	return "User has logged out successfully.", nil
}