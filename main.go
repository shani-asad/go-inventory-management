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
	staffRepository := repository.NewStaffRepository(db)
	productRepository := repository.NewProductRepository(db)

	// USECASE
	authUsecase := usecase.NewAuthUsecase(staffRepository, helper)
	skuUsecase := usecase.NewSkuUsecase(productRepository)

	// HANDLER
	authHandler := handler.NewAuthHandler(authUsecase)
	skuHandler := handler.NewSkuHandler(skuUsecase)

	//// ROUTE

	// Auth
	r.POST("/v1/staff/register", authHandler.Register)
	r.POST("/v1/staff/login", authHandler.Login)

	// Search AKU
	r.POST("/v1/product/customer", skuHandler.search)

	authorized := r.Group("")
	authorized.Use(middleware.AuthMiddleware)

	r.Run()
}
