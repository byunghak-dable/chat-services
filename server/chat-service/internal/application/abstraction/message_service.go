package abstraction

import "github.com/widcraft/chat-service/internal/application/dto"

type MessageService interface {
	SaveMessage(message *dto.MessageDto) error
	GetMessages(roomIdx uint) ([]dto.MessageDto, error)
}
