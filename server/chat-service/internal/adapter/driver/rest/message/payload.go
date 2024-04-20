package message

type messagesQuery struct {
	RoomId string `uri:"room_id"`
	Cursor string `form:"cursor"`
	Limit  int64  `form:"limit"`
}
