package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/user-service/internal/domain/dto"
	"github.com/widcraft/user-service/internal/port"
)

type UserHandler struct {
	logger log.FieldLogger
	app    port.UserApp
}

func NewUserHandler(logger log.FieldLogger, app port.UserApp) *UserHandler {
	return &UserHandler{
		logger: logger,
		app:    app,
	}
}

func (h *UserHandler) Register(router *gin.RouterGroup) {
	router.POST("/user", h.register)
	router.GET("/user/:idx", h.getUserByIdx)
}

func (h *UserHandler) register(c *gin.Context) {
	var params struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		ImageUrl string `json:"image_url" binding:"required,url"`
		Token    string `json:"token" binding:"required"`
	}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	h.logger.Info(params.Email)
	err = h.app.Register(&dto.RegisterReqDto{
		Email:    params.Email,
		Name:     params.Name,
		ImageUrl: params.ImageUrl,
		Token:    params.Token,
	})
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *UserHandler) getUserByIdx(c *gin.Context) {
	var params struct {
		Idx uint `uri:"idx"`
	}
	if err := c.ShouldBindUri(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, params)
}
