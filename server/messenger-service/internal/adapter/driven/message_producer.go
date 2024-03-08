package driven

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"messenger-service/internal/application/dto"
	"messenger-service/internal/port/driven"
)

type MessageProducer struct {
	*kafka.Producer
	topic string
}

func NewMessageProducer(configStore driven.ConfigStore) (*MessageProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": configStore.GetKafkaServers(),
		"client.id":         configStore.GetKafkaClientId(),
		"acks":              "all",
	})
	if err != nil {
		return nil, err
	}

	return &MessageProducer{producer, configStore.GetKafkaChatTopic()}, nil
}

func (producer *MessageProducer) Produce(message *dto.MessageDto) error {
	kafkaMessage, err := producer.makeMessage(message)
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

func (producer *MessageProducer) makeMessage(message *dto.MessageDto) (*kafka.Message, error) {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	return &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &producer.topic, Partition: kafka.PartitionAny},
		Value:          jsonMessage,
	}, nil
}

func (producer *MessageProducer) Close() error {
	producer.Producer.Close()
	return nil
}
