package driving

import "github.com/widcraft/chat-service/internal/application/dto"

type ChatService interface {
	Join(client MessengerClient)
	Leave(client MessengerClient)
	SendMessage(messageDto *dto.MessageDto) error
	GetMessages(roomIdx uint) ([]dto.MessageDto, error)
}
