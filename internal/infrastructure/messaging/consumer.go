package messaging

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp091.Connection
	channel   *amqp091.Channel
	queueName string
}

func NewConsumer(conn *amqp091.Connection, queueName string) (*Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn:      conn,
		channel:   ch,
		queueName: queueName,
	}, nil
}

func (c *Consumer) ConsumeMessages(handler func([]byte) error) error {
	q, err := c.channel.QueueDeclare(
		c.queueName, // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
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
			if err := handler(d.Body); err != nil {
				log.Printf("Error handling message: %v", err)
				d.Nack(false, true) // negative acknowledge, requeue the message
			} else {
				d.Ack(false) // acknowledge the message
			}
		}
	}()

	log.Printf("Consumer started for queue: %s", c.queueName)
	return nil
}
