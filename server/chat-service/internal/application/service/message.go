package service

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/application/mapper"
	"chat-service/internal/port/driven"
)

type Message struct {
	repository driven.MessageRepository
	mapper     *mapper.Message
}

func NewMessage(repository driven.MessageRepository, mapper *mapper.Message) *Message {
	return &Message{repository, mapper}
}

func (m *Message) Save(message *dto.Message) error {
	entity := m.mapper.ToEntity(message)

	return m.repository.Save(entity)
}

func (m *Message) GetSeveral(query *dto.MessagesQuery) ([]*dto.Message, error) {
	entities, err := m.repository.GetSeveral(query)
	if err != nil {
		return nil, err
	}

	return m.mapper.ToDtoList(entities), nil
}
