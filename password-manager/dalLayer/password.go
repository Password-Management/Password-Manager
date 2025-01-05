package dallayer

import (
	"errors"
	"log"
	"password-manager/db"
	"password-manager/models"

	"github.com/google/uuid"
)

func NewPasswordDalRequest() (Password, error) {
	return &PasswordImpl{}, nil
}

type PasswordImpl struct{}

type Password interface {
	Create(value *models.DbPassword) error
	FindWebsiteName(wesbite string, userId uuid.UUID) (*models.DbPassword, error)
	FindAll(userId uuid.UUID) ([]*models.DbPassword, error)
	DeleteUserPassword(userId uuid.UUID) error
	DeletePassword(website string) error
}

func (ps *PasswordImpl) Create(value *models.DbPassword) error {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request")
		return nil
	}
	dbConn, err := db.InitDB()
	if err != nil {
		return err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return transaction.Error
	}
	defer transaction.Rollback()
	state := transaction.Create(&value)
	if state.Error != nil {
		return state.Error
	}
	transaction.Commit()
	return nil
}

func (ps *PasswordImpl) FindWebsiteName(wesbite string, userId uuid.UUID) (*models.DbPassword, error) {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request")
		return nil, err
	}
	dbConn, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	var response *models.DbPassword
	defer transaction.Rollback()
	result := transaction.Last(&response, models.DbPassword{
		UserId:      userId,
		WebisteName: wesbite,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return response, nil

}

func (ps *PasswordImpl) FindAll(userId uuid.UUID) ([]*models.DbPassword, error) {
	db, err := db.NewDbRequest()
	if err != nil {
		return nil, errors.New("error while creating the DB request: " + err.Error())
	}
	dbConn, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	var response []*models.DbPassword
	defer transaction.Rollback()
	result := transaction.Find(&response, models.DbPassword{
		UserId: userId,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return response, nil
}

func (ps *PasswordImpl) DeleteUserPassword(userId uuid.UUID) error {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request")
		return err
	}
	dbConn, err := db.InitDB()
	if err != nil {
		return err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return transaction.Error
	}
	if err := transaction.Unscoped().Where("").Delete(nil, &models.DbPassword{
		UserId: userId,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (ps *PasswordImpl) DeletePassword(website string) error {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request")
		return err
	}
	dbConn, err := db.InitDB()
	if err != nil {
		return err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return transaction.Error
	}
	defer transaction.Rollback()
	if err := transaction.Unscoped().Where("website_name = ?", website).Delete(&models.DbPassword{
		WebisteName: website,
	}).Error; err != nil {
		return err
	}
	transaction.Commit()
	return nil
}