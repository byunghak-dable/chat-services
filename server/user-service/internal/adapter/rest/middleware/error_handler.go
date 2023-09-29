package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/widcraft/user-service/pkg/logger"
)

type ErrorHandler struct {
	logger logger.Logger
}

func NewErrorHandler(logger logger.Logger) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
	}
}

func (h *ErrorHandler) Register(router *gin.RouterGroup) {
	router.Use(h.handleError)
}

func (h *ErrorHandler) handleError(c *gin.Context) {
	c.Next()
	errs := c.Errors.Errors()
	if len(errs) == 0 {
		return
	}

	timestamp := time.Now()
	h.logger.Errorf("[%s] : %s", timestamp, errs)
	c.AbortWithStatusJSON(c.Writer.Status(), gin.H{
		"messages":  errs,
		"timestamp": timestamp,
	})
}
