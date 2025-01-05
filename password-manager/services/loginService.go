package services

import (
	"errors"
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
	LoginMaster(value *models.MasterLoginRequest) (*models.LoginResponseMaster, error)
	LoginUser(value *models.UserLoginRequest) (*models.LoginResponse, error)
	Logout(userID uuid.UUID) (*models.SuccessResponse, error)
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

func (ls *LoginServiceImpl) LoginMaster(value *models.MasterLoginRequest) (*models.LoginResponseMaster, error) {
	err := ls.SetupDalLayer()
	if err != nil {
		return nil, err
	}
	masterDetails, err := ls.MasterRepo.FindAll()
	if err != nil {
		return nil, errors.New("error while finding the master deatils: " + err.Error())
	}
	if masterDetails[0].SpecialKey != value.SpecialKey {
		return &models.LoginResponseMaster{
			Message:  "Master Not found or key provided is incorrect",
			MasterId: uuid.Nil,
		}, nil
	}
	loginDetails, err := ls.LoginRepo.FindById(masterDetails[0].Id)
	if err != nil {
		return nil, errors.New("error while connecting to the database: " + err.Error())
	}
	if loginDetails.UserId == masterDetails[0].Id {
		err := ls.LoginRepo.ReLogin(masterDetails[0].Id)
		if err != nil {
			return nil, errors.New("error while reloging the user: " + err.Error())
		}
		return &models.LoginResponseMaster{
			Message:  "Reloging in user",
			MasterId: masterDetails[0].Id,
		}, nil
	}

	err = ls.LoginRepo.Create(&models.DBLogin{
		UserId:   masterDetails[0].Id,
		IsLogin:  true,
		IsMaster: true,
	})
	if err != nil {
		return nil, errors.New("error while creating a login entry in the database: " + err.Error())
	}
	return &models.LoginResponseMaster{
		Message:  "Login of master successfull for user " + masterDetails[0].Name + " .",
		MasterId: masterDetails[0].Id,
	}, nil
}

func (ls *LoginServiceImpl) LoginUser(value *models.UserLoginRequest) (*models.LoginResponse, error) {
	err := ls.SetupDalLayer()
	if err != nil {
		return nil, err
	}
	config, err := helpers.ReadConfig("/app/config.yml")
	if err != nil {
		return nil, errors.New("error while connecting to env file: " + err.Error())
	}
	var Name string
	var userID uuid.UUID
	var masterID uuid.UUID
	if config.Algorithm == "RSA" {
		userDetails, userErr := ls.UserRepo.FindByRSA(&models.DBRSAUser{
			Email: value.Email,
		})
		if userErr != nil {
			return nil, errors.New("error while finding user from database: " + userErr.Error())
		}
		if userDetails.Email != value.Email {
			return &models.LoginResponse{
				Message:  "User not found",
				UserId:   uuid.Nil,
				MasterId: uuid.Nil,
			}, nil
		}
		loginDetails, err := ls.LoginRepo.FindById(userDetails.UserId)
		if err != nil {
			return nil, errors.New("error while connecting to the database: " + err.Error())
		}
		if loginDetails.UserId == userDetails.UserId {
			err := ls.LoginRepo.ReLogin(userDetails.UserId)
			if err != nil {
				return nil, errors.New("error while reloging the user: " + err.Error())
			}
			return &models.LoginResponse{
				Message:  "Reloging in user",
				UserId:   userDetails.UserId,
				MasterId: userDetails.MasterId,
			}, nil
		}
		if userDetails.Password != value.Password {
			return &models.LoginResponse{
				Message:  "Password is incorrect",
				UserId:   uuid.Nil,
				MasterId: uuid.Nil,
			}, nil
		}
		Name = userDetails.Name
		userID = userDetails.UserId
		masterID = userDetails.MasterId
	} else if config.Algorithm == "ASA" {
		userDetails, userErr := ls.UserRepo.FindByASA(&models.DBASAUser{
			Email: value.Email,
		})
		if userErr != nil {
			return nil, errors.New("error while finding user from database: " + userErr.Error())
		}
		if userDetails.Email != value.Email {
			return &models.LoginResponse{
				Message:  "User not found",
				UserId:   uuid.Nil,
				MasterId: uuid.Nil,
			}, nil
		}
		loginDetails, err := ls.LoginRepo.FindById(userDetails.UserId)
		if err != nil {
			return nil, errors.New("error while connecting to the database: " + err.Error())
		}
		if loginDetails.UserId == userDetails.UserId {
			err := ls.LoginRepo.ReLogin(userDetails.UserId)
			if err != nil {
				return nil, errors.New("error while reloging the user: " + err.Error())
			}
			return &models.LoginResponse{
				Message:  "Reloging in user",
				UserId:   userDetails.UserId,
				MasterId: userDetails.MasterId,
			}, nil
		}

		if userDetails.Password != value.Password {
			return &models.LoginResponse{
				Message:  "Password is incorrect",
				UserId:   uuid.Nil,
				MasterId: uuid.Nil,
			}, nil
		}
		Name = userDetails.Name
		userID = userDetails.UserId
		masterID = userDetails.MasterId
	}
	err = ls.LoginRepo.Create(&models.DBLogin{
		UserId:   userID,
		IsLogin:  true,
		IsMaster: false,
	})
	if err != nil {
		return nil, errors.New("error while creating a login entry in the database: " + err.Error())
	}
	return &models.LoginResponse{
		Message:  "User Logged in successfully: " + Name,
		UserId:   userID,
		MasterId: masterID,
	}, nil
}

func (ls *LoginServiceImpl) Logout(userID uuid.UUID) (*models.SuccessResponse, error) {
	err := ls.SetupDalLayer()
	if err != nil {
		return nil, err
	}
	err = ls.LoginRepo.Logout(userID)
	if err != nil {
		return nil, err
	}
	return &models.SuccessResponse{
		Message: "user logged out successfully",
	}, nil
}
