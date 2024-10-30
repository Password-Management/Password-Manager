package dallayer

import (
	"log"
	"password-manager/db"
	"password-manager/models"
)

func NewMasterDalRequest() (Master, error) {
	return &MasterImpl{}, nil
}

type Master interface {
	Create(value *models.DBMaster) error
	FindAll() ([]*models.DBMaster, error)
	FindBy(condition *models.DBMaster) (*models.DBMaster, error)
	Update(value *models.DBMaster) (*models.DBMaster, error)
}

type MasterImpl struct {}

func (mas *MasterImpl) Create(value *models.DBMaster) error {
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

func (mas *MasterImpl) FindAll() ([]*models.DBMaster, error) {
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
	var response []*models.DBMaster
	defer transaction.Rollback()
	result := transaction.Find(&response)
	if result.Error != nil {
		return nil, result.Error
	}
	return response, nil
}

func (mas *MasterImpl) FindBy(condition *models.DBMaster) (*models.DBMaster, error) {
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
	var response *models.DBMaster
	resp := transaction.Last(&response, condition)
	if resp.Error != nil {
		return nil, resp.Error
	}
	transaction.Commit()
	return response, nil
}

func (mas *MasterImpl) Update(value *models.DBMaster) (*models.DBMaster, error) {
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
	resp := transaction.Save(value)
	if resp.Error != nil {
		return nil, resp.Error
	}
	transaction.Commit()
	return value, nil
}
