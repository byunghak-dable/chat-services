package mapper

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/domain/entity"
)

type Room struct{}

func (r *Room) ToEntity(dto dto.Room) entity.Room {
	return entity.NewRoom(dto.Id, dto.Name, dto.Participants, dto.CreatedAt, dto.UpdatedAt)
}

func (r *Room) ToDto(entity entity.Room) dto.Room {
	return dto.Room{
		Id:           entity.Id(),
		Name:         entity.Name(),
		Participants: entity.Participants(),
		CreatedAt:    entity.CreatedAt(),
		UpdatedAt:    entity.UpdatedAt(),
	}
}
