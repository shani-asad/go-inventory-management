package usecase

import "cats-social/model/dto"

type AuthUsecase struct {
}

func NewAuthUsecase(u *AuthUsecase) AuthUsecaseInterface {
	return &AuthUsecase{}
}

func (u *AuthUsecase) Register(request dto.RequestAuth) {

}

func (u *AuthUsecase) Login(request dto.RequestAuth) {

}
