package models

type SuccessResponse struct {
	Message string `json:"message"`
}

type GetInfoResponse struct {
	Name      string `json:"name"`
	Algorithm string `json:"algorithm"`
	Email     string `json:"email"`
	Count     int `json:"count"`
}

type GetUserListResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
