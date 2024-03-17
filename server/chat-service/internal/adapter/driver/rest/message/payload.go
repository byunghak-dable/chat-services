package message

type messagesQuery struct {
	Cursor string `form:"cursor"`
	Limit  int64  `form:"limit" binding:"required"`
}
