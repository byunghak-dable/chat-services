package message

type messagesQuery struct {
	RoomId    string `uri:"room_id" binding:"required"`
	Cursor    string `form:"cursor"`
	UpdatedAt string `form:"updated_at"`
	Limit     int64  `form:"limit"`
}
