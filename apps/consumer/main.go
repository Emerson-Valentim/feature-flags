package main

import (
	"errors"
	"feature-flag/go/ports/kafka"
	"fmt"
	"log"
	"strings"
)

func main() {
	kafkaConnector := kafka.NewKafkaConnection("kafka:9092", "consumer")

	consumer, err := kafkaConnector.Consumer()

	if err != nil {
		log.Fatal("Error connecting:", err)
	}

	topic := "feature-flag"

	err = consumer.Start(topic, onMessage)

	if err != nil {
		log.Fatal("Error consuming partition:", err)
	}
}

func onMessage(message string) error {
	if strings.Contains(message, "fail") {
		return errors.New("Failed to consume the message")
	}

	fmt.Printf("Received message: %s\n", message)

	return nil
}
