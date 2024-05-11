package repository

import (
	"inventory-management/model/database"
	"inventory-management/model/dto"
	"context"
)

type StaffRepositoryInterface interface {
	CreateStaff(context.Context, database.Staff) (int, error)
	GetStaffByPhoneNumber(context.Context, string) (database.Staff, error)
}

type ProductRepositoryInterface interface {
	SearchSku(context.Context, dto.SearchSkuParams) ([]dto.SearchSkuResponse, error)
}
