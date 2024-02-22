package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/widcraft/messenger-service/internal/adapter/driving/rest/handler/chat"
	"github.com/widcraft/messenger-service/internal/port/driven"
	"github.com/widcraft/messenger-service/internal/port/driving"
)

type Rest struct {
	logger           driven.LoggerPort
	server           *http.Server
	messengerService driving.MessengerServicePort
}

func New(logger driven.LoggerPort, messaengerService driving.MessengerServicePort) *Rest {
	router := gin.Default()
	group := router.Group("/api/v1")

	chat.New(logger, messaengerService).Register(group)

	return &Rest{
		logger:           logger,
		messengerService: messaengerService,
		server: &http.Server{
			Handler:      router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
}

func (rest *Rest) Run(port string) {
	rest.server.Addr = ":" + port
	err := rest.server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		rest.logger.Errorf("websocket server error: %s", err)
	}
}

func (rest *Rest) OnExit() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return rest.server.Shutdown(ctx)
}
