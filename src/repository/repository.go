package repository

import (
	"context"
	"inventory-management/model/database"
	"inventory-management/model/dto"
)

type StaffRepositoryInterface interface {
	CreateStaff(context.Context, database.Staff) (int, error)
	GetStaffByPhoneNumber(context.Context, string) (database.Staff, error)
}

type ProductRepositoryInterface interface {
	CreateProduct(ctx context.Context, data database.Product) (response database.Product, err error)
	GetProduct(ctx context.Context, data dto.RequestGetProduct) (response []database.Product, err error)
	UpdateProduct(ctx context.Context, data database.Product) (response database.Product, err error)
	DeleteProduct(ctx context.Context, id int) (statusCode int)
	SearchSku(context.Context, dto.SearchSkuParams) ([]dto.SearchSkuResponse, error)
}

type CustomerRepositoryInterface interface {
	RegisterCustomer(context.Context, database.Customer) (string, error)
	SearchCustomers(context.Context, dto.SearchCustomersRequest) ([]dto.CustomerDTO, error)
	GetCustomerByPhoneNumber(context.Context, string) (database.Customer, error)
}
