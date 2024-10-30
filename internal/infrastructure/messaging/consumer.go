package messaging

import (
	"fmt"
	"os"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type Consumer struct {
	conn      *amqp091.Connection
	channel   *amqp091.Channel
	queueName string
	logger    zerolog.Logger
}

func NewConsumer(conn *amqp091.Connection, queueName string) (*Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"user_events", // Exchange name
		"direct",      // Exchange type
		true,          // Durable
		false,         // Auto-deleted
		false,         // Internal
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	q, err := ch.QueueDeclare(
		queueName, // Queue name
		true,      // Durable
		false,     // Auto-deleted
		false,     // Exclusive
		false,     // No-wait
		amqp091.Table{
			"x-dead-letter-exchange":    "dlx",
			"x-dead-letter-routing-key": "dlq",
		}, // Arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	err = ch.QueueBind(
		q.Name,        // Queue name
		queueName,     // Routing key
		"user_events", // Exchange
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	return &Consumer{
		conn:      conn,
		channel:   ch,
		queueName: queueName,
		logger:    zerolog.New(os.Stdout).With().Timestamp().Str("component", "consumer").Logger(),
	}, nil
}

func (c *Consumer) handleMessage(handler func([]byte) error, delivery amqp091.Delivery) {
	if err := handler(delivery.Body); err != nil {
		c.logger.Error().Err(err).Msg("Failed to handle message")
		delivery.Nack(false, true)
		return
	}

	delivery.Ack(false)
	c.logger.Info().Msg("Message successfully processed")
}

func (c *Consumer) ConsumeMessages(handler func([]byte) error) error {
	msgs, err := c.channel.Consume(
		c.queueName, // queue
		"",          // consumer
		false,       // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to register consumer")
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	go func() {
		for d := range msgs {
			c.handleMessage(handler, d)
		}
	}()

	c.logger.Info().Msgf("Consumer started for queue: %s", c.queueName)
	return nil
}
