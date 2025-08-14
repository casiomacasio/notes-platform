package events

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

type RabbitMQBus struct {
	ch *amqp.Channel
}

func NewRabbitMQBus(conn *amqp.Connection) (*RabbitMQBus, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &RabbitMQBus{ch: ch}, nil
}

func (r *RabbitMQBus) Publish(queueName string, event any) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}
	_, err = r.ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return r.ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
