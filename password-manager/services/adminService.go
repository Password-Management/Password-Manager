// This is the admin Service no UI is there.
package services

import (
	"errors"
	"log"
	dallayer "password-manager/dalLayer"
	"password-manager/db"
	"password-manager/helpers"
	"password-manager/models"
	"time"

	"github.com/google/uuid"
)

type AdminSericeImpl struct {
	MasterRepo dallayer.Master
	CredsRepo  dallayer.Creds
}

type AdminService interface {
	Create() error
	CreateOTP(userId string) (response *models.SuccessResponse, err error)
	VerifyOTP(userId string, otp string) (response *models.SuccessResponse, err error)
	GetPlanInformation(customerId string) (response *models.SuccessResponse, err error)
}

func NewAdminService() (AdminService, error) {
	return &AdminSericeImpl{}, nil
}

func (ad *AdminSericeImpl) SetupDalLayer() error {
	var err error
	ad.MasterRepo, err = dallayer.NewMasterDalRequest()
	if err != nil {
		return errors.New("error while connecting to the master Dal layer: " + err.Error())
	}
	ad.CredsRepo, err = dallayer.NewCredsDalRequest()
	if err != nil {
		return errors.New("error while connecting to the creds Dal layer: " + err.Error())
	}
	return nil
}

func (ad *AdminSericeImpl) Create() error {
	log.Println("Inside the service")
	err := ad.SetupDalLayer()
	if err != nil {
		return err
	}
	resp, err := ad.MasterRepo.FindAll()
	if err != nil {
		return errors.New("error while connecting to the master repo for checking the master count")
	}
	if len(resp) == 1 {
		return errors.New("already a master exist can create more")
	}
	config, err := helpers.ReadConfig("/app/config.yml")
	if err != nil {
		return errors.New("error while reading the config: " + err.Error())
	}
	if config.Email == "" || config.Name == "" || config.ProductType == "" {
		return errors.New("the config file is not present please get the config file")
	}
	specialKey := helpers.GenerateSpecialKey()
	err = CreateUserTable(config.Algorithm)
	if err != nil {
		return err
	}
	value := &models.DBMaster{
		CreatedAt:  time.Now(),
		Name:       config.Name,
		Email:      config.Email,
		Algorithm:  config.Algorithm,
		SpecialKey: specialKey,
		Count:      0,
		Plan:       config.ProductType,
	}
	err = ad.MasterRepo.Create(value)
	if err != nil {
		return errors.New("error while creating the master entry: " + err.Error())
	}

	body := "This is your SpecialKey" + " " + specialKey + "." + "\n" + "You can edit this key only once and its manditaory"
	err = helpers.SendEmail(body, "This the AuthKey")
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func CreateUserTable(algo string) error {
	log.Println("Inside the update alogrithm function")
	db, err := db.NewDbRequest()
	if err != nil {
		return errors.New("error while creating a DB request: " + err.Error())
	}
	database, err := db.InitDB()
	log.Println("After database init")
	if err != nil {
		return errors.New("error in starting the DataBase: " + err.Error())
	}
	if algo == "ASA" {
		err := database.AutoMigrate(&models.DBASAUser{})
		if err != nil {
			return errors.New("error while creating  the table RSA: " + err.Error())
		}
	} else {
		err := database.AutoMigrate(&models.DBRSAUser{})
		if err != nil {
			return errors.New("error while creating the table ASA: " + err.Error())
		}
	}
	log.Println("After creating the database")
	return nil
}

func (ad *AdminSericeImpl) CreateOTP(userId string) (response *models.SuccessResponse, err error) {
	err = ad.SetupDalLayer()
	if err != nil {
		return nil, err
	}
	otpDetails, err := ad.CredsRepo.FindBy(userId)
	if err != nil {
		return nil, errors.New("error while finding already existing otp for the user: " + err.Error())
	}
	if otpDetails != "" {
		return &models.SuccessResponse{Message: "OTP already exists."}, nil
	}
	otp := helpers.GenerateRandomString(8)
	err = ad.CredsRepo.Create(&models.DbCreds{
		UserId: uuid.MustParse(userId),
		Otp:    otp,
		IsUsed: false,
	})
	if err != nil {
		return nil, errors.New("error while creating a OTP entry: " + err.Error())
	}
	return &models.SuccessResponse{
		Message: "OTP creation was successfull.",
	}, nil
}

func (ad *AdminSericeImpl) VerifyOTP(userId string, otp string) (response *models.SuccessResponse, err error) {
	err = ad.SetupDalLayer()
	if err != nil {
		return nil, err
	}
	otpDetails, err := ad.CredsRepo.FindBy(userId)
	if err != nil {
		return nil, errors.New("error while finding the otp details: " + err.Error())
	}
	if otpDetails == otp {
		updateDetailErr := ad.CredsRepo.Update(userId)
		if updateDetailErr != nil {
			return nil, errors.New("error while updating the status of the otp: " + updateDetailErr.Error())
		}
	} else {
		return &models.SuccessResponse{
			Message: "OTP doesnt match",
		}, nil
	}
	return &models.SuccessResponse{
		Message: "SUCCESS",
	}, nil
}

func (ad *AdminSericeImpl) GetPlanInformation(customerId string) (response *models.SuccessResponse, err error) {
	err = ad.SetupDalLayer()
	if err != nil {
		return nil, err
	}
	resp, err := ad.MasterRepo.FindBy(&models.DBMaster{
		CustomerId: uuid.MustParse(customerId),
	})
	if err != nil {
		return nil, errors.New("error while getting the master information")
	}
	return &models.SuccessResponse{
		Message: resp.Plan,
	}, nil
}
