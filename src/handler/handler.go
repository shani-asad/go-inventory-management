package handler

import "github.com/gin-gonic/gin"

type CatHandlerInterface interface {
	GetCatById(c *gin.Context)
}

type AuthHandlerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type MatchHandlerInterface interface {
	CreateMatch(c *gin.Context)	
}