package repository

import (
	"context"
	"database/sql"
	"fmt"
	"inventory-management/model/database"
	"inventory-management/model/dto"
	"time"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepositoryInterface {
	return &ProductRepository{db}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, data database.Product) (database.Product, error) {
	createdAt := time.Now().Format(time.RFC3339)
	updatedAt := time.Now().Format(time.RFC3339)

	query := `INSERT INTO products (name, sku, category, image_url, notes, price, stock, location, is_available, created_at, updated_at)
			   VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
			   RETURNING id`

	// Execute the SQL query
	var id int
	err := r.db.QueryRowContext(ctx, query, data.Name, data.SKU, data.Category, data.ImageURL, data.Notes, data.Price, data.Stock, data.Location, data.IsAvailable, createdAt, updatedAt).Scan(&id)
	if err != nil {
		return database.Product{}, fmt.Errorf("failed to create product: %v", err)
	}

	// Set the generated ID in the product object
	data.ID = id

	return data, nil
}

func (r *ProductRepository) GetProduct(context.Context, dto.RequestGetProduct) ([]database.Product, error) {
	return []database.Product{}, nil
}

func (r *ProductRepository) UpdateProduct(context.Context, database.Product) (database.Product, error) {
	return database.Product{}, nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id int) int {
	// Prepare the SQL query
	query := "DELETE FROM products WHERE id = $1"

	// Execute the SQL query
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		fmt.Printf("failed to delete product: %v", err)
		return 500
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("failed to get rows affected: %v", err)
		return 404
	}
	if rowsAffected == 0 {
		fmt.Printf("product with ID %d not found", id)
		return 404
	}

	return 200
}
