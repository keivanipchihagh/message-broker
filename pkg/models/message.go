package models

import "time"

type Message struct {
	Topic     string
	Payload   any
	CreatedAt time.Time
}

func NewMessage(topic string, payload any) Message {
	return Message{
		Topic:     topic,
		Payload:   payload,
		CreatedAt: time.Now(),
	}
}
