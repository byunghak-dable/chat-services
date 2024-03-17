package entity

import (
	"time"
)

type Room struct {
	Id           string `bson:"_id"`
	Name         string
	Participants []string
	CreatedAt    time.Time `bson:"created_at"`
}
