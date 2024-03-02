package driving

import "messenger-service/internal/application/dto"

type MessengerServicePort interface {
	Join(client MessengerClientPort) error
	Leave(client MessengerClientPort) error
	SendMessage(message *dto.MessageDto) error
}

type BroadcastServicePort interface {
	Broadcast(messageDto *dto.MessageDto) error
}

type MessengerClientPort interface {
	GetRoomIdx() uint
	GetUserIdx() uint
	SendMessage(message *dto.MessageDto) error
}
