package chat

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/port"
)

type Handler struct {
	logger *log.Logger
	app    port.ChatApp
}

func New(logger *log.Logger, app port.ChatApp) *Handler {
	return &Handler{
		logger: logger,
		app:    app,
	}
}

func (h *Handler) Register(router *gin.RouterGroup) {
	router.POST("chat", h.makeChatHandler())
}

func (h *Handler) makeChatHandler() gin.HandlerFunc {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			h.logger.Info("checking origin ", r)
			return true
		},
	}
	return func(ctx *gin.Context) {
		ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			h.logger.Errorf("socket failed: %s", err)
			return
		}
		defer ws.Close()
		h.handleMessage(ws)
	}
}

func (h *Handler) handleMessage(ws *websocket.Conn) {
	for {
		var msg *message
		err := ws.ReadJSON(msg)

		if websocket.IsCloseError(err) || websocket.IsUnexpectedCloseError(err) {
			h.logger.Info("connection closed")
			return
		}
		if err != nil {
			h.logger.Error(err)
			continue
		}
	}
}
