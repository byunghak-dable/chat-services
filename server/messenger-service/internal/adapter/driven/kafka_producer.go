package driven

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"messenger-service/internal/port/driven"
)

type KafkaProducer[T any] struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaProducer[T any](configStore driven.ConfigStore) (*KafkaProducer[T], error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": configStore.GetKafkaServers(),
		"client.id":         configStore.GetKafkaClientId(),
		"acks":              "all",
	})
	if err != nil {
		return nil, err
	}

	return &KafkaProducer[T]{producer, configStore.GetKafkaChatTopic()}, nil
}

func (kp *KafkaProducer[T]) Produce(message *T) error {
	kafkaMessage, err := kp.makeMessage(message)
	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	err = kp.producer.Produce(kafkaMessage, nil)
	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	switch event := (<-kp.producer.Events()).(type) {
	case *kafka.Message:
		if event.TopicPartition.Error != nil {
			return fmt.Errorf("delivery failed: %v", event.TopicPartition.Error)
		}
	default:
		return fmt.Errorf("unknown event: %v", event)
	}

	return nil
}

func (kp *KafkaProducer[T]) makeMessage(message *T) (*kafka.Message, error) {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	return &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kp.topic, Partition: kafka.PartitionAny},
		Value:          jsonMessage,
	}, nil
}

func (kp *KafkaProducer[T]) Close() error {
	kp.producer.Close()
	return nil
}
