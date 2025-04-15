package main

import (
	"log"
	"time"

	"github.com/keivanipchihagh/message-broker/pkg/models"
)

func main() {
	// Create a new broker
	broker := models.NewBroker()
	defer broker.Close()

	// Create and register publishers
	pub1 := models.NewPublisher("pub1", "news", broker)
	broker.AddPublisher(pub1)

	pub2 := models.NewPublisher("pub2", "sports", broker)
	broker.AddPublisher(pub2)

	// Create and register consumers
	con1 := models.NewConsumer("con1", "news", broker)
	broker.AddConsumer(con1)
	con1.Start()
	defer con1.Stop()

	con2 := models.NewConsumer("con2", "sports", broker)
	broker.AddConsumer(con2)
	con2.Start()
	defer con2.Stop()

	// Publish some messages
	pub1.Publish("Hello1")
	pub1.Publish("Hello2")
	pub2.Publish("Hello1")
	pub2.Publish("Hello2")

	// Give consumers time to process messages
	time.Sleep(100 * time.Millisecond)

	// Print registered consumers and publishers
	log.Println("Registered consumers per topic:")
	for topic, consumers := range broker.Consumers {
		var ids []string
		for _, c := range consumers {
			ids = append(ids, c.Id)
		}
		log.Printf("Topic %s: %v", topic, ids)
	}
}
