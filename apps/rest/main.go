package main

import (
	"feature-flag/go/domains/flags"
	"feature-flag/go/ports/kafka"
	"log"

	"github.com/gin-gonic/gin"
)

var topic = "feature-flag"

func main() {
	router := gin.Default()
	kafkaConnector := kafka.NewKafkaConnection("kafka:9092", "never-mind")

	producer, err := kafkaConnector.Producer()
	flagsDomain, err := flags.NewFlags("redis:6379")

	if err != nil {
		log.Fatal("Error creating flag domain")
	}

	if err != nil {
		log.Fatal("Error creating kafka producer")
	}

	router.POST("/v1/flag", postFlag(producer, flagsDomain))

	router.Run(":8080")
}

type PostFlagReq struct {
	Name string
}

func postFlag(producer kafka.ProducerInterface, flagsDomain flags.FlagsDomain) func(c *gin.Context) {
	return func(c *gin.Context) {
		var body PostFlagReq
		c.BindJSON(&body)

		flag, err := flagsDomain.Create(flags.CreateInput{
			Name: body.Name,
		})

		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})

			return
		}

		err = producer.Publish(topic, "Creating new flag")

		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})

			return
		}

		c.JSON(200, flag)
	}
}
