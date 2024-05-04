package usecase

import "cats-social/model/dto"

type CatUsecaseInterface interface {
	GetCatById(id int) (err error)
	AddCat(request dto.RequestCreateCat) (id int64, err error)
	GetCat(request dto.RequestGetCat) (cats []dto.CatDetail, err error)
	UpdateCat(request dto.RequestCreateCat, id int64) (err error)
	CheckHasMatch(id int) (hasMatched bool, err error)
	DeleteCat(id int) (err error)
}

type AuthUsecaseInterface interface {
	Register(request dto.RequestCreateUser) error
	Login(request dto.RequestAuth) (token string, err error)
}

type MatchUsecaseInterface interface {
	CreateMatch(request dto.RequestCreateMatch) error
}
