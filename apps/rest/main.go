package main

import (
	"feature-flag/go/ports/kafka"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	kafkaConnector := kafka.NewKafkaConnection("kafka:9092", "never-mind")

	producer, err := kafkaConnector.Producer()

	topic := "test-topic"
	message := "Hello, Kafka!"

	router.GET("/", func(c *gin.Context) {
		err = producer.Publish(topic, message)

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
