package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/weed082/user-api/internal/adapter/rest/handler"
	"github.com/weed082/user-api/internal/adapter/rest/middleware"
	"github.com/weed082/user-api/internal/port"
)

type Rest struct {
	logger log.FieldLogger
	server *http.Server
}

func New(logger log.FieldLogger, userApp port.UserApp) *Rest {
	router := gin.Default()
	group := router.Group("/api/v1")
	middleware.NewErrorHandler(logger).Register(group)
	handler.NewUserHandler(logger, userApp).Register(group)

	return &Rest{
		logger: logger,
		server: &http.Server{
			Handler:      router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
}

func (r *Rest) Run(port string) {
	r.server.Addr = ":" + port
	err := r.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		r.logger.Errorf("rest serer error: %s", err)
	}
}

func (r *Rest) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.server.Shutdown(ctx); err != nil {
		r.logger.Errorf("shutting down rest server failed: %s", err)
	}
	r.logger.Info("shutting down rest server")
}
