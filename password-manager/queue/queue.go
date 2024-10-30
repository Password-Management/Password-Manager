package queue

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"password-manager/helpers"
	"password-manager/models"
	"strings"

	"github.com/streadway/amqp"
)



const URL = "amqp://guest:guest@"

func getRabbitMQContainerIP() (string, error) {
	// Execute Docker inspect command to get the IP address of the RabbitMQ container
	cmd := exec.Command("docker", "inspect", "-f", "{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}", "rabbitmq")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get RabbitMQ container IP: %v", err)
	}

	ipAddress := strings.TrimSpace(string(output))
	return ipAddress, nil
}

func QueueConsumer() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/") // Adjust the URL based on your setup
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	// Create a channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer channel.Close()

	// Declare the queue to ensure it exists
	queueName := "product_queue" // Change this to your actual queue name
	_, err = channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Start consuming messages from the queue
	msgs, err := channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	log.Println("Waiting for messages. To exit press CTRL+C")

	// Consume messages
	for msg := range msgs {
		var productDetail *models.Config
		err := json.Unmarshal(msg.Body, &productDetail)
		if err != nil {
			log.Printf("Error unmarshalling message: %s", err)
			continue
		}
		response, err := helpers.CreateConfig(productDetail)
		if err != nil {
			log.Printf("error while wrting the config: %s", err)
			continue
		}
		log.Println("The response: ", response)
		// Print the received message
		log.Printf("Message received and updated in the config.yml")
	}
}
