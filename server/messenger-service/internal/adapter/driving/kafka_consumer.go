package driving

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"messenger-service/internal/application/dto"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
)

type KafkaConsumer struct {
	logger driven.LoggerPort
	*kafka.Consumer
	broadcaster driving.BroadcastServicePort
	topics      []string
}

func NewKafkaConsumer(logger driven.LoggerPort, config *kafka.ConfigMap, broadcaster driving.BroadcastServicePort, topics []string) (*KafkaConsumer, error) {
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{logger, consumer, broadcaster, topics}, nil
}

func (consumer *KafkaConsumer) Run() error {
	if err := consumer.SubscribeTopics(consumer.topics, nil); err != nil {
		return err
	}

	for {
		message, err := consumer.ReadMessage(-1)

		if err != nil {
			consumer.logger.Errorf("kafka consumer read message failed: %s", err)
			continue
		}

		if err := consumer.broadcast(message.Value); err != nil {
			consumer.logger.Errorf("kafka consumer broadcast message failed: %s", err)
		}
	}
}

func (consumer *KafkaConsumer) broadcast(bytes []byte) error {
	var message dto.MessageDto

	if err := json.Unmarshal(bytes, &message); err != nil {
		return err
	}

	return consumer.broadcaster.Broadcast(&message)
}

func (consumer *KafkaConsumer) Close() error {
	return consumer.Close()
}
