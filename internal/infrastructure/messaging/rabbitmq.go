package messaging

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbitMQ() (*amqp.Connection, error) {
	amqpServerURL := os.Getenv("RABBITMQ_URL")
	if amqpServerURL == "" {
		return nil, fmt.Errorf("RABBITMQ_URL is not set in the environment variables")
	}
	conn, err := amqp.Dial(amqpServerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	return conn, nil
}
