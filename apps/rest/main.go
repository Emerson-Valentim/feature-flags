package main

import (
	"feature-flag/go/domains/flags"
	"feature-flag/go/ports/kafka"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var topic = "feature-flag"

func main() {
	router := gin.Default()
	kafkaConnector := kafka.NewKafkaConnection("kafka:9092", "never-mind")

	producer, err := kafkaConnector.Producer()

	if err != nil {
		log.Fatal("Error creating kafka producer")
	}

	flagsDomain, err := flags.NewFlags("redis://redis:6379")

	if err != nil {
		log.Fatal("Error creating flag domain")
	}

	router.GET("/v1/flag/:id", getFlag(flagsDomain))
	router.POST("/v1/flag", postFlag(producer, flagsDomain))
	router.DELETE("/v1/flag/:id", deleteFlag(producer, flagsDomain))

	router.Run(":8080")
}

type PostFlagReq struct {
	Name string
}

func getFlag(flagsDomain flags.FlagsDomain) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		flag, err := flagsDomain.Get(id)

		if err != nil {
			msg := err.Error()

			if msg == "not found" {
				c.JSON(404, gin.H{
					"message": id + " not found",
				})
				return
			}

			c.JSON(500, gin.H{
				"message": "Internal server error",
			})

			return
		}

		c.JSON(200, flag)
	}
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
			fmt.Println("failed to publish event")
		}

		c.JSON(200, flag)
	}
}

func deleteFlag(producer kafka.ProducerInterface, flagsDomain flags.FlagsDomain) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := flagsDomain.Delete(id)

		if err != nil {
			msg := err.Error()

			if msg == "not found" {
				c.JSON(404, gin.H{
					"message": id + " not found",
				})
				return
			}

			c.JSON(500, gin.H{
				"message": "Internal server error",
			})

			return
		}

		err = producer.Publish(topic, "Deleting flag")

		if err != nil {
			fmt.Println("failed to publish event")
		}

		c.JSON(200, gin.H{
			"message": id + " deleted successfully",
		})
	}
}
