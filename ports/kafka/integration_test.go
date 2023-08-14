package kafka

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKafkaConnection(t *testing.T) {
	conn := NewKafkaConnection("host", "TestNewKafkaConnection")
	assert.NotNil(t, conn.Host, "host")
}

func TestConsumerConnectionWithInvalidURL(t *testing.T) {
	conn := NewKafkaConnection("invalid-url", "TestConsumerConnectionWithInvalidURL")
	consumer, err := conn.Consumer()
	assert.Nil(t, consumer)
	assert.NotNil(t, err, "Error creating Kafka Consumer")
}

func TestConsumerConnection(t *testing.T) {
	conn := NewKafkaConnection("localhost:9092", "TestConsumerConnection")
	consumer, err := conn.Consumer()
	assert.NotNil(t, consumer)
	assert.Nil(t, err)
}

func TestProducerConnectionWithInvalidURL(t *testing.T) {
	conn := NewKafkaConnection("invalid-url", "TestProducerConnectionWithInvalidURL")
	producer, err := conn.Producer()
	assert.Nil(t, producer)
	assert.NotNil(t, err, "Error creating Kafka Producer")
}

func TestProducerConnection(t *testing.T) {
	conn := NewKafkaConnection("localhost:9092", "TestProducerConnection")
	producer, err := conn.Producer()
	assert.NotNil(t, producer)
	assert.Nil(t, err)
}

// func TestEventPublishAndConsume(t *testing.T) {
// 	topic := "topic-for-integration-test"

// 	conn := NewKafkaConnection("localhost:9092", "TestEventPublishAndConsume")

// 	producer, errP := conn.Producer()
// 	consumer, errC := conn.Consumer()

// 	assert.Nil(t, errP)
// 	assert.Nil(t, errC)

// 	consumed := false

// 	onMessage := func(message string) error {
// 		consumed = true

// 		return nil
// 	}

// 	err := consumer.Start(topic, onMessage)

// 	assert.Nil(t, err)

// 	producer.Publish(topic, "message")

// 	assert.Eventually(t, func() bool {
// 		return consumed
// 	}, 10*time.Second, 100*time.Millisecond)
// }
