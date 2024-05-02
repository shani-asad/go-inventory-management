package main

import (
	"cats-social/app/server"
	"cats-social/db"
	"cats-social/helpers"
	"cats-social/model/properties"
	"cats-social/src/handler"
	"cats-social/src/middleware"
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
	helper := helpers.NewHelper()

	// MIDDLEWARE
	middleware := middleware.NewMiddleware(helper)
	// REPOSITORY
	userRepository := repository.NewUserRepository(db)

	// USECASE
	authUsecase := usecase.NewAuthUsecase(userRepository, helper)

	// HANDLER
	catHandler := handler.NewCatHandler()
	authHandler := handler.NewAuthHandler(authUsecase)

	//	ROUTE
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware)

	authorized.GET("/cat", catHandler.GetCatById)

	r.Run()
}
