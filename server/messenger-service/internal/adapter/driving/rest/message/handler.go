package message

import (
	"github.com/gin-gonic/gin"
	"messenger-service/internal/application/dto"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
	"net/http"
)

type Handler struct {
	logger driven.Logger
	app    driving.Message
}

func NewHandler(logger driven.Logger, app driving.Message) *Handler {
	return &Handler{logger, app}
}

func (h *Handler) Register(router *gin.RouterGroup) {
	router.GET("/messages", h.getMessages)
}

func (h *Handler) getMessages(ctx *gin.Context) {
	messages, err := h.app.GetMessages(&dto.MessagesQuery{})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, messages)
}
