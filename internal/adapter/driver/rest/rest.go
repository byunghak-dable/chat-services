package rest

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/weed082/chat-server/internal/adapter/driver/rest/handler"
)

type Rest struct {
	logger *log.Logger
	server *http.Server
}

func New(logger *log.Logger) *Rest {
	router := gin.Default()
	group := router.Group("/api/v1")
	handler.NewUserHandler(logger).Register(group)

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
		r.logger.Fatalf("rest serer error: %s", err)
	}
}

func (r *Rest) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.server.Shutdown(ctx); err != nil {
		r.logger.Printf("shutting down rest server failed: %s", err)
	}
	r.logger.Println("shutting down rest server")
}
