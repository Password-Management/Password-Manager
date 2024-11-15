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
	CreateRSA(value *models.DBRSAUser) error
	CreateASA(value *models.DBASAUser) error
	FindAllRSAUser() ([]*models.DBRSAUser, error)
	FindAllASAUser() ([]*models.DBASAUser, error)
	FindById(userId uuid.UUID, masterId uuid.UUID) (*models.DBRSAUser, error)
	FindByRSA(condition *models.DBRSAUser) (*models.DBRSAUser, error)
	FindByASA(condition *models.DBASAUser) (*models.DBASAUser, error)
	Delete(id uuid.UUID, userType string) error
}

func (us *UserImpl) CreateRSA(value *models.DBRSAUser) error {
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

func (us *UserImpl) FindAllRSAUser() ([]*models.DBRSAUser, error) {
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
	var response []*models.DBRSAUser
	defer transaction.Rollback()
	result := transaction.Find(&response)
	if result.Error != nil {
		return nil, result.Error
	}
	return response, nil
}

func (us *UserImpl) FindAllASAUser() ([]*models.DBASAUser, error) {
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
	var response []*models.DBASAUser
	defer transaction.Rollback()
	result := transaction.Find(&response)
	if result.Error != nil {
		return nil, result.Error
	}
	return response, nil
}

func (us *UserImpl) FindById(userId uuid.UUID, masterId uuid.UUID) (*models.DBRSAUser, error) {
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
	var response *models.DBRSAUser
	defer transaction.Rollback()
	result := transaction.Last(&response, models.DBRSAUser{
		Id:       userId,
		MasterId: masterId,
	})
	log.Println("the result value = ", result)
	if result.Error != nil {
		return nil, result.Error
	}
	return response, nil
}

func (us *UserImpl) CreateASA(value *models.DBASAUser) error {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request for the user Create service: ", err)
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
	state := transaction.Create(&value)
	if state.Error != nil {
		return state.Error
	}
	transaction.Commit()
	return nil
}

func (us *UserImpl) FindByRSA(condition *models.DBRSAUser) (*models.DBRSAUser, error) {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request for the RSA user find service: ", err)
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
	var response *models.DBRSAUser
	resp := transaction.Find(&response, &condition)
	if resp.Error != nil {
		return nil, resp.Error
	}
	defer transaction.Rollback()
	return response, nil
}

func (us *UserImpl) FindByASA(condition *models.DBASAUser) (*models.DBASAUser, error) {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request for the RSA user find service: ", err)
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
	var response *models.DBASAUser
	resp := transaction.Find(&response, &condition)
	if resp.Error != nil {
		return nil, resp.Error
	}
	defer transaction.Rollback()
	return response, nil
}

func (us *UserImpl) Delete(id uuid.UUID, userType string) error {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error while creating the db request for deleteuser: ", err)
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
	if userType == "RSA" {
		deleteUser := transaction.Unscoped().Delete(nil, &models.DBRSAUser{
			Id: id,
		})
		if deleteUser.Error != nil {
			return deleteUser.Error
		}
	} else {
		deleteUser := transaction.Unscoped().Delete(nil, &models.DBASAUser{
			Id: id,
		})
		if deleteUser.Error != nil {
			return deleteUser.Error
		}
	}
	return nil
}
