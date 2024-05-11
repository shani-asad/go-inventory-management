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

func (r *ProductRepository) DeleteProduct(context.Context, int) error {
	return nil
}
