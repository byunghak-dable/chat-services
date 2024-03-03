package driving

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
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
	//consumer.broadcaster.Broadcast()
	err := consumer.SubscribeTopics(consumer.topics, nil)

	if err != nil {
		return err
	}

	for {
		event := consumer.Poll(100)

		switch e := event.(type) {
		case *kafka.Message:
			consumer.logger.Printf("event: %v", e)
		case kafka.Error:
			return e
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}
}

func (consumer *KafkaConsumer) Close() error {
	return consumer.Close()
}
