package driving

import "messenger-service/internal/application/dto"

type MessengerServicePort interface {
	Join(client MessengerClientPort)
	Leave(client MessengerClientPort)
	SendMessage(message *dto.MessageDto) error
}

type MessengerClientPort interface {
	GetRoomIdx() uint
	GetUserIdx() uint
	SendMessage(message *dto.MessageDto) error
}
