package models

import (
	"errors"
	"slices"
	"sync"
)

type Broker struct {
	mu     sync.RWMutex
	topics map[string][]*Consumer
}

func NewBroker() *Broker {
	return &Broker{
		topics: make(map[string][]*Consumer),
	}
}

// CreateTopic initializes a new topic
func (b *Broker) CreateTopic(topic string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.topics[topic]; !exists {
		b.topics[topic] = make([]*Consumer, 0)
	}
}

// Subscribe adds a consumer to a topic
func (b *Broker) Subscribe(topic string, consumer *Consumer) error {

	b.mu.Lock()
	defer b.mu.Unlock()

	// Check if topic exists
	if _, exists := b.topics[topic]; !exists {
		return errors.New("topic does not exist")
	}

	// Check if consumer is already subscribed
	if slices.Contains(b.topics[topic], consumer) {
		return errors.New("consumer already subscribed to this topic")
	}

	b.topics[topic] = append(b.topics[topic], consumer)
	return nil
}

// Unsubscribe removes a consumer from a topic
func (b *Broker) Unsubscribe(topic string, consumer *Consumer) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Check if topic exists
	consumers, exists := b.topics[topic]
	if !exists {
		return errors.New("topic does not exist")
	}

	// Remove the consumer
	for i, c := range consumers {
		if c == consumer {
			b.topics[topic] = slices.Delete(consumers, i, i+1)
			return nil
		}
	}

	return errors.New("consumer not found in topic")
}

// Publish sends a message to all consumers of a topic
func (b *Broker) Publish(topic string, payload any) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	// Check if topic exists
	consumers, exists := b.topics[topic]
	if !exists {
		return errors.New("topic does not exist")
	}

	message := NewMessage(topic, payload)

	for _, consumer := range consumers {
		go func() {
			consumer.channel <- message
		}()
	}

	return nil
}
