package messenger

import (
	"chat-service/internal/application/abstraction"
	"chat-service/internal/port/driver"
)

type LeaveUseCase struct {
	roomManager abstraction.RoomManager
}

func NewLeaveUseCase(roomManager abstraction.RoomManager) *LeaveUseCase {
	return &LeaveUseCase{roomManager}
}

func (u *LeaveUseCase) Handle(client driver.MessengerClient) {
	u.roomManager.Leave(client)
}
