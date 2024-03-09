package service

import (
	"messenger-service/internal/application/dto"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
)

type Messenger struct {
	producer     driven.MessageProducer
	messageStore driving.MessageStore
	roomManager  *RoomManager
}

func NewMessenger(producer driven.MessageProducer, messageStore driving.MessageStore, roomManager *RoomManager) *Messenger {
	return &Messenger{producer, messageStore, roomManager}
}

func (m *Messenger) Join(client driving.MessengerClient) error {
	return m.roomManager.Join(client)
}

func (m *Messenger) Leave(client driving.MessengerClient) error {
	return m.roomManager.Leave(client)
}

func (m *Messenger) Broadcast(message *dto.Message) error {
	return m.roomManager.Broadcast(message)
}

func (m *Messenger) SendMessage(message *dto.Message) error {
	if err := m.messageStore.SaveMessage(message); err != nil {
		return err
	}
	return m.producer.Produce(message)
}
