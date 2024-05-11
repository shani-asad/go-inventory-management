package repository

import (
	"context"
	"database/sql"
	"inventory-management/model/database"
	"inventory-management/model/dto"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepositoryInterface {
	return &CustomerRepository{db}
}

func (r *CustomerRepository) RegisterCustomer(ctx context.Context, data database.Customer) (string, error) {
	var id string

	query := `INSERT INTO customers (phone_number, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.db.QueryRowContext(
			ctx,
			query,
			data.PhoneNumber,
			data.Name,
			data.CreatedAt,
			data.UpdatedAt,
	).Scan(&id)

	if err != nil {
			return "", err
	}

	return id, nil
}

func (r *CustomerRepository) SearchCustomers(ctx context.Context, data dto.SearchCustomersRequest) ([]dto.CustomerDTO, error) {
	// Perform database operation to find customers
	// Example implementation:
	// Query the database based on phoneNumber and name
	// Return the found customers
	return []dto.CustomerDTO{}, nil
}