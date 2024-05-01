package handler

import "github.com/gin-gonic/gin"

type CatHandler struct {
}

func NewCatHandler() CatHandlerInterface {
	return &CatHandler{}
}

func (h *CatHandler) GetCatById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "cat by id",
	})
}
