package handler

import "github.com/gin-gonic/gin"

type CatHandlerInterface interface {
	GetCatById(c *gin.Context)
}
