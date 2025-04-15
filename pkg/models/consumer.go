package models

import "fmt"

// Consumer represents a message consumer
type Consumer struct {
	Id      string
	Topic   string
	broker  *Broker
	channel chan Message
}

// NewConsumer creates a new Consumer instance
func NewConsumer(id, topic string, broker *Broker) *Consumer {
	return &Consumer{
		Id:      id,
		Topic:   topic,
		broker:  broker,
		channel: make(chan Message, 10), // buffer to prevent blocking
	}
}

// Start begins consuming messages
func (c *Consumer) Start() {
	fmt.Printf("%s) started on toppic=%s\n", c.Id, c.Topic)
	go func() {
		for msg := range c.channel {
			c.handler(msg)
		}
	}()
}

// handler prints a message
func (c *Consumer) handler(msg Message) {
	fmt.Printf("%s) received: topic=%s, message=%v\n", c.Id, msg.Topic, msg)
}

// Close shuts down the channel and removes the consumer
func (c *Consumer) Stop() {
	if err := c.broker.RemoveConsumer(c); err != nil {
		fmt.Printf("%s) error during stop: %v\n", c.Id, err)
	}
	close(c.channel)
}
