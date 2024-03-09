package service

import (
	"messenger-service/internal/application/dto"
	"messenger-service/internal/application/mapper"
	"messenger-service/internal/port/driven/persistence"
)

type Message struct {
	repository persistence.MessageRepository
	mapper     *mapper.Message
}

func NewMessageStore(repository persistence.MessageRepository, mapper *mapper.Message) *Message {
	return &Message{repository, mapper}
}

func (ms *Message) SaveMessage(message *dto.Message) error {
	return ms.repository.SaveMessage(ms.mapper.ToEntity(message))
}

func (ms *Message) GetMessages(query *dto.MessagesQuery) ([]*dto.Message, error) {
	entities, err := ms.repository.GetMessages(query)
	if err != nil {
		return nil, err
	}

	return ms.mapper.ToDtoList(entities), nil
}
