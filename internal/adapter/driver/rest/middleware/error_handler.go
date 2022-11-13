package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type ErrorHandler struct {
	logger *log.Logger
}

func NewErrorHandler(logger *log.Logger) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
	}
}

func (h ErrorHandler) Register(router *gin.RouterGroup) {
	router.Use(h.handleError)
}

func (h ErrorHandler) handleError(c *gin.Context) {
	c.Next()
	errs := c.Errors.Errors()
	timestamp := time.Now()

	h.logger.Printf("[%s] : %s", timestamp, errs)
	c.AbortWithStatusJSON(c.Writer.Status(), gin.H{
		"messages":  errs,
		"timestamp": timestamp,
	})
}
