package dto

import "time"

type Room struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Id           string
	Name         string
	Participants []string
}
