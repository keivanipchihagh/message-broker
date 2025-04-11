package main

import (
	"fmt"
	"time"

	"github.com/keivanipchihagh/message-broker/pkg/models"
)

func main() {
	// Create the broker
	broker := models.NewBroker()

	// Create topics
	broker.CreateTopic("news")
	broker.CreateTopic("sports")

	// Create consumers
	newsConsumer1 := models.NewConsumer("news")
	newsConsumer2 := models.NewConsumer("news")
	sportsConsumer := models.NewConsumer("sports")

	// Subscribe consumers to topics
	if err := broker.Subscribe("news", newsConsumer1); err != nil {
		panic(err)
	}
	defer broker.Unsubscribe("news", newsConsumer1)

	if err := broker.Subscribe("news", newsConsumer2); err != nil {
		panic(err)
	}
	defer broker.Unsubscribe("news", newsConsumer2)

	if err := broker.Subscribe("sports", sportsConsumer); err != nil {
		panic(err)
	}
	defer broker.Unsubscribe("sports", sportsConsumer)

	// Start consuming messages
	newsConsumer1.Start()
	defer newsConsumer1.Close()

	newsConsumer2.Start()
	defer newsConsumer2.Close()

	sportsConsumer.Start()
	defer sportsConsumer.Close()

	// Publish some messages
	go func() {
		for i := range 1 {
			broker.Publish("news", fmt.Sprintf("News update %d", i))
			broker.Publish("sports", fmt.Sprintf("Sports update %d", i))
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Let the system run for a while
	time.Sleep(3 * time.Second)
}
