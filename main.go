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
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	r := server.InitServer()

	if os.Getenv("DEVELOPER_NAME") != "naja" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file:", err)
			return
		}
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

	if os.Getenv("DEVELOPER_NAME") == "naja" {
		// run migrations
		m, err := migrate.New(os.Getenv("MIGRATION_PATH"), os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal("Error creating migration instance: ", err)
		}

		// Run the migration up to the latest version
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Error applying migrations:", err)
		}
	}

	fmt.Println("Migration successfully applied")

	helper := helpers.NewHelper()

	// MIDDLEWARE
	middleware := middleware.NewMiddleware(helper)

	// REPOSITORY
	staffRepository := repository.NewStaffRepository(db)
	productRepository := repository.NewProductRepository(db)
	customerRepository := repository.NewCustomerRepository(db)

	// USECASE
	authUsecase := usecase.NewAuthUsecase(staffRepository, helper)
	productUsecase := usecase.NewProductUsecase(productRepository, helper)
	skuUsecase := usecase.NewSkuUsecase(productRepository)
	customerUsecase := usecase.NewCustomerUsecase(customerRepository)

	// HANDLER
	authHandler := handler.NewAuthHandler(authUsecase)
	productHandler := handler.NewProductHandler(productUsecase)
	skuHandler := handler.NewSkuHandler(skuUsecase)
	customerHandler := handler.NewCustomerHandler(customerUsecase)

	//// ROUTE

	// Auth
	r.POST("/v1/staff/register", authHandler.Register)
	r.POST("/v1/staff/login", authHandler.Login)
	r.POST("/v1/customer/register", customerHandler.RegisterCustomer)
	r.GET("v1/customer", customerHandler.SearchCustomers)

	// Search AKU
	r.POST("/v1/product/customer", skuHandler.Search)

	authorized := r.Group("")
	authorized.Use(middleware.AuthMiddleware)

	authorized.POST("/v1/product", productHandler.CreateProduct)
	authorized.GET("/v1/product", productHandler.GetProduct)
	authorized.PUT("/v1/product/:id", productHandler.UpdateProduct)
	authorized.DELETE("/v1/product/:id", productHandler.DeleteProduct)

	r.Run()
}
