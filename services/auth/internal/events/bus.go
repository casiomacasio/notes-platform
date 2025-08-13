package events

type EventBus interface {
    Publish(topic string, event any) error
}