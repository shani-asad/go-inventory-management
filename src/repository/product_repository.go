package repository

import (
	"context"
	"database/sql"
	"inventory-management/model/database"
	"inventory-management/model/dto"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepositoryInterface {
	return &ProductRepository{db}
}

func (r *ProductRepository) CreateProduct(context.Context, database.Product) (database.Product, error) {
	return database.Product{}, nil
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
