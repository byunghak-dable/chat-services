package abstraction

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/port/driver"
)

type RoomManager interface {
	Join(client driver.MessengerClient)
	Leave(client driver.MessengerClient)
	Broadcast(message dto.Message) error
}
