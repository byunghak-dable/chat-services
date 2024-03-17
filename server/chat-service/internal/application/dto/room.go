package dto

import "time"

type Room struct {
	Id           string
	Name         string
	Participants []string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
