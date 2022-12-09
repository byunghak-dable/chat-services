package chat

type connection struct {
	UserIdx uint `form:"user_idx" binding:"required"`
	RoomIdx uint `form:"room_idx" binding:"required"`
}

// chat message
type message struct {
	Type    uint8  `json:"type"`
	RoomIdx uint   `json:"room_idx"`
	UserIdx uint   `json:"user_idx"`
	Message string `json:"message"`
}
