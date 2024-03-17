package driver

import "chat-service/internal/application/dto"

type Room interface {
	Save(room *dto.Room) error
	GetRooms() []dto.Room
}
