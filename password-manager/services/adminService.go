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
)

type AdminSericeImpl struct {
	MasterRepo dallayer.Master
}

type AdminService interface {
	Create() error
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
		return errors.New("error while connecting to the master repo for checking the master count.")
	}
	if len(resp) == 1 {
		return errors.New("Already a master exist can create more")
	}
	config, err := helpers.ReadConfig("/app/config.yml")
	if err != nil {
		return errors.New("error while reading the config: " + err.Error())
	}
	if config.Email == "" || config.Name == "" || config.ProductType == "" {
		return errors.New("The queue has not receive any value can't create the user")
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
	var body string
	body = "This is your SpecialKey" + " " + specialKey + "." + "\n" + "You can edit this key only once and its manditaory"
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
