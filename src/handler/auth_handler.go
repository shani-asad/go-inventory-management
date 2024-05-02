package handler

import (
	"cats-social/model/dto"
	"cats-social/src/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	iAuthUsecase usecase.AuthUsecaseInterface
}

func NewAuthHandler(iAuthUsecase usecase.AuthUsecaseInterface) AuthHandlerInterface {
	return &AuthHandler{iAuthUsecase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request dto.RequestCreateUser
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	err = h.iAuthUsecase.Register(request)
	if err != nil {
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	c.JSON(200, gin.H{
		"message": "user register success",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request dto.RequestAuth
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	token, err := h.iAuthUsecase.Login(request)
	if err != nil {
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
