package driven

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"messenger-service/internal/adapter/driven/config"
)

type KafkaProducer[T any] struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaProducer[T any](configStore *config.Store) (*KafkaProducer[T], error) {
	configs := configStore.GetKafkaProducerConfig()
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": configs.Servers,
		"client.id":         configs.ClientId,
		"acks":              "all",
	})
	if err != nil {
		return nil, err
	}

	return &KafkaProducer[T]{producer, configs.Topic}, nil
}

func (kp *KafkaProducer[T]) Produce(message *T) error {
	kafkaMessage, err := kp.makeMessage(message)
	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	if err := kp.producer.Produce(kafkaMessage, nil); err != nil {
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
