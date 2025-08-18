package events

import "github.com/streadway/amqp"

type EventBus interface {
	Publish(topic string, event any) error
	Consume(queueName string, handler func(amqp.Delivery)) error
}
