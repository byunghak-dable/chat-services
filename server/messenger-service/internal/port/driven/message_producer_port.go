package driven

import "github.com/widcraft/messenger-service/internal/application/dto"

type MessageProducerPort interface {
	ProduceMessage(message *dto.MessageDto) error
}
