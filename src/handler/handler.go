package handler

import "github.com/gin-gonic/gin"

type CatHandlerInterface interface {
	GetCatById(c *gin.Context)
	AddCat(c *gin.Context)
	GetCat(c *gin.Context)
	UpdateCat(c *gin.Context)
	DeleteCat(c *gin.Context)
}

type AuthHandlerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type MatchHandlerInterface interface {
	CreateMatch(c *gin.Context)	
	GetMatch(c *gin.Context)	
	DeleteMatch(c *gin.Context)
}