package orm

import (
	"chat-service/internal/domain/entity"
	"time"
)

type Room struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Id           string
	Name         string
	Participants []string
}

func FromRoom(entity entity.Room) Room {
	return Room{
		Id:           entity.Id(),
		Name:         entity.Name(),
		Participants: entity.Participants(),
		CreatedAt:    entity.CreatedAt(),
		UpdatedAt:    entity.UpdatedAt(),
	}
}
