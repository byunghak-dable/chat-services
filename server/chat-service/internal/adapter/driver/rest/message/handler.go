package message

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/port/driven"
	"chat-service/internal/port/driver"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger           driven.Logger
	readMultiUseCase driver.ReadMultiMessageUseCase
}

func NewHandler(logger driven.Logger, getMultiUseCase driver.ReadMultiMessageUseCase) *Handler {
	return &Handler{logger, getMultiUseCase}
}

func (h *Handler) Register(router gin.IRoutes) {
	router.GET("/messages/room/:room_id", h.getMulti)
}

func (h *Handler) getMulti(ctx *gin.Context) {
	query := messagesQuery{Limit: 10}

	if err := ctx.ShouldBindUri(&query); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	messages, err := h.readMultiUseCase.Handle(dto.MessagesQuery{
		RoomId: query.RoomId,
		Cursor: query.Cursor,
		Limit:  query.Limit,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, messages)
}
