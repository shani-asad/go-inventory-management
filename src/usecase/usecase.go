package usecase

import (
	"inventory-management/model/database"
	"inventory-management/model/dto"
)

type AuthUsecaseInterface interface {
	Register(request dto.RequestCreateStaff) (token string, err error)
	Login(request dto.RequestAuth) (token string, user database.Staff, err error)
	GetStaffByPhoneNumber(email string) (exists bool, err error)
}

type CustomerUsecaseInterface interface {
  RegisterCustomer(request dto.RegisterCustomerRequest) (customer string, err error)
  SearchCustomers(request dto.SearchCustomersRequest) (customers []dto.CustomerDTO, err error)
	GetCustomerByPhoneNumber(phoneNumber string) (exists bool, err error)
}