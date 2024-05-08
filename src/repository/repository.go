package repository

import (
	"inventory-management/model/database"
	"inventory-management/model/dto"
	"context"
)

type StaffRepositoryInterface interface {
	CreateStaff(context.Context, database.Staff) (dto.RegistrationResponse, error)
	GetStaffByPhoneNumber(context.Context, string) (database.Staff, error)
}
