package repository

import (
	"context"
	"database/sql"
	"inventory-management/model/database"
	"inventory-management/model/dto"
	"strconv"
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
	query := "SELECT id, phone_number, name FROM customers WHERE 1=1"

	var args []interface{}

	if data.PhoneNumber != "" {
		query += " AND phone_number LIKE $" + strconv.Itoa(len(args)+1)
		args = append(args, "%"+data.PhoneNumber+"%")
	}

	if data.Name != "" {
		query += " AND name LIKE $" + strconv.Itoa(len(args)+1)
		args = append(args, "%"+data.Name+"%")
	}

	query += " ORDER BY created_at DESC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []dto.CustomerDTO
	for rows.Next() {
		var customer dto.CustomerDTO
		if err := rows.Scan(&customer.UserId, &customer.PhoneNumber, &customer.Name); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *CustomerRepository) GetCustomerByPhoneNumber(ctx context.Context, phoneNumber string) (response database.Customer, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT id, name, phone_number FROM customers WHERE phone_number = $1", phoneNumber).Scan(&response.Id, &response.Name, &response.PhoneNumber)
	return response, err
}
