package rest

import (
	"context"
	"errors"
	"messenger-service/internal/adapter/driving/rest/chat"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	logger           driven.Logger
	server           *http.Server
	messengerService driving.Messenger
}

func New(configStore driven.ConfigStore, logger driven.Logger, messenger driving.Messenger) *Rest {
	router := gin.Default()
	group := router.Group("/api/v1")

	chat.NewHandler(logger, messenger).Register(group)

	return &Rest{
		logger:           logger,
		messengerService: messenger,
		server: &http.Server{
			Handler:      router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
			Addr:         ":" + configStore.GetRestPort(),
		},
	}
}

func (rest *Rest) Run() error {
	err := rest.server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (rest *Rest) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return rest.server.Shutdown(ctx)
}
