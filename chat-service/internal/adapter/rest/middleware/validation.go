package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ValidationHandler struct {
	logger *log.Logger
	// TODO: need grpc client
}

func NewValidationHandler(logger *log.Logger) *ValidationHandler {
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
