package driving

import "messenger-service/internal/application/dto"

type Messenger interface {
	Join(client MessengerClient) error
	Leave(client MessengerClient) error
	SendMessage(message *dto.Message) error
}

type MessengerClient interface {
	GetRoomIdx() uint
	GetUserIdx() uint
	SendMessage(message *dto.Message) error
}
