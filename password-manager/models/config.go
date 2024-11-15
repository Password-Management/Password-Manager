package models

type Config struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	ProductType string `json:"product_type"`
	Algorithm   string `json:"algorithm"`
}
