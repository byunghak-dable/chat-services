package room

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/application/mapper"
	"chat-service/internal/port/driven"
)

type CreateUsecase struct {
	repository driven.RoomRepository
	mapper     *mapper.Room
}

func NewCreateUsecase(repository driven.RoomRepository, mapper *mapper.Room) *CreateUsecase {
	return &CreateUsecase{repository, mapper}
}

func (u *CreateUsecase) Handle(room dto.Room) error {
	return u.repository.Save(u.mapper.ToEntity(room))
}
