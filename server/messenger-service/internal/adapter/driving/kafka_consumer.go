package driving

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"messenger-service/internal/port/driving"
)

type KafkaConsumer struct {
	*kafka.Consumer
	broadcaster driving.BroadcastServicePort
}

func NewKafkaConsumer(config *kafka.ConfigMap, broadcaster driving.BroadcastServicePort) (*KafkaConsumer, error) {
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{consumer, broadcaster}, nil
}

func (consumer *KafkaConsumer) Run() {
	//consumer.broadcaster.Broadcast()
}

func (consumer *KafkaConsumer) Close() error {
	return consumer.Close()
}
