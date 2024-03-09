package driving

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"messenger-service/internal/application/dto"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
)

type KafkaConsumer struct {
	logger      driven.Logger
	consumer    *kafka.Consumer
	broadcaster driving.MessageBroadcaster
	topics      []string
}

func NewKafkaConsumer(configStore driven.ConfigStore, logger driven.Logger, broadcaster driving.MessageBroadcaster) (*KafkaConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": configStore.GetKafkaServers(),
		"group.id":          configStore.GetKafkaGroupId(),
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{logger, consumer, broadcaster, []string{configStore.GetKafkaChatTopic()}}, nil
}

func (kc *KafkaConsumer) Run() error {
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

func (kc *KafkaConsumer) handleMessage(bytes []byte) error {
	var message dto.Message

	if err := json.Unmarshal(bytes, &message); err != nil {
		return err
	}

	return kc.broadcaster.Broadcast(&message)
}

func (kc *KafkaConsumer) Close() error {
	return kc.consumer.Close()
}
