package websocket

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/adapter/websocket/handler"
	"github.com/widcraft/chat-service/port"
)

type Websocket struct {
	logger  *log.Logger
	chatApp port.ChatApp
	server  *http.Server
}

func New(logger *log.Logger, chatApp port.ChatApp) *Websocket {
	router := gin.Default()
	group := router.Group("/chat/v1")
	handler.NewChatHandler(logger, chatApp).Register(group)
	return &Websocket{
		logger:  logger,
		chatApp: chatApp,
		server: &http.Server{
			Handler:      router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
}

func (ws *Websocket) Run(port string) {
	ws.server.Addr = ":" + port
	err := ws.server.ListenAndServe()
	if err != nil {
		ws.logger.Errorf("websocket server error: %s", err)
	}
}

func (ws *Websocket) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := ws.server.Shutdown(ctx); err != nil {
		ws.logger.Errorf("shutting down websocket server failed: %s", err)
	}
	ws.logger.Info("shutting down websocket server")
}
