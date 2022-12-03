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
	socketBufferSize := 1024
	upgrader := websocket.Upgrader{
		ReadBufferSize:  socketBufferSize,
		WriteBufferSize: socketBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			h.logger.Info(r)
			return true
		},
	}
	router.GET("ws", func(ctx *gin.Context) {
		ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			h.logger.Error("socket failed: %s", err)
			return
		}
		defer ws.Close()
		h.readMessage(ws)
	})
}

func (h *ChatHandler) readMessage(ws *websocket.Conn) error {
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			return err
		}
		h.logger.Info(messageType, message)
	}
}
