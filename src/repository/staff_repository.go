package repository

import (
	"context"
	"database/sql"
	"inventory-management/model/database"
	"inventory-management/model/dto"
)

type StaffRepository struct {
	db *sql.DB
}

func NewStaffRepository(db *sql.DB) StaffRepositoryInterface {
	return &StaffRepository{db}
}

func (r *StaffRepository) CreateStaff(ctx context.Context, data database.Staff) (dto.RegistrationResponse, error) {
	query := `
	INSERT INTO staffs (phone_number, name, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	var userId string
	err := r.db.QueryRowContext(
		ctx,
		query,
		data.PhoneNumber,
		data.Name,
		data.Password,
		data.CreatedAt,
		data.UpdatedAt,
	).Scan(&userId)

	if err != nil {
		return dto.RegistrationResponse{}, err
	}

	response := dto.RegistrationResponse{
		Message: 201,
		Data: dto.StaffData{
			UserId:       userId,
			PhoneNumber:  data.PhoneNumber,
			Name:         data.Name,
			AccessToken:  "your_access_token_here",
		},
	}

	return response, nil
}

func (r *StaffRepository) GetStaffByPhoneNumber(ctx context.Context, phoneNumber string) (response database.Staff, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT id, name, phone_number FROM staffs WHERE phone_number = $1", phoneNumber).Scan(&response.Id, &response.Name, &response.PhoneNumber)
	return response, err
}
