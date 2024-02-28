package driven

import "messenger-service/internal/application/dto"

type MessageProducerPort interface {
	ProduceMessage(message *dto.MessageDto) error
}
