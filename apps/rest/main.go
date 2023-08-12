package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route
	router.GET("/", func(c *gin.Context) {
		producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)

		if err != nil {
			log.Fatal("Error creating Kafka producer:", err)
		}
		defer producer.Close()

		// Define the Kafka topic and message
		topic := "test-topic"
		message := "Hello, Kafka!"

		// Produce the message to the Kafka topic
		_, _, err = producer.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(message),
		})
		if err != nil {
			log.Fatal("Error sending message:", err)
		}

		fmt.Println("Message sent successfully!")

		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// Run the server on port 8080
	router.Run(":8080")
}
