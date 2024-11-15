package models

import "github.com/google/uuid"

type SuccessResponse struct {
	Message string `json:"message"`
}

type GetInfoResponse struct {
	Name      string `json:"name"`
	Algorithm string `json:"algorithm"`
	Email     string `json:"email"`
	Count     int    `json:"count"`
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
	LoginId uuid.UUID `json:"login_id"`
}

type DeleteWebsiteResponse struct {
	Response string `json:"response"`
}
