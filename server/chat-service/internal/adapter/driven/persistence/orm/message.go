package orm

import (
	"chat-service/internal/domain/entity"
	"time"
)

type Message struct {
	CreatedAt time.Time `bson:"created_at,"`
	UpdatedAt time.Time
	Id        string `bson:"_id,omitempty"`
	RoomId    string `bson:"room_id"`
	UserId    string `bson:"user_id"`
	Message   string
}

func FromMessage(entity entity.Message) Message {
	return Message{
		CreatedAt: entity.CreatedAt(),
		UpdatedAt: entity.UpdatedAt(),
		Id:        entity.Id(),
		RoomId:    entity.RoomId(),
		UserId:    entity.UserId(),
		Message:   entity.Message(),
	}
}

func (m *Message) ToDomain() entity.Message {
	return entity.NewMessage(m.Id, m.RoomId, m.UserId, m.Message, m.CreatedAt, m.UpdatedAt)
}
