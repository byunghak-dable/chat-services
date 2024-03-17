package room

import (
	"chat-service/internal/port/driven"
	"chat-service/internal/port/driver"
)

type Handler struct {
	logger driven.Logger
	app    driver.Room
}
