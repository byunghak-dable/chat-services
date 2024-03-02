package chat

import (
	"messenger-service/internal/application/dto"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	logger           driven.LoggerPort
	messengerService driving.MessengerServicePort
	upgrader         *websocket.Upgrader
}

func New(logger driven.LoggerPort, messengerService driving.MessengerServicePort) *Handler {
	return &Handler{
		logger:           logger,
		messengerService: messengerService,
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
		ctx.Status(http.StatusBadRequest)
		return
	}

	conn, upgradeErr := h.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if upgradeErr != nil {
		h.logger.Errorf("webSocket upgrade failed: %s", upgradeErr)
		ctx.Status(http.StatusBadRequest)
		return
	}

	defer func() {
		if err := conn.Close(); err != nil {
			h.logger.Errorf("webSocket close failed: %v", err)
		}
	}()

	connectionErr := h.handleConnection(conn, &param)

	if connectionErr != nil {
		h.logger.Errorf("webSocket connection handling failed: %s", connectionErr)
		ctx.Status(http.StatusInternalServerError)
	}
}

func (h *Handler) handleConnection(conn *websocket.Conn, param *connection) error {
	client := &client{
		websocketConn: conn,
		roomIdx:       param.RoomIdx,
		userIdx:       param.UserIdx,
	}

	err := h.messengerService.Join(client)

	if err != nil {
		return err
	}

	defer func() {
		err := h.messengerService.Leave(client)

		if err != nil {
			h.logger.Errorf("messenger leave failed: %v", err)
		}
	}()

	return h.handleMessage(client)
}

func (h *Handler) handleMessage(client *client) error {
	for {
		var msg message
		err := client.websocketConn.ReadJSON(&msg)

		if websocket.IsCloseError(err) || websocket.IsUnexpectedCloseError(err) {
			return err
		}

		if err != nil {
			h.logger.Errorf("read message failed: %s", err)
			continue
		}

		h.sendMessage(client, &msg)
	}
}

func (h *Handler) sendMessage(client *client, msg *message) {
	message := &dto.MessageDto{
		RoomIdx:  client.roomIdx,
		UserIdx:  client.userIdx,
		Name:     client.name,
		ImageUrl: client.imageUrl,
		Message:  msg.Message,
	}
	err := h.messengerService.SendMessage(message)

	if err != nil {
		h.logger.Errorf("send message failed: %s", err)
	}
}
