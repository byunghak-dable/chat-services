package dto

import "time"

type Message struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Id        string
	Message   string
	RoomId    string
	UserId    string
}

type MessagesQuery struct {
	RoomId string
	Cursor string
	Limit  int64
}
