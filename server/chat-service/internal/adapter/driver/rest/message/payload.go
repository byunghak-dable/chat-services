package message

type messagesQuery struct {
	RoomId string `uri:"roomId"`
	Cursor string `uri:"cursor"`
	Limit  int64  `form:"limit" binding:"required"`
}
