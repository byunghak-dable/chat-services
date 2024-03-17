package driven

import (
	"chat-service/internal/domain/entity"
)

type RoomRepository interface {
	Save(room *entity.Room) error
	GetRooms() []entity.Room
}
