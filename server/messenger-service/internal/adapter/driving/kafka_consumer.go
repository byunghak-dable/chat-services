package driving

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"messenger-service/internal/adapter/driven/config"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
)

type KafkaConsumer[T any] struct {
	logger      driven.Logger
	consumer    *kafka.Consumer
	broadcaster driving.Broadcaster[T]
	topics      []string
}

func NewKafkaConsumer[T any](configStore *config.Store, logger driven.Logger, broadcaster driving.Broadcaster[T]) (*KafkaConsumer[T], error) {
	configs := configStore.GetKafkaConsumerConfig()
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": configs.Servers,
		"group.id":          configs.GroupId,
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer[T]{logger, consumer, broadcaster, []string{configs.Topic}}, nil
}

func (kc *KafkaConsumer[T]) Run() error {
	if err := kc.consumer.SubscribeTopics(kc.topics, nil); err != nil {
		return err
	}

	for {
		message, err := kc.consumer.ReadMessage(-1)

		if err != nil {
			kc.logger.Errorf("kafka consumer read message failed: %s", err)
			continue
		}

		if err := kc.handleMessage(message.Value); err != nil {
			kc.logger.Errorf("kafka consumer handleMessagehandleMessage message failed: %s", err)
		}
	}
}

func (kc *KafkaConsumer[T]) handleMessage(bytes []byte) error {
	var message T

	if err := json.Unmarshal(bytes, &message); err != nil {
		return err
	}

	return kc.broadcaster.Broadcast(&message)
}

func (kc *KafkaConsumer[T]) Close() error {
	return kc.consumer.Close()
}
