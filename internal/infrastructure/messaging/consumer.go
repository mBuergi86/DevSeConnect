package messaging

import (
	"encoding/json"
	"log"

	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
	"github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn     *amqp091.Connection
	channel  *amqp091.Channel
	userRepo repository.UserRepository
}

func NewConsumer(conn *amqp091.Connection, userRepo repository.UserRepository) (*Consumer, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn:     conn,
		channel:  channel,
		userRepo: userRepo,
	}, nil
}

func (c *Consumer) Start() error {
	q, err := c.channel.QueueDeclare(
		"user_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return err
	}

	err = c.channel.QueueBind(
		q.Name,        // queue name
		"user_queue",  // routing key
		"user_events", // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := c.channel.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			var event map[string]interface{}
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Printf("Error unmarshaling event: %v", err)
				d.Nack(false, false)
				continue
			}

			eventType, _ := event["type"].(string)
			userData, _ := event["user"].(map[string]interface{})

			var user entity.User
			userJSON, _ := json.Marshal(userData)
			if err := json.Unmarshal(userJSON, &user); err != nil {
				log.Printf("Error unmarshaling user data: %v", err)
				d.Nack(false, false)
				continue
			}

			switch eventType {
			case "user_created":
				log.Printf("Processing user_created event for user %s", user.UserID)
				// Handle user creation event
			case "user_updated":
				log.Printf("Processing user_updated event for user %s", user.UserID)
				// Handle user update event
			case "user_deleted":
				log.Printf("Processing user_deleted event for user %s", user.UserID)
				// Handle user deletion event
			default:
				log.Printf("Unknown event type: %s", eventType)
			}

			d.Ack(false)
		}
	}()

	log.Printf("RabbitMQ Consumer started. Waiting for messages...")
	return nil
}
