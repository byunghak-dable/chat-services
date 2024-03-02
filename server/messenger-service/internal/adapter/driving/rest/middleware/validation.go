package middleware

import (
	"messenger-service/internal/port/driven"

	"github.com/gin-gonic/gin"
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
