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

type ProductUsecaseInterface interface {
	CreateProduct(dto.RequestUpsertProduct) (dto.ResponseCreateProduct, error)
	GetProduct(dto.RequestGetProduct) (dto.ResponseGetProduct, error)
	UpdateProduct(dto.RequestUpsertProduct) (statusCode int)
	DeleteProduct(id int) (statusCode int)
}
