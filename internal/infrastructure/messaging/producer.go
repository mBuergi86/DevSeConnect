package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/mBuergi86/devseconnect/internal/domain/models"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type Producer struct {
	channel *amqp091.Channel
	logger  zerolog.Logger
}

func NewProducer(conn *amqp091.Connection) (*Producer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	err = ch.ExchangeDeclare(
		"user_events",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	return &Producer{channel: ch, logger: zerolog.New(os.Stdout).With().Timestamp().Str("component", "producer").Logger()}, nil
}

func (p *Producer) PublishMessage(exchange, routingKey string, message models.EventMessage) error {
	body, err := json.Marshal(message)
	if err != nil {
		p.logger.Error().Err(err).Msg("Failed to marshal complete message")
		return fmt.Errorf("failed to marshal complete message: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = p.channel.PublishWithContext(
		ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		p.logger.Error().Err(err).Msg("Failed to publish message")
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
