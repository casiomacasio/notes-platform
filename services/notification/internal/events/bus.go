package events

import (
    "github.com/streadway/amqp"
)    

type EventBus interface {
    Consume(queueName string, handler func(amqp.Delivery))
}