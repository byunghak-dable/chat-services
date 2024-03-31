package messenger

import (
	"chat-service/internal/application/abstraction"
	"chat-service/internal/port/driver"
)

type JoinUseCase struct {
	roomManager abstraction.RoomManager
}

func NewJoinUseCase(roomManager abstraction.RoomManager) *JoinUseCase {
	return &JoinUseCase{roomManager}
}

func (u *JoinUseCase) Handle(client driver.MessengerClient) {
	u.roomManager.Join(client)
}
