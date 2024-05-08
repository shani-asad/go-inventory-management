package main

import (
	"fmt"
	"inventory-management/app/server"
	"inventory-management/db"
	"inventory-management/helpers"
	"inventory-management/model/properties"
	"inventory-management/src/handler"
	"inventory-management/src/middleware"
	"inventory-management/src/repository"
	"inventory-management/src/usecase"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	r := server.InitServer()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbParams := os.Getenv("DB_PARAMS")

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?%s",
		dbUsername, dbPassword, dbHost, dbPort, dbName, dbParams,
	)

	postgreConfig := properties.PostgreConfig{
		DatabaseURL: connectionString,
	}

	db := db.InitPostgreDB(postgreConfig)


	helper := helpers.NewHelper()

	// MIDDLEWARE
	middleware := middleware.NewMiddleware(helper)

	// REPOSITORY
	userRepository := repository.NewStaffRepository(db)

	// USECASE
	authUsecase := usecase.NewAuthUsecase(userRepository, helper)

	// HANDLER
	authHandler := handler.NewAuthHandler(authUsecase)

	// ROUTE
	r.POST("/v1/staff/register", authHandler.Register)
	r.POST("/v1/staff/login", authHandler.Login)

	authorized := r.Group("")
	authorized.Use(middleware.AuthMiddleware)

	r.Run()
}
