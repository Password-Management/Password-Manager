package services

import (
	"demo-server/models"
	"demo-server/queue"
	"errors"
	"log"
)

type Product interface {
	GetProductDetails(request *models.ProductDetailRequestBody) (string, error)
}

type ProductType struct{}

func NewProductRequest() (Product, error) {
	return &ProductType{}, nil
}

func (pt *ProductType) GetProductDetails(request *models.ProductDetailRequestBody) (string, error) {
	log.Println(request)
	queue, err := queue.NewQueue()
	if err != nil {
		return "", errors.New("error while starting the queue: " + err.Error())
	}
	resp, err := queue.StartQueueAndPushProduct(request)
	if err != nil {
		log.Println("errro from queue: ", err)
		return "", err
	}
	return resp, nil
}
