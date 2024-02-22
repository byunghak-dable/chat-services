package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/widcraft/chat-service/internal/port/driven"
)

type ValidationHandler struct {
	logger driven.Logger
	// TODO: need grpc client
}

func NewValidationHandler(logger driven.Logger) *ValidationHandler {
	return &ValidationHandler{
		logger: logger,
	}
}

func (h *ValidationHandler) Register(router *gin.RouterGroup) {
	router.Use(h.validateUser)
}

func (h *ValidationHandler) validateUser(c *gin.Context) {
	// TODO: need to request user service to check validation
	c.Next()
}
