package models

import "fmt"

type Consumer struct {
	Id      string
	topic   string
	channel chan Message
}

// NewConsumer creates a new consumer
func NewConsumer(topic string) *Consumer {
	return &Consumer{
		topic:   topic,
		channel: make(chan Message),
	}
}

// Close stops the consumer
func (c *Consumer) Close() {
	close(c.channel)
}

func (c *Consumer) Start() {
	go func() {
		for {
			select {
			case message, ok := <-c.channel:
				if !ok {
					return
				}
				fmt.Println(message)
			}
		}
	}()
}
