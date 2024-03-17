package messenger

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/port/driven"
	"chat-service/internal/port/driver"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	logger   driven.Logger
	app      driver.Messenger
	upgrader *websocket.Upgrader
}

func NewHandler(logger driven.Logger, app driver.Messenger) *Handler {
	return &Handler{
		logger,
		app,
		&websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(request *http.Request) bool {
				return true
			},
		},
	}
}

func (h *Handler) Register(router *gin.RouterGroup) {
	router.GET("/chat", h.chat)
}

func (h *Handler) chat(ctx *gin.Context) {
	var param connection

	if err := ctx.ShouldBindQuery(&param); err != nil {
		h.logger.Errorf("connection params binding failed: %s", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	conn, err := h.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		h.logger.Errorf("upgrade connections failed: %s", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	defer func() {
		if err := conn.Close(); err != nil {
			h.logger.Errorf("webSocket close failed: %v", err)
		}
	}()

	if err := h.handleConnection(conn, &param); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) handleConnection(conn *websocket.Conn, param *connection) error {
	client := &client{
		conn:   conn,
		roomId: param.RoomId,
		userId: param.UserId,
	}

	h.app.Join(client)
	defer h.app.Leave(client)

	return h.handleMessage(client)
}

func (h *Handler) handleMessage(client *client) error {
	for {
		var message string

		err := client.conn.ReadJSON(&message)
		if websocket.IsCloseError(err) || websocket.IsUnexpectedCloseError(err) {
			return err
		}

		if err != nil {
			h.logger.Errorf("read message failed: %s", err)
			continue
		}

		h.sendMessage(client, message)
	}
}

func (h *Handler) sendMessage(client *client, message string) {
	messageDto := &dto.Message{
		RoomId:  client.roomId,
		UserId:  client.userId,
		Message: message,
	}

	if err := h.app.Send(messageDto); err != nil {
		h.logger.Errorf("send messageDto failed: %s", err)
	}
}
