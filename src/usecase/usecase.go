package usecase

import "cats-social/model/dto"

type CatUsecaseInterface interface {
	GetCatById(id int) interface{}
}

type AuthUsecaseInterface interface {
	Register(request dto.RequestCreateUser) error
	Login(request dto.RequestAuth) (token string, err error)
}

type MatchUsecaseInterface interface {
	CreateMatch(request dto.RequestCreateMatch) error
}
