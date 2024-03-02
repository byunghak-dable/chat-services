package driven

import "messenger-service/internal/application/dto"

type MessageProducerPort interface {
	Produce(topic string, message *dto.MessageDto) error
}
