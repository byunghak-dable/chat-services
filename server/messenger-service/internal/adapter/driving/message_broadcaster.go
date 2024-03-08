package driving

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"messenger-service/internal/application/dto"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
)

type MessageBroadcaster struct {
	logger driven.Logger
	*kafka.Consumer
	broadcaster driving.MessageBroadcaster
	topics      []string
}

func NewMessageBroadcaster(configStore driven.ConfigStore, logger driven.Logger, broadcaster driving.MessageBroadcaster) (*MessageBroadcaster, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": configStore.GetKafkaServers(),
		"group.id":          configStore.GetKafkaGroupId(), // TODO: need to add suffix for scale out
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		return nil, err
	}

	return &MessageBroadcaster{logger, consumer, broadcaster, []string{configStore.GetKafkaChatTopic()}}, nil
}

func (consumer *MessageBroadcaster) Run() error {
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

func (consumer *MessageBroadcaster) broadcast(bytes []byte) error {
	var message dto.MessageDto

	if err := json.Unmarshal(bytes, &message); err != nil {
		return err
	}

	return consumer.broadcaster.Broadcast(&message)
}

func (consumer *MessageBroadcaster) Close() error {
	return consumer.Close()
}
