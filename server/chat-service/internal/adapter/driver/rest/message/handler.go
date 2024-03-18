package message

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/port/driven"
	"chat-service/internal/port/driver"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	logger driven.Logger
	app    driver.Message
}

func NewHandler(logger driven.Logger, app driver.Message) *Handler {
	return &Handler{logger, app}
}

func (h *Handler) Register(router *gin.RouterGroup) {
	router.GET("/messages/room/:roomId/cursor/:cursor", h.getMulti)
}

func (h *Handler) getMulti(ctx *gin.Context) {
	var query messagesQuery

	if err := ctx.ShouldBindUri(&query); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	messages, err := h.app.GetMulti(dto.MessagesQuery{
		RoomId: query.RoomId,
		Cursor: query.Cursor,
		Limit:  query.Limit,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, messages)
}
