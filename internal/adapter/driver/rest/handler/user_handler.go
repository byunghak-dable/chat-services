package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	logger *log.Logger
}

func NewUserHandler(logger *log.Logger) *UserHandler {
	return &UserHandler{
		logger: logger,
	}
}

func (h UserHandler) Register(router *gin.RouterGroup) {
	router.POST("/user", h.register)
	router.GET("/user/:idx", h.getUserByIdx)
}

func (h UserHandler) register(c *gin.Context) {
	var params struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
		Token string `json:"token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, params)
}

func (h UserHandler) getUserByIdx(c *gin.Context) {
	var params struct {
		Idx uint `uri:"idx"`
	}
	if err := c.ShouldBindUri(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, params)
}
