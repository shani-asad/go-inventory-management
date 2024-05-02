package usecase

import "cats-social/model/dto"

type CatUsecaseInterface interface {
	GetCatById(id int) interface{}
}

type AuthUsecaseInterface interface {
	Register(request dto.RequestAuth)
	Login(request dto.RequestAuth)
}