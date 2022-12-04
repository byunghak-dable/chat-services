package chat

type user struct {
	UserIdx  uint   `json:"user_idx" binding:"required"`
	Name     string `json:"name" binding:"required"`
	ImageUrl string `json:"image_url" binding:"required"`
}

// chat message
type message struct {
	Type    uint8  `json:"type"`
	UserIdx uint   `json:"user_idx"`
	Message string `json:"message"`
}
