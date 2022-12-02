package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/port"
)

type ChatHandler struct {
	logger *log.Logger
	app    port.ChatApp
}

func NewChatHandler(logger *log.Logger, app port.ChatApp) *ChatHandler {
	return &ChatHandler{
		logger: logger,
		app:    app,
	}
}

func (h *ChatHandler) Register(router *gin.RouterGroup) {
	upgrader := websocket.Upgrader{
		// TODO; check for more info
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	router.GET("ws", func(ctx *gin.Context) {
		socket, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			h.logger.Error("socket failed: %s", err)
			return
		}
		defer socket.Close()
	})
}
