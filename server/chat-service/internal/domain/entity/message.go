package entity

import (
	"time"
)

type Message struct {
	createdAt time.Time
	updatedAt time.Time
	id        string
	roomId    string
	userId    string
	message   string
}

func NewMessage(id, roomId, userId, message string, createdAt, updatedAt time.Time) Message {
	return Message{
		id:        id,
		roomId:    roomId,
		userId:    userId,
		message:   message,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (m *Message) SetId(id string) {
	m.id = id
}

func (m *Message) Id() string {
	return m.id
}

func (m *Message) RoomId() string {
	return m.roomId
}

func (m *Message) UserId() string {
	return m.userId
}

func (m *Message) Message() string {
	return m.message
}

func (m *Message) CreatedAt() time.Time {
	return m.createdAt
}

func (m *Message) UpdatedAt() time.Time {
	return m.updatedAt
}
