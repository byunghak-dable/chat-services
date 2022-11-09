package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weed082/chat-server/internal/adapter/driving/rest/dto"
)

type UserHandler struct {
	logger *log.Logger
}

func NewUserHandler(logger *log.Logger) *UserHandler {
	return &UserHandler{
		logger: logger,
	}
}

func (h UserHandler) Register(router *gin.RouterGroup) {
	router.POST("/user", h.register)
	router.GET("/user/:idx", h.getUserByIdx)
}

func (h UserHandler) register(c *gin.Context) {

}

func (h UserHandler) getUserByIdx(c *gin.Context) {
	var params dto.GetUserByIdxDto
	if err := c.ShouldBindUri(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, params)
}
