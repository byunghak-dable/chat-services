package driven

import "messenger-service/internal/application/dto"

type MessageProducerPort interface {
	Produce(message *dto.MessageDto) error
}
