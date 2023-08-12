package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// Create a new Kafka consumer
	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatal("Error creating Kafka consumer:", err)
	}
	defer consumer.Close()

	// Define the Kafka topic
	topic := "test-topic"

	// Consume messages from the Kafka topic
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("Error consuming partition:", err)
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Received message: %s\n", string(msg.Value))
		}
	}
}
