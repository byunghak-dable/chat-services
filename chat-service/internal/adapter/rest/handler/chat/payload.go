package chat

type connection struct {
	UserIdx uint32 `form:"user_idx" binding:"required"`
	RoomIdx uint32 `form:"room_idx" binding:"required"`
}

// chat message
type message struct {
	Type    uint8  `json:"type"`
	Message string `json:"message"`
}
