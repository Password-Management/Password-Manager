package dallayer

import (
	"log"
	"password-manager/db"
	"password-manager/models"
)

func NewLoginDalRequest() (Login, error) {
	return &LoginImpl{}, nil
}

type LoginImpl struct{}

type Login interface {
	Create(*models.DBLogin) error
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
