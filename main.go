package main

import (
	"cats-social/app/server"
	"cats-social/db"
	"cats-social/model/properties"
	"cats-social/src/handler"
	"cats-social/src/repository"
	"cats-social/src/usecase"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := server.InitServer()

	postgreConfig := properties.PostgreConfig{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	db := db.InitPostgreDB(postgreConfig)

	// REPOSITORY
	userRepository := repository.NewUserRepository(db)

	// USECASE
	authUsecase := usecase.NewAuthUsecase(userRepository)

	// HANDLER
	catHandler := handler.NewCatHandler()
	authHandler := handler.NewAuthHandler(authUsecase)

	//	ROUTE
	r.GET("/cat", catHandler.GetCatById)
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.Run()
}
