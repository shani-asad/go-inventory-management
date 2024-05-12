package repository

import (
	"context"
	"database/sql"
	"inventory-management/model/database"
)

type StaffRepository struct {
	db *sql.DB
}

func NewStaffRepository(db *sql.DB) StaffRepositoryInterface {
	return &StaffRepository{db}
}

func (r *StaffRepository) CreateStaff(ctx context.Context, data database.Staff) (int, error) {
	query := `
	INSERT INTO staffs (phone_number, name, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	var staffId int
	err := r.db.QueryRowContext(
		ctx,
		query,
		data.PhoneNumber,
		data.Name,
		data.Password,
		data.CreatedAt,
		data.UpdatedAt,
	).Scan(&staffId)

	return staffId, err
}

func (r *StaffRepository) GetStaffByPhoneNumber(ctx context.Context, phoneNumber string) (response database.Staff, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT id, name, phone_number, password FROM staffs WHERE phone_number = $1", phoneNumber).Scan(&response.Id, &response.Name, &response.PhoneNumber, &response.Password)
	return response, err
}
