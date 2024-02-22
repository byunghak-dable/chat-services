package abstraction

import "github.com/widcraft/chat-service/internal/application/dto"

type MessageService interface {
	SaveMessage(message *dto.MessageDto) error          // NOTE: should move to kafka message consumer pipeline
	GetMessages(roomIdx uint) ([]dto.MessageDto, error) // TODO: cursor pagination functionality
}
