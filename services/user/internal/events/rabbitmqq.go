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

func (r *RabbitMQBus) Consume(queueName string, handler func(amqp.Delivery)) error {
	_, err := r.ch.QueueDeclare(
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

	msgs, err := r.ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			handler(msg)
		}
	}()

	return nil
}
