package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/widcraft/user-service/internal/domain/dto"
	"github.com/widcraft/user-service/internal/port"
	"github.com/widcraft/user-service/pkg/logger"
)

type Handler struct {
	logger logger.Logger
	app    port.UserApp
}

func New(logger logger.Logger, app port.UserApp) *Handler {
	return &Handler{
		logger: logger,
		app:    app,
	}
}

func (h *Handler) Register(router *gin.RouterGroup) {
	router.POST("/user", h.register)
}

func (h *Handler) register(ctx *gin.Context) {
	var param *register

	err := ctx.ShouldBindJSON(param)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	h.logger.Info(param.Email)
	err = h.app.Register(&dto.RegisterReqDto{
		Email:    param.Email,
		Name:     param.Name,
		ImageUrl: param.ImageUrl,
		Token:    param.Token,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
