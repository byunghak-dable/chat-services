package abstraction

import (
	"github.com/widcraft/chat-service/internal/application/dto"
	"github.com/widcraft/chat-service/internal/port/driving"
)

type MessengerService interface {
	Participate(client driving.MessengerClient)
	Quit(client driving.MessengerClient)
	SendMessage(message *dto.MessageDto) error
}
