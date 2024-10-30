package services

import (
	"errors"
	"log"
	"math/rand"
	dallayer "password-manager/dalLayer"
	"password-manager/encryption"
	"password-manager/helpers"
	"password-manager/models"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var customerEmail string
var customerName string

func NewMasterServiceRequest() (MasterService, error) {
	return &MasterServiceImpl{}, nil
}

type MasterServiceImpl struct {
	MasterRepo dallayer.Master
	UserRepo   dallayer.User
}

type MasterService interface {
	Create() error
	EditKey(value *models.EditKeyRequest) (response *models.SuccessResponse, err error)
	GetInfo(specialKey string) (response *models.GetInfoResponse, err error)
	UpdateAlgorithm(value *models.UpdateAlgorithmRequest) (response *models.SuccessResponse, err error)
	CreateUser(value *models.CreateUserRequest) (response *models.SuccessResponse, err error)
	ListUser(specialKey string) (response []*models.GetUserListResponse, err error)
}

func (ms *MasterServiceImpl) SetupDalLayer() error {
	var err error
	ms.MasterRepo, err = dallayer.NewMasterDalRequest()
	if err != nil {
		return errors.New("error while connecting to the master Dal layer: " + err.Error())
	}

	ms.UserRepo, err = dallayer.NewUserDalRequest()
	if err != nil {
		return errors.New("error while connecting to the user Dal layer: " + err.Error())
	}
	return nil
}

func generateSpecialKey() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := rand.Intn(900000) + 100000
	return strconv.Itoa(randomNumber)
}

func (ms *MasterServiceImpl) Create() error {
	err := ms.SetupDalLayer()
	if err != nil {
		return err
	}
	specialKey := generateSpecialKey()
	config, err := helpers.ReadConfig("/app/config.yml")
	if err != nil {
		return errors.New("error while reading the config: " + err.Error())
	}
	value := &models.DBMaster{
		Name:       config.Name,
		Email:      config.Email,
		Algorithm:  "test",
		SpecialKey: specialKey,
		Count:      0,
	}
	err = ms.MasterRepo.Create(value)
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

func (ms *MasterServiceImpl) EditKey(value *models.EditKeyRequest) (response *models.SuccessResponse, err error) {
	err = ms.SetupDalLayer()
	if err != nil {
		return nil, err
	}
	specialKey := value.SpecialKey
	resp, err := ms.MasterRepo.FindAll()
	if err != nil {
		return nil, errors.New("unable get the details of the master entry: " + err.Error())
	}
	if specialKey != resp[0].SpecialKey {
		return nil, errors.New("Please check your mail wrong special key is given")
	}
	if resp[0].Count == 1 {
		return nil, errors.New("You have already updated the key if lost contact support team")
	}
	resp[0].SpecialKey = value.NewKey
	resp[0].Count = 1
	respUpdate, err := ms.MasterRepo.Update(resp[0])
	if err != nil {
		return nil, errors.New("error while updating the entry for the master: " + err.Error())
	}
	return &models.SuccessResponse{
		Message: "Your Key is updated to " + respUpdate.SpecialKey,
	}, nil
}

func (ms *MasterServiceImpl) GetInfo(specialKey string) (response *models.GetInfoResponse, err error) {
	err = ms.SetupDalLayer()
	if err != nil {
		return nil, err
	}
	resp, err := ms.MasterRepo.FindAll()
	if err != nil {
		return nil, errors.New("unable get the details of the master entry: " + err.Error())
	}
	if specialKey != resp[0].SpecialKey {
		return nil, errors.New("Please check your mail wrong special key is given")
	}
	return &models.GetInfoResponse{
		Name:      resp[0].Name,
		Email:     resp[0].Email,
		Algorithm: resp[0].Algorithm,
		Count:     resp[0].Count,
	}, nil
}

func (ms *MasterServiceImpl) UpdateAlgorithm(value *models.UpdateAlgorithmRequest) (response *models.SuccessResponse, err error) {
	err = ms.SetupDalLayer()
	if err != nil {
		return nil, err
	}
	resp, err := ms.MasterRepo.FindAll()
	if err != nil {
		return nil, errors.New("unable get the details of the master entry: " + err.Error())
	}
	if value.SpecialKey != resp[0].SpecialKey {
		return nil, errors.New("Please check your mail wrong special key is given")
	}
	resp[0].Algorithm = value.NewAlgorithm
	respUpdate, err := ms.MasterRepo.Update(resp[0])
	if err != nil {
		return nil, errors.New("error while updating the algorithm " + err.Error())
	}
	return &models.SuccessResponse{
		Message: "Your Algorithm is updated to " + respUpdate.Algorithm,
	}, nil
}

func (ms *MasterServiceImpl) CreateUser(value *models.CreateUserRequest) (response *models.SuccessResponse, err error) {
	err = ms.SetupDalLayer()
	log.Println("the requestBody: ", value)
	if err != nil {
		return nil, err
	}
	resp, err := ms.MasterRepo.FindAll()
	if err != nil {
		return nil, errors.New("unable get the details of the master entry: " + err.Error())
	}
	if uuid.MustParse(value.MasterId) != resp[0].Id {
		return nil, errors.New("Please check your tenant information")
	}
	userInfo, err := ms.UserRepo.FindAll()
	if err != nil {
		return nil, errors.New("error while fetching all the userInfo: " + err.Error())
	}
	for _, info := range userInfo {
		if info.Email == value.Email || info.Name == value.Name {
			return nil, errors.New("this user already exists.")
		}
	}
	privateKey, publicKey, err := encryption.GenerateRSAKeys()
	if err != nil {
		return nil, errors.New("error while generating the keys: " + err.Error())
	}
	publicKeyDb, err := encryption.PublicKeyToPEM(publicKey)
	if err != nil {
		return nil, errors.New("error while converting the public key to the PEM: " + err.Error())
	}
	hashedSpecial, err := encryption.HashPassword(value.SpecialKey)
	if err != nil {
		return nil, errors.New("error while hashing the user's special key: " + err.Error())
	}
	privateKeyDb := encryption.PrivateKeyToPEM(privateKey)
	User := models.DBUser{
		Name:       value.Name,
		Password:   value.Password,
		PublicKey:  publicKeyDb,
		PrivateKey: privateKeyDb,
		MasterId:   uuid.MustParse(value.MasterId),
		Email:      value.Email,
		SpecialKey: hashedSpecial,
	}
	err = ms.UserRepo.Create(&User)
	if err != nil {
		return nil, errors.New("error while creating the user entry by the master: " + err.Error())
	}
	return &models.SuccessResponse{
		Message: "Creation of User " + value.Name + " was successfull.",
	}, nil
}

func (ms *MasterServiceImpl) ListUser(specialKey string) (response []*models.GetUserListResponse, err error) {
	err = ms.SetupDalLayer()
	if err != nil {
		return nil, err
	}
	resp, err := ms.MasterRepo.FindAll()
	if err != nil {
		return nil, errors.New("unable get the details of the master entry: " + err.Error())
	}
	if specialKey != resp[0].SpecialKey {
		return nil, errors.New("Please check your mail wrong special key is given")
	}
	userInfo, err := ms.UserRepo.FindAll()
	if err != nil {
		return nil, errors.New("error while fetching all the userInfo: " + err.Error())
	}
	for _, info := range userInfo {
		var userResp models.GetUserListResponse
		userResp.Email = info.Email
		userResp.Name = info.Name
		response = append(response, &userResp)
	}
	return response, nil
}
