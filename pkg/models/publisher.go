package models

type Publisher struct {
	broker *Broker
	topic  string
}

func NewPublisher(broker *Broker, topic string) *Publisher {
	broker.CreateTopic(topic)
	return &Publisher{
		broker: broker,
		topic:  topic,
	}
}

func (p *Publisher) Publish(payload any) {
	msg := NewMessage(p.topic, payload)
	p.broker.Publish(p.topic, msg)
}
