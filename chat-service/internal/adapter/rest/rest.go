package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/adapter/rest/handler/chat"
	"github.com/widcraft/chat-service/internal/port"
)

type Rest struct {
	logger  *log.Logger
	server  *http.Server
	chatApp port.ChatApp
}

func New(logger *log.Logger, chatApp port.ChatApp) *Rest {
	router := gin.Default()
	group := router.Group("/api/v1")
	chat.New(logger, chatApp).Register(group)

	return &Rest{
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

func (ws *Rest) Run(port string) {
	ws.server.Addr = ":" + port
	err := ws.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		ws.logger.Errorf("websocket server error: %s", err)
	}
}

func (ws *Rest) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := ws.server.Shutdown(ctx); err != nil {
		ws.logger.Errorf("shutting down websocket server failed: %s", err)
	}
	ws.logger.Info("shutting down websocket server")
}
