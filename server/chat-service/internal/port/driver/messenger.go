package driver

import "chat-service/internal/application/dto"

type Messenger interface {
	Join(client MessengerClient)
	Leave(client MessengerClient)
	Send(message *dto.Message) error
}

type MessengerClient interface {
	Send(message *dto.Message) error
	RoomId() string
	UserId() string
}
