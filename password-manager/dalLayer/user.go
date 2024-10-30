package dallayer

import (
	"log"
	"password-manager/db"
	"password-manager/models"

	"github.com/google/uuid"
)

func NewUserDalRequest() (User, error) {
	return &UserImpl{}, nil
}

type UserImpl struct{}

type User interface {
	Create(value *models.DBUser) error
	FindAll() ([]*models.DBUser, error)
	FindById(userId uuid.UUID, masterId uuid.UUID) (*models.DBUser, error)
}

func (us *UserImpl) Create(value *models.DBUser) error {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request for the user Create service: ", err)
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

func (us *UserImpl) FindAll() ([]*models.DBUser, error) {
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
	var response []*models.DBUser
	defer transaction.Rollback()
	result := transaction.Find(&response)
	if result.Error != nil {
		return nil, result.Error
	}
	return response, nil
}

func (us *UserImpl) FindById(userId uuid.UUID, masterId uuid.UUID) (*models.DBUser, error) {
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
	var response *models.DBUser
	defer transaction.Rollback()
	result := transaction.Last(&response, models.DBUser{
		Id: userId,
		MasterId: masterId,
	})
	log.Println("the result value = ", result)
	if result.Error != nil {
		return nil, result.Error
	}
	return response, nil
}

