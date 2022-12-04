package chat

type connection struct {
	UserIdx uint `query:"user_idx" binding:"required"`
	RoomIdx uint `query:"room_idx" binding:"required"`
}

// chat message
type message struct {
	Type    uint8  `json:"type"`
	UserIdx uint   `json:"user_idx"`
	Message string `json:"message"`
}
