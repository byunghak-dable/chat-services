package driver

import "chat-service/internal/application/dto"

type MessengerJoinUseCase interface {
	Handle(client MessengerClient)
}

type MessengerLeaveUseCase interface {
	Handle(client MessengerClient)
}

type MessengerSendUseCase interface {
	Handle(message dto.Message) error
}

type MessengerClient interface {
	Send(message dto.Message) error
	RoomId() string
	UserId() string
}
