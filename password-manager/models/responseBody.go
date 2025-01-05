package models

import "github.com/google/uuid"

type SuccessResponse struct {
	Message string `json:"message"`
}

type GetInfoResponse struct {
	Name       string `json:"name"`
	Algorithm  string `json:"algorithm"`
	Email      string `json:"email"`
	Count      int    `json:"count"`
	Plan       string `json:"plan"`
	CustomerId string `json:"customer_id"`
}

type GetUserListResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetUserInfoResponse struct {
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	WebsiteNames []string `json:"webiste_name"`
}

type LoginResponse struct {
	Message  string    `json:"message"`
	UserId   uuid.UUID `json:"user_id"`
	MasterId uuid.UUID `json:"master_id"`
}

type LoginResponseMaster struct {
	Message  string    `json:"message"`
	MasterId uuid.UUID `json:"master_id"`
}

type DeleteWebsiteResponse struct {
	Response string `json:"response"`
}

type ListWebsiteResponse struct {
	WebsiteName string `json:"website_name"`
	UserName    string `json:"user_name"`
}

type GetUserByEmailResponse struct {
	Email  string `json:"email"`
	UserId string `json:"user_id"`
}
