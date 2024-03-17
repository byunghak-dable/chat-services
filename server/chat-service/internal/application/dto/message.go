package dto

import "time"

type Message struct {
	Id        string
	Message   string
	RoomId    string
	UserId    string
	UpdatedAt time.Time
}

type MessagesQuery struct {
	RoomId string
	Cursor string
	Limit  int64
}
