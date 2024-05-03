package usecase

import (
	"cats-social/model/database"
	"cats-social/model/dto"
	"cats-social/src/repository"
	"context"
)

type CatUsecase struct {
	iCatRepository repository.CatRepositoryInterface
}

func NewCatUsecase(
	iCatRepository repository.CatRepositoryInterface) CatUsecaseInterface {
  return &CatUsecase{iCatRepository}
}

func (u *CatUsecase) GetCatById(id int) interface{} {
	return nil
}

func (u *CatUsecase) AddCat(request dto.RequestCreateCat) error {
	data := database.Cat{
		Name: request.Name,
    Race: request.Race,
    Sex: request.Sex,
    AgeInMonth: request.AgeInMonth,
    Description: request.Description,
    ImageUrls: request.ImageUrls,
	}

	err := u.iCatRepository.CreateCat(context.TODO(), data)
	return err
}
