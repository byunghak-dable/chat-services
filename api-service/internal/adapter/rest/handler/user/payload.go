package user

type register struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	ImageUrl string `json:"image_url" binding:"required,url"`
	Token    string `json:"token" binding:"required"`
}
