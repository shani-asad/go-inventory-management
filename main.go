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
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
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
	// run migrations
	m, err := migrate.New(os.Getenv("MIGRATION_PATH"), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error creating migration instance: ", err)
	}


	// Run the migration up to the latest version
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error applying migrations:", err)
	}

	fmt.Println("Migration successfully applied")

	helper := helpers.NewHelper()

	// MIDDLEWARE
	middleware := middleware.NewMiddleware(helper)
	// REPOSITORY
	userRepository := repository.NewUserRepository(db)
	matchRepository := repository.NewMatchRepository(db)
	catRepository := repository.NewCatRepository(db)

	// USECASE
	authUsecase := usecase.NewAuthUsecase(userRepository, helper)
	matchUseCase := usecase.NewMatchUsecase(matchRepository, helper)
	catUsecase := usecase.NewCatUsecase(catRepository)


	// HANDLER
	catHandler := handler.NewCatHandler(catUsecase)
	authHandler := handler.NewAuthHandler(authUsecase)
	matchHandler := handler.NewMatchHandler(matchUseCase)

	// ROUTE
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.POST("/v1/register", authHandler.Register)
	r.POST("/v1/login", authHandler.Login)

	authorized := r.Group("")
	authorized.Use(middleware.AuthMiddleware)

	// CAT
	authorized.POST("v1/cat", catHandler.AddCat)
	authorized.GET("v1/cat", catHandler.GetCat)
	authorized.PUT("v1/cat/:id", catHandler.UpdateCat)
	authorized.DELETE("v1/cat/:id", catHandler.DeleteCat)

	authorized.POST("/v1/cat/match", matchHandler.CreateMatch)
	authorized.GET("/v1/cat/match", matchHandler.GetMatch)
	authorized.DELETE("v1/cat/match/:id", matchHandler.DeleteMatch)

	r.Run()
}
