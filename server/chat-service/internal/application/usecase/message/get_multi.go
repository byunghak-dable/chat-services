package message

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/application/mapper"
	"chat-service/internal/port/driven"
)

type GetMultiUseCase struct {
	repository driven.MessageRepository
	mapper     *mapper.Message
}

func NewGetMultiUseCase(service driven.MessageRepository, mapper *mapper.Message) *GetMultiUseCase {
	return &GetMultiUseCase{service, mapper}
}

func (u *GetMultiUseCase) Handle(query dto.MessagesQuery) ([]dto.Message, error) {
	entities, err := u.repository.GetMulti(query)
	if err != nil {
		return nil, err
	}

	return u.mapper.ToDtoList(entities), nil
}
