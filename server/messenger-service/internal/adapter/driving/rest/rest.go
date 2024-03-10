package rest

import (
	"context"
	"errors"
	"messenger-service/internal/adapter/driven/config"
	"messenger-service/internal/adapter/driving/rest/message"
	"messenger-service/internal/adapter/driving/rest/messenger"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	logger driven.Logger
	server *http.Server
}

func New(configStore *config.Store, logger driven.Logger, messengerApp driving.Messenger, messageApp driving.Message) *Rest {
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
