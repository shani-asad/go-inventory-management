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
	"github.com/joho/godotenv"
	"fmt"
)

func main() {
	r := server.InitServer()

	err := 	godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file:", err)
        return
    }

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

	r.POST("/v1/register", authHandler.Register)
	r.POST("/v1/login", authHandler.Login)

	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware)

	authorized.GET("/v1/cat", catHandler.GetCatById)

	r.Run()
}
