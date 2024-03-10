package entity

import "time"

type Message struct {
	Id        string `bson:"_id"`
	Message   string
	RoomIdx   uint      `bson:"room_idx"`
	UserIdx   uint      `bson:"user_idx"`
	CreatedAt time.Time `bson:"created_at"`
}
