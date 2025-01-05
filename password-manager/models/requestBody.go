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
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsMaster bool   `json:"is_master"`
}

type ListUserRequest struct {
	SpecialKey string `json:"special_key"`
}

type CreatePasswordRequest struct {
	WebisteName string `json:"website_name"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
}

type GetPasswordRequest struct {
	WebisteName string `json:"website_name"`
}

type ListWebsiteRequest struct {
	UserId   string `json:"user_id"`
	MasterId string `json:"master_id"`
}

type GetUserInfoRequest struct {
	UserId   string `json:"user_id"`
	MasterId string `json:"master_id"`
}

type MasterLoginRequest struct {
	SpecialKey string `json:"special_key"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserPassKeyUpdateRequest struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
