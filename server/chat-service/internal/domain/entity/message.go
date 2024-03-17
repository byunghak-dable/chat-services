package entity

import (
	"time"
)

// Message TODO: need to handle timestamp to auto generate if empty
type Message struct {
	Id        string `bson:"_id,omitempty"`
	RoomId    string `bson:"room_id"`
	UserId    string `bson:"user_id"`
	Message   string
	CreatedAt time.Time `bson:"created_at,"`
	UpdatedAt time.Time
}

func (m *Message) SetId(id string) {
	m.Id = id
}
