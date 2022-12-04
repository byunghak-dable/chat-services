package chat

import (
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/port"
)

const (
	JOIN_ROOM_REQUEST = 1
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

func (h *Handler) Register(router *http.ServeMux) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			h.logger.Info("checking origin ", r)
			return true
		},
	}
	router.HandleFunc("/chat/v1/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			h.logger.Errorf("socket failed: %s", err)
			return
		}
		defer ws.Close()
		h.handleMessage(ws) // handle incoming message
	})
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

		switch msg.request {
		default:
			h.logger.Warn("unknown request")
		}
	}
}
