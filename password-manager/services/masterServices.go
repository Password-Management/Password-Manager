package services

import (
	"errors"
	"log"
	dallayer "password-manager/dalLayer"
	"password-manager/db"
	"password-manager/encryption"
	"password-manager/helpers"
	"password-manager/models"
	"time"

	"github.com/google/uuid"
)

func NewMasterServiceRequest() (MasterService, error) {
	return &MasterServiceImpl{}, nil
}

type MasterServiceImpl struct {
	MasterRepo   dallayer.Master
	UserRepo     dallayer.User
	PasswordRepo dallayer.Password
	LoginRepo    dallayer.Login
}

type MasterService interface {
	EditKey(value *models.EditKeyRequest) (response *models.SuccessResponse, err error)
	GetInfo(specialKey string) (response *models.GetInfoResponse, err error)
	UpdateAlgorithm(value *models.UpdateAlgorithmRequest) (response *models.SuccessResponse, err error)
	CreateUser(value *models.CreateUserRequest, masterId string) (response *models.SuccessResponse, err error)
	ListUser(specialKey string) (response []*models.GetUserListResponse, err error)
	DeleteUser(userId uuid.UUID) (response *models.SuccessResponse, err error)
	GetUserByEmail(email string, masterId string) (response *models.GetUserByEmailResponse, err error)
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

	ms.PasswordRepo, err = dallayer.NewPasswordDalRequest()
	if err != nil {
		return errors.New("error while connecting to the master Dal layer: " + err.Error())
	}

	ms.LoginRepo, err = dallayer.NewLoginDalRequest()
	if err != nil {
		return errors.New("error while connecting to the login Dal layer from service: " + err.Error())
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
		return nil, errors.New("please check your mail wrong special key is given")
	}
	if resp[0].Count == 1 {
		return nil, errors.New("you have already updated the key if lost contact support team")
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
	if uuid.MustParse(specialKey) != resp[0].Id {
		return nil, errors.New("please check your mail wrong special key is given")
	}
	return &models.GetInfoResponse{
		Name:       resp[0].Name,
		Email:      resp[0].Email,
		Algorithm:  resp[0].Algorithm,
		Count:      resp[0].Count,
		Plan:       resp[0].Plan,
		CustomerId: resp[0].CustomerId.String(),
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
		return nil, errors.New("please check your mail wrong special key is given")
	}
	if resp[0].Algorithm == value.NewAlgorithm {
		return nil, errors.New("already the same algorithm")
	}
	resp[0].Algorithm = value.NewAlgorithm
	respUpdate, err := ms.MasterRepo.Update(resp[0])
	if err != nil {
		return nil, errors.New("error while updating the algorithm " + err.Error())
	}
	tableErr := ms.UpdateUserTable(value.NewAlgorithm)
	if tableErr != nil {
		return nil, errors.New("error while updating the algorithm: " + tableErr.Error())
	}
	return &models.SuccessResponse{
		Message: "Your Algorithm is updated to " + respUpdate.Algorithm,
	}, nil
}

func (ms *MasterServiceImpl) CreateUser(value *models.CreateUserRequest, masterId string) (response *models.SuccessResponse, err error) {
	log.Println("Inside the service file")
	err = ms.SetupDalLayer()
	if err != nil {
		return nil, err
	}
	resp, err := ms.MasterRepo.FindAll()
	if err != nil {
		return nil, errors.New("unable get the details of the master entry: " + err.Error())
	}
	if uuid.MustParse(masterId) != resp[0].Id {
		return nil, errors.New("please check your tenant information")
	}
	userInfo, err := ms.UserRepo.FindAllRSAUser()
	if err != nil {
		return nil, errors.New("error while fetching all the userInfo: " + err.Error())
	}
	for _, info := range userInfo {
		if info.Email == value.Email || info.Name == value.Name {
			return nil, errors.New("this user already exists")
		}
	}
	log.Println("The Alogorithm: ", resp[0].Algorithm)
	if resp[0].Algorithm == "ASA" {
		err := ms.createASA(value, masterId)
		if err != nil {
			return nil, err
		}
	} else if resp[0].Algorithm == "RSA" {
		err := ms.createRSA(value, masterId)
		if err != nil {
			return nil, err
		}
	}
	return &models.SuccessResponse{
		Message: "Creation of User " + value.Name + " was successfull.",
	}, nil
}

func (ms *MasterServiceImpl) createASA(value *models.CreateUserRequest, masterId string) error {
	passwordKey, err := encryption.GenerateKey(32)
	if err != nil {
		return errors.New("error while generating ASA key for the user: " + err.Error())
	}
	password := helpers.GenerateRandomString(8)
	specialKey := helpers.GenerateSpecialKey()
	User := models.DBASAUser{
		CreatedAt:   time.Now(),
		CreatedBy:   uuid.MustParse(masterId),
		Name:        value.Name,
		Password:    password,
		SpecialKey:  specialKey,
		PasswordKey: passwordKey,
		MasterId:    uuid.MustParse(masterId),
		Email:       value.Email,
		IsMaster:    value.IsMaster,
		UserId:      uuid.New(),
	}
	err = ms.UserRepo.CreateASA(&User)
	if err != nil {
		return errors.New("error while creating the user entry by the master: " + err.Error())
	}
	body := "This is your SpecialKey" + " " + specialKey + "." + "\n" + "You can edit this key as this is generated by the admin while registering you." + "\n" + "This is your auto-generated password please edit before logging it in else it's an security breach " + password
	err = createEmailforUser(body)
	if err != nil {
		return err
	}
	return nil
}

func (ms *MasterServiceImpl) createRSA(value *models.CreateUserRequest, masterId string) error {
	privateKey, publicKey, err := encryption.GenerateRSAKeys()
	if err != nil {
		return errors.New("error while generating the keys: " + err.Error())
	}
	publicKeyDb, err := encryption.PublicKeyToPEM(publicKey)
	if err != nil {
		return errors.New("error while converting the public key to the PEM: " + err.Error())
	}
	password := helpers.GenerateRandomString(8)
	specialKey := helpers.GenerateSpecialKey()
	privateKeyDb := encryption.PrivateKeyToPEM(privateKey)
	User := models.DBRSAUser{
		CreatedAt:  time.Now(),
		CreatedBy:  uuid.MustParse(masterId),
		Name:       value.Name,
		Email:      value.Email,
		UserId:     uuid.New(),
		Password:   password,
		PublicKey:  publicKeyDb,
		PrivateKey: privateKeyDb,
		SpecialKey: specialKey,
		MasterId:   uuid.MustParse(masterId),
		IsMaster:   value.IsMaster,
	}
	err = ms.UserRepo.CreateRSA(&User)
	if err != nil {
		return errors.New("error while creating the user entry by the master: " + err.Error())
	}
	body := "This is your SpecialKey" + " " + specialKey + "." + "\n" + "You can edit this key as this is generated by the admin while registering you." + "\n" + "This is your auto-generated password please edit before logging it in else it's an security breach " + password
	err = createEmailforUser(body)
	if err != nil {
		return err
	}
	return nil
}

func createEmailforUser(body string) error {
	subject := "Registered to the KeyPass as a user"
	err := helpers.SendEmail(body, subject)
	if err != nil {
		return err
	}
	return nil
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
		return nil, errors.New("please check your mail wrong special key is given")
	}
	userInfo, err := ms.UserRepo.FindAllRSAUser()
	if err != nil {
		return nil, errors.New("error while fetching all the userInfo: " + err.Error())
	}
	if len(userInfo) == 0 {
		return []*models.GetUserListResponse{}, nil
	}
	for _, info := range userInfo {
		var userResp models.GetUserListResponse
		userResp.Email = info.Email
		userResp.Name = info.Name
		response = append(response, &userResp)
	}
	return response, nil
}

func (ms *MasterServiceImpl) GetUserByEmail(email string, masterId string) (response *models.GetUserByEmailResponse, err error) {
	err = ms.SetupDalLayer()
	if err != nil {
		return nil, err
	}
	resp, err := ms.MasterRepo.FindAll()
	if err != nil {
		return nil, errors.New("unable get the details of the master entry: " + err.Error())
	}
	if resp[0].Id.String() != masterId {
		return nil, errors.New("the masterId is wrong")
	}
	if resp[0].Algorithm == "RSA" {
		userInfo, err := ms.UserRepo.FindByRSA(&models.DBRSAUser{
			Email: email,
		})
		if err != nil {
			return nil, errors.New("error while finding the userdetails: " + err.Error())
		}
		return &models.GetUserByEmailResponse{
			Email:  userInfo.Email,
			UserId: userInfo.UserId.String(),
		}, nil
	} else {
		userInfo, err := ms.UserRepo.FindByASA(&models.DBASAUser{
			Email: email,
		})
		if err != nil {
			return nil, errors.New("error while finding the userdetails: " + err.Error())
		}
		return &models.GetUserByEmailResponse{
			Email:  userInfo.Email,
			UserId: userInfo.UserId.String(),
		}, nil
	}
}

func (ms *MasterServiceImpl) DeleteUser(userId uuid.UUID) (response *models.SuccessResponse, err error) {
	err = ms.SetupDalLayer()
	if err != nil {
		return nil, err
	}
	config, err := helpers.ReadConfig("/app/config.yml")
	if err != nil {
		return nil, errors.New("error while reading the config: " + err.Error())
	}
	flag := false
	userType := config.Algorithm
	if userType == "ASA" {
		userDetails, err := ms.UserRepo.FindAllASAUser()
		if err != nil {
			return nil, err
		}
		for _, info := range userDetails {
			if info.UserId == userId {
				flag = true
			}
		}
	} else {
		userDetails, err := ms.UserRepo.FindAllRSAUser()
		if err != nil {
			return nil, err
		}
		for _, info := range userDetails {
			if info.UserId == userId {
				flag = true
			}
		}
	}
	if !flag {
		return nil, errors.New("user doesn't exist")
	}

	passwordDetails, err := ms.PasswordRepo.FindAll(userId)
	if err != nil {
		return nil, err
	}
	if len(passwordDetails) > 0 {
		for _, details := range passwordDetails {
			deletePassword := ms.PasswordRepo.DeletePassword(details.WebisteName)
			if deletePassword != nil {
				return nil, deletePassword
			}
		}
	}

	loginDetails, err := ms.LoginRepo.FindById(userId)
	if err != nil {
		return nil, err
	}

	deleteLoginEntryError := ms.LoginRepo.Delete(loginDetails.UserId)
	if deleteLoginEntryError != nil {
		return nil, deleteLoginEntryError
	}

	// Deleting user from user database
	deleteUserError := ms.UserRepo.Delete(userId, config.Algorithm)
	if deleteUserError != nil {
		return nil, deleteUserError
	}

	return &models.SuccessResponse{
		Message: "User deleted Successfully",
	}, nil
}

func (ms *MasterServiceImpl) UpdateUserTable(updatevalue string) error {
	db, err := db.NewDbRequest()
	if err != nil {
		return errors.New("error while creating a DB request: " + err.Error())
	}
	database, err := db.InitDB()
	log.Println("After database init")
	if err != nil {
		return errors.New("error in starting the DataBase: " + err.Error())
	}
	if updatevalue == "ASA" {
		userDetails, err := ms.UserRepo.FindAllRSAUser()
		if err != nil {
			return errors.New("error while finding all the users: " + err.Error())
		}
		err = database.Migrator().DropTable("user_tbl")
		if err != nil {
			return errors.New("error while dropping the table: " + err.Error())
		}
		err = database.AutoMigrate(&models.DBASAUser{})
		if err != nil {
			return errors.New("error while creating  the table ASA: " + err.Error())
		}
		for _, user := range userDetails {
			singleUser := &models.DBASAUser{}
			singleUser.CreatedAt = user.CreatedAt
			singleUser.CreatedBy = user.CreatedBy
			singleUser.Email = user.Email
			singleUser.UserId = user.UserId
			singleUser.IsMaster = user.IsMaster
			singleUser.MasterId = user.MasterId
			singleUser.Name = user.Name
			singleUser.Password = user.Password
			singleUser.SpecialKey = user.SpecialKey
			passwordKey, err := encryption.GenerateKey(32)
			if err != nil {
				return errors.New("error while generating ASA key for the user while migrating the algorithm: " + err.Error())
			}
			singleUser.PasswordKey = passwordKey
			err = ms.UserRepo.CreateASA(singleUser)
			if err != nil {
				return errors.New("error while creating the user and migrating: " + err.Error())
			}
		}
	} else {
		userDetails, err := ms.UserRepo.FindAllASAUser()
		if err != nil {
			return errors.New("error while finding all the users: " + err.Error())
		}
		err = database.Migrator().DropTable("user_tbl")
		if err != nil {
			return errors.New("error while dropping the table: " + err.Error())
		}
		err = database.AutoMigrate(&models.DBRSAUser{})
		if err != nil {
			return errors.New("error while creating  the table RSA: " + err.Error())
		}
		for _, user := range userDetails {
			singleUser := &models.DBRSAUser{}
			singleUser.CreatedAt = user.CreatedAt
			singleUser.CreatedBy = user.CreatedBy
			singleUser.Email = user.Email
			singleUser.UserId = user.UserId
			singleUser.IsMaster = user.IsMaster
			singleUser.MasterId = user.MasterId
			singleUser.Name = user.Name
			singleUser.Password = user.Password
			singleUser.SpecialKey = user.SpecialKey
			privateKey, publicKey, err := encryption.GenerateRSAKeys()
			if err != nil {
				return errors.New("error while generating the keys: " + err.Error())
			}
			publicKeyDb, err := encryption.PublicKeyToPEM(publicKey)
			if err != nil {
				return errors.New("error while converting the public key to the PEM: " + err.Error())
			}
			privateKeyDb := encryption.PrivateKeyToPEM(privateKey)
			singleUser.PrivateKey = privateKeyDb
			singleUser.PublicKey = publicKeyDb
			err = ms.UserRepo.CreateRSA(singleUser)
			if err != nil {
				return errors.New("error while creating the user while migrating the algorithm: " + err.Error())
			}
		}
	}
	return nil
}
