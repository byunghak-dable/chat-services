package messenger

type connection struct {
	UserId string `form:"user_id" binding:"required"`
	RoomId string `form:"room_id" binding:"required"`
}
