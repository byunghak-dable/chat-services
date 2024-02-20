package chat

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/widcraft/chat-service/internal/application/dto"
	"github.com/widcraft/chat-service/internal/port/primary"
	"github.com/widcraft/chat-service/internal/port/secondary"
)

type Handler struct {
	logger   secondary.Logger
	app      primary.MessageService
	upgrader *websocket.Upgrader
}

func New(logger secondary.Logger, app primary.MessageService) *Handler {
	return &Handler{
		logger: logger,
		app:    app,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(request *http.Request) bool {
				logger.Info("checking origin ", request)
				return true
			},
		},
	}
}

func (h *Handler) Register(router *gin.RouterGroup) {
	router.GET("chat", h.chat)
}

func (h *Handler) chat(ctx *gin.Context) {
	var param connection
	err := ctx.ShouldBindQuery(&param)
	if err != nil {
		h.logger.Error(err)
		return
	}

	conn, err := h.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		h.logger.Errorf("socket failed: %s", err)
		return
	}
	defer conn.Close()

	h.handleConnection(conn, param)
}

func (h *Handler) handleConnection(conn *websocket.Conn, param connection) {
	client := &client{
		websocketConn: conn,
		roomIdx:       param.RoomIdx,
		userIdx:       param.UserIdx,
	}

	h.app.Join(client)
	defer h.app.Leave(client)

	h.handleMessage(client)
}

func (h *Handler) handleMessage(client *client) {
	for {
		var msg *message
		err := client.websocketConn.ReadJSON(msg)

		if websocket.IsCloseError(err) || websocket.IsUnexpectedCloseError(err) {
			return
		}

		if err != nil {
			h.logger.Errorf("read message failed: %s", err)
		}

		h.sendMessge(client, msg)
	}
}

func (h *Handler) sendMessge(client *client, msg *message) {
	err := h.app.SendMessage(&dto.MessageDto{
		RoomIdx:  client.roomIdx,
		UserIdx:  client.userIdx,
		Name:     client.name,
		ImageUrl: client.imageUrl,
		Message:  msg.Message,
	})
	if err != nil {
		h.logger.Errorf("send message failed: %s", err)
	}
}
