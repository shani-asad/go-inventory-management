package main

import (
	"cats-social/app/server"
	"cats-social/src/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := server.InitServer()

	catHandler := handler.NewCatHandler()

	// cat route
	r.GET("/cat", catHandler.GetCatById)

	// health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.Run()
}
