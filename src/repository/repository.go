package repository

import (
	"context"
	"inventory-management/model/database"
	"inventory-management/model/dto"
)

type StaffRepositoryInterface interface {
	CreateStaff(context.Context, database.Staff) (dto.RegistrationResponse, error)
	GetStaffByPhoneNumber(context.Context, string) (database.Staff, error)
}

type ProductRepositoryInterface interface {
	CreateProduct(context.Context, database.Product) (database.Product, error)
	GetProduct(context.Context, dto.RequestGetProduct) ([]database.Product, error)
	UpdateProduct(context.Context, database.Product) (database.Product, error)
	DeleteProduct(context.Context, int) error
}
