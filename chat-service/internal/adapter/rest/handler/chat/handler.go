package chat

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/port"
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
	router.GET("chat", h.makeChatHandler())
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
		var param connection
		err := ctx.ShouldBindQuery(&param)
		if err != nil {
			h.logger.Error(err)
			return
		}

		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			h.logger.Errorf("socket failed: %s", err)
			return
		}
		defer conn.Close()

		client := &client{userIdx: param.UserIdx, conn: conn}
		if err = h.app.Connect(param.RoomIdx, client); err != nil {
			h.logger.Errorf("connect client failed", err)
			return
		}

		h.handleConnection(conn)

		if err = h.app.Disconnect(param.RoomIdx, client); err != nil {
			h.logger.Errorf("disconnect client failed: %s", err)
		}
	}
}

func (h *Handler) handleConnection(conn *websocket.Conn) {
	for {
		var msg message
		err := conn.ReadJSON(&msg)
		if websocket.IsCloseError(err) || websocket.IsUnexpectedCloseError(err) {
			h.logger.Info("websocket connection closed")
			return
		}
		if err != nil {
			h.logger.Error(err)
			continue
		}
		h.app.SendMessge(&dto.MessageDto{
			RoomIdx: msg.RoomIdx,
			UserIdx: msg.UserIdx,
			Message: msg.Message,
		})
	}
}
