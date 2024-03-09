package driven

import "messenger-service/internal/application/dto"

type MessageProducer interface {
	Produce(message *dto.Message) error
}
