package dto

type GetUserByIdxDto struct {
	Idx uint `uri:"idx" binding:"required"`
}
