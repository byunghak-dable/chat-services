package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/api-service/internal/domain/dto"
	"github.com/widcraft/api-service/internal/port"
)

type Handler struct {
	logger log.FieldLogger
	app    port.UserApp
}

func New(logger log.FieldLogger, app port.UserApp) *Handler {
	return &Handler{
		logger: logger,
		app:    app,
	}
}

func (h *Handler) Register(router *gin.RouterGroup) {
	router.POST("/user", h.register)
}

func (h *Handler) register(c *gin.Context) {
	var register *register

	err := c.ShouldBindJSON(register)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	h.logger.Info(register.Email)
	err = h.app.Register(&dto.RegisterReqDto{
		Email:    register.Email,
		Name:     register.Name,
		ImageUrl: register.ImageUrl,
		Token:    register.Token,
	})
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
