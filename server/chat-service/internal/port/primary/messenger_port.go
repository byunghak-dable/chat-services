package primary

import "github.com/widcraft/chat-service/internal/application/dto"

type MessengerClient interface {
	GetRoomIdx() uint
	GetUserIdx() uint
	SendMessage(message *dto.MessageDto) error
}
