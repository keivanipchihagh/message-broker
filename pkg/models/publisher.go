package models

import (
	"fmt"
	"sync"
)

// Publisher represents a message publisher

type Publisher struct {
	Id     string
	Topic  string
	broker *Broker
	mu     sync.Mutex
}

// NewPublisher creates a new Publisher instance
func NewPublisher(id, topic string, broker *Broker) *Publisher {
	return &Publisher{
		Id:     id,
		Topic:  topic,
		broker: broker,
	}
}

// Publish sends a message to the broker
func (p *Publisher) Publish(payload any) {
	p.mu.Lock()
	defer p.mu.Unlock()

	msg := NewMessage(p.Topic, payload)
	p.broker.Publish(msg)
	fmt.Printf("%s) sent: topic=%s, message=%v\n", p.Id, msg.Topic, msg)
}
