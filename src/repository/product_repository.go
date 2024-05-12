package repository

import (
	"context"
	"database/sql"
	"fmt"
	"inventory-management/model/database"
	"inventory-management/model/dto"
	"strings"
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
			   RETURNING id, created_at`

	// Execute the SQL query
	var id int
	err := r.db.QueryRowContext(ctx, query, data.Name, data.SKU, data.Category, data.ImageURL, data.Notes, data.Price, data.Stock, data.Location, data.IsAvailable, createdAt, updatedAt).Scan(&id, &createdAt)
	if err != nil {
		return database.Product{}, fmt.Errorf("failed to create product: %v", err)
	}

	// Set the generated ID in the product object
	data.ID = id
	data.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)

	return data, nil
}

func (r *ProductRepository) GetProduct(ctx context.Context, param dto.RequestGetProduct) (products []database.Product, err error) {
	query := "SELECT * FROM products WHERE 1=1"

	// Add conditions based on the request parameters
	var args []interface{}

	if param.ID != "" {
		query += " AND id = $1"
		args = append(args, param.ID)
	}

	if param.Name != "" {
		query += " AND LOWER(name) LIKE LOWER($2)"
		args = append(args, "%"+param.Name+"%")
	}

	if param.IsAvailable {
		query += " AND is_available = true"
	}

	if param.Category != "" {
		query += " AND category = $3"
		args = append(args, param.Category)
	}

	if param.SKU != "" {
		query += " AND sku = $4"
		args = append(args, param.SKU)
	}

	if param.Instock {
		query += " AND stock > 0"
	}

	if param.CreatedAt != "" {
		if strings.ToLower(param.CreatedAt) == "asc" {
			query += " ORDER BY created_at ASC"
		} else if strings.ToLower(param.CreatedAt) == "desc" {
			query += " ORDER BY created_at DESC"
		}
	}

	query += " LIMIT $5 OFFSET $6"

	// Add limit and offset
	args = append(args, param.Limit)
	args = append(args, param.Offset)

	// Execute the SQL query
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %v", err)
	}
	defer rows.Close()

	// Iterate over the result set
	for rows.Next() {
		var product database.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.SKU,
			&product.Category,
			&product.ImageURL,
			&product.Notes,
			&product.Price,
			&product.Stock,
			&product.Location,
			&product.IsAvailable,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration of product rows: %v", err)
	}

	return products, nil
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

func (r *ProductRepository) SearchSku(ctx context.Context, params dto.SearchSkuParams) (response []dto.SearchSkuResponse, err error) {

	query := constructQuery(params)
	rows, err := r.db.QueryContext(ctx, query)

	for rows.Next() {
		var sku dto.SearchSkuResponse
		err := rows.Scan(&sku)
		fmt.Println(err)
		if err != nil {
			return nil, err
		}
		response = append(response, sku)
	}
	return response, err
}

func constructQuery(params dto.SearchSkuParams) string {
	query := "SELECT * FROM products WHERE 1=1"
	if params.Name != "" {
		query += fmt.Sprintf(" AND LOWER(name) LIKE LOWER('%%%s%%')", params.Name)
	}
	if params.Category != "" {
		query += fmt.Sprintf(" AND category = '%s'", params.Category)
	}
	if params.Sku != "" {
		query += fmt.Sprintf(" AND sku = '%s'", params.Sku)
	}
	if params.IsInstockValid {
		if params.InStock {
			query += " AND stock > 0"
		} else {
			query += " AND stock < 1"
		}
	}
	if params.Price == "asc" {
		query += " ORDER BY price ASC"
	} else if params.Price == "desc" {
		query += " ORDER BY price DESC"
	}
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", params.Limit, params.Offset)
	return query
}
