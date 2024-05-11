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

type SkuUsecaseInterface interface {
	Search(request dto.SearchSkuParams) ([]dto.SearchSkuResponse, error)
}