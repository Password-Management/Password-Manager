package dallayer

import (
	"log"
	"password-manager/db"
	"password-manager/models"
)

func NewCredsDalRequest() (Creds, error) {
	return &CredsImpl{}, nil
}

type CredsImpl struct{}

type Creds interface {
	Create(value *models.DbCreds) error
	FindBy(userId string) (otp string, err error)
	Update(userId string) error
}

func (c *CredsImpl) Create(value *models.DbCreds) error {
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

func (c *CredsImpl) FindBy(userId string) (otp string, err error) {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request")
		return "", err
	}
	dbConn, err := db.InitDB()
	if err != nil {
		return "", err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return "", transaction.Error
	}
	var response models.DbCreds
	defer transaction.Rollback()
	credsDetails := transaction.Model(models.DbCreds{}).Where("user_id = ? and is_used = ?", userId, false).Find(&response)
	if credsDetails.Error != nil {
		return "", credsDetails.Error
	}
	return response.Otp, nil
}

func (c *CredsImpl) Update(userId string) error {
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
	updateDetails := transaction.Model(models.DbCreds{}).Where("user_id = ?", userId).Update("is_used", true)
	if updateDetails.Error != nil {
		return updateDetails.Error
	}
	transaction.Commit()
	return nil
}
