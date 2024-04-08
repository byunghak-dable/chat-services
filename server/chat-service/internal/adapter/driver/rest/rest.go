package rest

import (
	"chat-service/internal/adapter/driven/config"
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
	logger        driven.Logger
	server        *http.Server
	routerGroupV1 *gin.RouterGroup
}

func New(configStore *config.Config, logger driven.Logger) *Rest {
	handler := gin.Default()
	routeGroupV1 := handler.Group("/api/v1")

	return &Rest{
		logger: logger,
		server: &http.Server{
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
			Addr:         ":" + configStore.GetRestPort(),
		},
		routerGroupV1: routeGroupV1,
	}
}

func (r *Rest) RegisterMessenger(joinUseCase driver.MessengerJoinUseCase, leaveUseCase driver.MessengerLeaveUseCase, sendUseCase driver.MessengerSendUseCase) {
	messenger.NewHandler(r.logger, joinUseCase, leaveUseCase, sendUseCase).Register(r.routerGroupV1)
}

func (r *Rest) RegisterMessage(getMultiUseCase driver.GetMultiMessageUseCase) {
	message.NewHandler(r.logger, getMultiUseCase).Register(r.routerGroupV1)
}

func (r *Rest) Run() error {
	if err := r.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (r *Rest) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return r.server.Shutdown(ctx)
}
