package message

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/application/mapper"
	"chat-service/internal/port/driven"
)

type ReadMultiUseCase struct {
	repository driven.MessageRepository
	mapper     *mapper.Message
}

func NewReadMultiUseCase(service driven.MessageRepository, mapper *mapper.Message) *ReadMultiUseCase {
	return &ReadMultiUseCase{service, mapper}
}

func (u *ReadMultiUseCase) Handle(query dto.MessagesQuery) ([]dto.Message, error) {
	entities, err := u.repository.GetMulti(query)
	if err != nil {
		return nil, err
	}

	return u.mapper.ToDtoList(entities), nil
}
