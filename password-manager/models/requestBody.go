package models

type EditKeyRequest struct {
	SpecialKey string `json:"special_key"`
	NewKey     string `json:"new_key"`
}

type GetInfoRequest struct {
	SpecialKey string `json:"special_key"`
}

type UpdateAlgorithmRequest struct {
	SpecialKey   string `json:"special_key"`
	NewAlgorithm string `json:"new_algorithm"`
}

type CreateUserRequest struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	MasterId   string `json:"master_id"`
	SpecialKey string `json:"special_key"`
}

type ListUserRequest struct {
	SpecialKey string `json:"special_key"`
}

type CreatePasswordRequest struct {
	UserId      string `json:"user_id"`
	WebisteName string `json:"website_name"`
	Password    string `json:"password"`
	MasterId    string `json:"master_id"`
}

type GetPasswordRequest struct {
	UserId      string `json:"user_id"`
	WebisteName string `json:"website_name"`
	MasterId    string `json:"master_id"`
}

type ListWebsiteRequest struct {
	UserId   string `json:"user_id"`
	MasterId string `json:"master_id"`
}
