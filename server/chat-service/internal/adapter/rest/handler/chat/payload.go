package chat

type connection struct {
	UserIdx uint `form:"user_idx" binding:"required"`
	RoomIdx uint `form:"room_idx" binding:"required"`
}

// chat message
type message struct {
	Message string `json:"message"`
	Type    uint8  `json:"type"`
}
