package models

import "time"

type Message struct {
	Topic     string    `json:"topic"`
	Payload   any       `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

func NewMessage(topic string, payload any) Message {
	return Message{
		Topic:     topic,
		Payload:   payload,
		CreatedAt: time.Now(),
	}
}
