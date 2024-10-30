package services

func LoginServiceRequest() (LoginService, error) {
	return &LoginServiceImpl{}, nil
}

type LoginServiceImpl struct{}

type LoginService interface{}
