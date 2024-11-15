package dallayer

import (
	"errors"
	"fmt"
	"log"
	"password-manager/db"
	"password-manager/models"

	"github.com/google/uuid"
)

func NewLoginDalRequest() (Login, error) {
	return &LoginImpl{}, nil
}

type LoginImpl struct{}

type Login interface {
	Create(*models.DBLogin) error
	Delete(userId uuid.UUID) error
	Logout(userId uuid.UUID) error
	FindById(userid uuid.UUID) (*models.DBLogin, error)
}

func (ls *LoginImpl) Create(value *models.DBLogin) error {
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

func (ls *LoginImpl) FindById(userid uuid.UUID) (*models.DBLogin, error) {
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
	var response *models.DBLogin
	loginDetails := transaction.Find(response, &models.DBLogin{
		UserId: userid,
	})
	if loginDetails.Error != nil {
		return nil, loginDetails.Error
	}
	return response, nil
}

func (ls *LoginImpl) Logout(userId uuid.UUID) error {
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
	var userDetails *models.DBLogin
	resp := transaction.Find(&userDetails, &models.DBLogin{
		UserId: userId,
	})
	if resp.Error != nil {
		return resp.Error
	}
	if userDetails.IsLogin == false {
		return errors.New("user is not logged in.")
	}
	fmt.Println("From the repo layer = ", userDetails)
	updateDetails := transaction.Model(models.DBLogin{}).Where("user_id = ?", userId).Update("is_login", false)
	if updateDetails.Error != nil {
		return updateDetails.Error
	}
	transaction.Commit()
	return nil
}

func (ls *LoginImpl) Delete(userId uuid.UUID) error {
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
	deleteUser := transaction.Unscoped().Delete(nil, &models.DBLogin{
		UserId: userId,
	})
	if deleteUser.Error != nil {
		return deleteUser.Error
	}
	return nil
}
