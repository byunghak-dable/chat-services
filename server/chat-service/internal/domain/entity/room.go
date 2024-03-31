package entity

import (
	"time"
)

type Room struct {
	createdAt    time.Time
	updatedAt    time.Time
	id           string
	name         string
	participants []string
}

func NewRoom(id, name string, participants []string, createdAt, updatedAt time.Time) Room {
	return Room{
		id:           id,
		name:         name,
		participants: participants,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}
}

func (r *Room) Id() string {
	return r.id
}

func (r *Room) Name() string {
	return r.name
}

func (r *Room) Participants() []string {
	return r.participants
}

func (r *Room) CreatedAt() time.Time {
	return r.createdAt
}

func (r *Room) UpdatedAt() time.Time {
	return r.updatedAt
}
