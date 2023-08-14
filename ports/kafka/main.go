package kafka

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type KafkaConnection struct {
	Host    string
	GroupId string
}

func NewKafkaConnection(host string, groupId string) KafkaConnection {
	return KafkaConnection{
		Host:    host,
		GroupId: groupId,
	}
}

type SaramaProducerAdapter struct {
	sarama.SyncProducer
}

type SaramaConsumerAdapter struct {
	sarama.ConsumerGroup
}

type OnMessage func(message string) error

type ConsumerWrapper struct {
	ready     chan bool
	onMessage OnMessage
}

func (consumer *ConsumerWrapper) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *ConsumerWrapper) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *ConsumerWrapper) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			err := consumer.onMessage(string(message.Value))

			if err != nil {
				return err
			}

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}

func (pa *SaramaProducerAdapter) Publish(topic string, message string) error {
	_, _, err := pa.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	})

	return err
}

func (ca *SaramaConsumerAdapter) Start(topic string, onMessage OnMessage) error {
	ctx, _ := context.WithCancel(context.Background())

	wrappedOnMessage := ConsumerWrapper{
		ready:     make(chan bool),
		onMessage: onMessage,
	}

	err := ca.Consume(ctx, []string{topic}, &wrappedOnMessage)

	if err != nil {
		return err
	}

	return nil
}

type ProducerInterface interface {
	Publish(topic string, message string) error
}

type ConsumerInterface interface {
	Start(topic string, onMessage OnMessage) error
}

func (conn *KafkaConnection) Consumer() (ConsumerInterface, error) {
	config := sarama.NewConfig()
	config.Admin.Retry.Max = 3
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}

	consumer, err := sarama.NewConsumerGroup([]string{conn.Host}, conn.GroupId, config)

	if err != nil {
		msg := "Error creating Kafka consumer"

		fmt.Printf(msg, err)

		return nil, errors.New(msg)
	}

	return &SaramaConsumerAdapter{consumer}, nil
}

func (conn *KafkaConnection) Producer() (ProducerInterface, error) {
	producer, err := sarama.NewSyncProducer([]string{conn.Host}, nil)

	if err != nil {
		msg := "Error creating Kafka producer"

		fmt.Printf(msg, err)

		return nil, errors.New(msg)
	}

	return &SaramaProducerAdapter{producer}, nil
}
