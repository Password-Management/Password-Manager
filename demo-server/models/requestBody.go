package models

type ProductDetailRequestBody struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	ProductType string `json:"product_type"`
	Algorithm   string `json:"algorithm"`
}
