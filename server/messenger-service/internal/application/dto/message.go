package dto

import "time"

type Message struct {
	Id        string
	Message   string
	RoomIdx   uint
	UserIdx   uint
	CreatedAt time.Time
}
