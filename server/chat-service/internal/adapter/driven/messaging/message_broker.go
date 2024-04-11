package messaging

import (
	"chat-service/internal/adapter/driven/config"
	"chat-service/internal/application/dto"
	"chat-service/internal/port/driven"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type MessageBroker struct {
	logger      driven.Logger
	topic       string
	producer    *kafka.Producer
	consumer    *kafka.Consumer
	subscribers []driven.MessageSubscriber
}

func NewMessageBroker(configStore *config.Store, logger driven.Logger) (*MessageBroker, error) {
	configs := configStore.GetMessageBrokerConfig()
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": configs.Servers,
		"client.id":         configs.ClientId,
		"acks":              "all",
	})
	if err != nil {
		return nil, err
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": configs.Servers, "group.id": configs.GroupId,
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		return nil, err
	}

	return &MessageBroker{logger, configs.Topic, producer, consumer, []driven.MessageSubscriber{}}, nil
}

func (mb *MessageBroker) Subscribe(subscriber driven.MessageSubscriber) {
	mb.subscribers = append(mb.subscribers, subscriber)
}

func (mb *MessageBroker) Publish(message dto.Message) error {
	kafkaMessage, err := mb.makeMessage(message)
	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	if err := mb.producer.Produce(kafkaMessage, nil); err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	switch event := (<-mb.producer.Events()).(type) {
	case *kafka.Message:
		if event.TopicPartition.Error != nil {
			return fmt.Errorf("delivery failed: %v", event.TopicPartition.Error)
		}
	default:
		return fmt.Errorf("unknown event: %v", event)
	}

	return nil
}

func (mb *MessageBroker) Run(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	if err := mb.consumer.SubscribeTopics([]string{mb.topic}, nil); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			mb.close()
			mb.logger.Infoln("[MESSAGE_BROKER] successfully closed")
			return nil
		default:
			message, err := mb.consumer.ReadMessage(time.Second)

			if err == nil {
				mb.emitMessage(message.Value)
				continue
			}

			if err.(kafka.Error).IsTimeout() {
				continue
			}

			mb.logger.Errorf("[MESSAGE_BROKER] consumer read message failed: %s", err)
		}
	}
}

func (mb *MessageBroker) makeMessage(message dto.Message) (*kafka.Message, error) {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	return &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &mb.topic, Partition: kafka.PartitionAny},
		Value:          jsonMessage,
	}, nil
}

func (mb *MessageBroker) emitMessage(bytes []byte) {
	var message dto.Message

	if err := json.Unmarshal(bytes, &message); err != nil {
		mb.logger.Errorf("[MESSAGE_BROKER] parse message failed: %s", err)
		return
	}

	for _, subscriber := range mb.subscribers {
		subscriber.OnReceive(message)
	}
}

func (mb *MessageBroker) close() error {
	mb.producer.Flush(1000)
	mb.producer.Close()

	if err := mb.consumer.Close(); err != nil {
		return err
	}

	return nil
}
