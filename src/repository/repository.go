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

type CustomerRepositoryInterface interface {
  RegisterCustomer(context.Context, database.Customer) (string, error)
	SearchCustomers(context.Context, dto.SearchCustomersRequest) ([]dto.CustomerDTO, error)
	GetCustomerByPhoneNumber(context.Context, string) (database.Customer, error)
}
