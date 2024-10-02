package messaging

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	channel *amqp091.Channel
}

func NewProducer(conn *amqp091.Connection) (*Producer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &Producer{channel: ch}, nil
}

func (p *Producer) PublishMessage(exchange, routingKey string, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return p.channel.PublishWithContext(ctx, exchange, routingKey, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}
