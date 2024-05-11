package handler

import "github.com/gin-gonic/gin"

type AuthHandlerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type SkuHandlerInterface interface {
	Search(c *gin.Context)
}

type CustomerHandlerInterface interface {
	RegisterCustomer(c *gin.Context)
	SearchCustomers(c *gin.Context)
}