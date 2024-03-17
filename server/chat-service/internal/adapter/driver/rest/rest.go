package rest

import (
	driven2 "chat-service/internal/adapter/driven/config"
	"chat-service/internal/adapter/driver/rest/message"
	"chat-service/internal/adapter/driver/rest/messenger"
	"chat-service/internal/port/driven"
	"chat-service/internal/port/driver"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	logger driven.Logger
	server *http.Server
}

func New(configStore *driven2.Config, logger driven.Logger, messengerApp driver.Messenger, messageApp driver.Message) *Rest {
	router := gin.Default()
	group := router.Group("/api/v1")

	messenger.NewHandler(logger, messengerApp).Register(group)
	message.NewHandler(logger, messageApp).Register(group)

	return &Rest{
		logger: logger,
		server: &http.Server{
			Handler:      router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
			Addr:         ":" + configStore.GetRestPort(),
		},
	}
}

func (r *Rest) Run() error {
	err := r.server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (r *Rest) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return r.server.Shutdown(ctx)
}
