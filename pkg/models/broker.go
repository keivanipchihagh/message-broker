package models

import (
	"errors"
	"fmt"
	"slices"
	"sync"
)

type Broker struct {
	Publishers map[string][]*Publisher // topic -> []Publisher
	Consumers  map[string][]*Consumer  // topic -> []Consumer
	mu         sync.RWMutex
	channel    chan Message
}

func NewBroker() *Broker {
	b := &Broker{
		Consumers:  make(map[string][]*Consumer),
		Publishers: make(map[string][]*Publisher),
		channel:    make(chan Message),
	}
	go b.dispatchMessages()
	return b
}

func (b *Broker) AddPublisher(publisher *Publisher) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Check if already exists
	for _, p := range b.Publishers[publisher.Topic] {
		if p.Id == publisher.Id {
			return errors.New("publisher already exists")
		}
	}

	b.Publishers[publisher.Topic] = append(b.Publishers[publisher.Topic], publisher)
	return nil
}

func (b *Broker) AddConsumer(consumer *Consumer) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Check if already exists
	for _, c := range b.Consumers[consumer.Topic] {
		if c.Id == consumer.Id {
			return errors.New("consumer already exists")
		}
	}

	b.Consumers[consumer.Topic] = append(b.Consumers[consumer.Topic], consumer)
	return nil
}

func (b *Broker) RemovePublisher(publisher *Publisher) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	for i, p := range b.Publishers[publisher.Topic] {
		if p.Id == publisher.Id {
			b.Publishers[publisher.Topic] = slices.Delete(b.Publishers[publisher.Topic], i, i+1)
			return nil
		}
	}

	return errors.New("publisher not found")
}

func (b *Broker) RemoveConsumer(consumer *Consumer) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	for i, c := range b.Consumers[consumer.Topic] {
		if c.Id == consumer.Id {
			b.Consumers[consumer.Topic] = slices.Delete(b.Consumers[consumer.Topic], i, i+1)
			return nil
		}
	}

	return errors.New("consumer not found")
}

// Publish sends a message to the broker
func (b *Broker) Publish(msg Message) {
	b.channel <- msg
}

// Close shuts down the broker
func (b *Broker) Close() {
	close(b.channel)
}

func (b *Broker) dispatchMessages() {
	for msg := range b.channel {
		b.mu.RLock()
		consumers := make([]*Consumer, len(b.Consumers[msg.Topic]))
		copy(consumers, b.Consumers[msg.Topic])
		b.mu.RUnlock()

		for _, consumer := range consumers {
			select {
			case consumer.channel <- msg: // Non-blocking send
			default:
				fmt.Printf("consumer %s) channel full, message dropped\n", consumer.Id)
			}
		}
	}
}
