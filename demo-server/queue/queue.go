package queue

import (
	"demo-server/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

const (
	maxRetries = 3               
	retryDelay = 2 * time.Second 
)

var (
	conn        *amqp.Connection
	channel     *amqp.Channel
	queueName   = "product_queue"
	rabbitmqURL = "amqp://guest:guest@rabbitmq:5672/"
)

type QueueService struct{}

type Queue interface {
	StartQueueAndPushProduct(productData *models.ProductDetailRequestBody) (string, error)
}

func NewQueue() (Queue, error) {
	log.Println("The queue request has been called.")
	return &QueueService{}, nil
}

func (q *QueueService) StartQueueAndPushProduct(productData *models.ProductDetailRequestBody) (string, error) {
	log.Println("INSIDE THE START QUEUE FUNCTION")

	// Retry connection
	for i := 0; i < maxRetries; i++ {
		if conn == nil {
			var err error
			conn, err = amqp.Dial(rabbitmqURL)
			if err == nil {
				log.Println("RabbitMQ connection established")
				break
			}
			log.Printf("Connection error: %v. Retrying in %v... (Attempt %d/%d)", err, retryDelay, i+1, maxRetries)
			time.Sleep(retryDelay)
			if i == maxRetries-1 {
				return "", errors.New("error while creating a connection: " + err.Error())
			}
		}
	}

	// Retry channel creation
	for i := 0; i < maxRetries; i++ {
		if channel == nil {
			var err error
			channel, err = conn.Channel()
			if err == nil {
				log.Println("RabbitMQ channel created")
				break
			}
			log.Printf("Channel error: %v. Retrying in %v... (Attempt %d/%d)", err, retryDelay, i+1, maxRetries)
			time.Sleep(retryDelay)
			if i == maxRetries-1 {
				return "", errors.New("error while creating a channel: " + err.Error())
			}
		}
	}

	// Retry queue declaration
	for i := 0; i < maxRetries; i++ {
		_, err := channel.QueueDeclare(
			queueName,
			false,
			false,
			false,
			false,
			nil,
		)
		if err == nil {
			log.Println("Queue declared")
			break
		}
		log.Printf("Queue declare error: %v. Retrying in %v... (Attempt %d/%d)", err, retryDelay, i+1, maxRetries)
		time.Sleep(retryDelay)
		if i == maxRetries-1 {
			return "", errors.New("error while declaring the queue: " + err.Error())
		}
	}

	// Convert product data to JSON
	messageBody, err := json.Marshal(productData)
	if err != nil {
		return "", errors.New("error while JSON marshalling: " + err.Error())
	}

	// Retry message publishing
	for i := 0; i < maxRetries; i++ {
		err = channel.Publish(
			"",
			queueName,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        messageBody,
			},
		)
		if err == nil {
			fmt.Printf("Message sent: %s\n", string(messageBody))
			log.Printf("Message sent to queue %s: %s", queueName, string(messageBody))
			return "success", nil
		}
		log.Printf("Publish error: %v. Retrying in %v... (Attempt %d/%d)", err, retryDelay, i+1, maxRetries)
		time.Sleep(retryDelay)
		if i == maxRetries-1 {
			return "", errors.New("error while publishing: " + err.Error())
		}
	}

	return "success", nil
}
