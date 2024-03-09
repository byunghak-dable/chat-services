package driving

import "messenger-service/internal/application/dto"

type Message interface {
	SaveMessage(message *dto.Message) error
	GetMessages(query *dto.MessagesQuery) ([]*dto.Message, error)
}
