package driven

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"messenger-service/internal/application/dto"
)

type KafkaProducer struct {
	*kafka.Producer
}

func NewKafkaProducer(config *kafka.ConfigMap) (*KafkaProducer, error) {
	producer, err := kafka.NewProducer(config)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{producer}, nil
}

func (producer *KafkaProducer) Produce(topic string, message *dto.MessageDto) error {
	kafkaMessage, err := producer.makeMessage(&topic, message)
	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	err = producer.Producer.Produce(kafkaMessage, nil)
	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	switch event := (<-producer.Events()).(type) {
	case *kafka.Message:
		if event.TopicPartition.Error != nil {
			return fmt.Errorf("delivery failed: %v", event.TopicPartition.Error)
		}
	default:
		return fmt.Errorf("unknown event: %v", event)
	}

	return nil
}

func (producer *KafkaProducer) makeMessage(topic *string, message *dto.MessageDto) (*kafka.Message, error) {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	return &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
		Value:          jsonMessage,
	}, nil
}

func (producer *KafkaProducer) Close() error {
	producer.Producer.Close()
	return nil
}
