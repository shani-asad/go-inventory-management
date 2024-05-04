package usecase

import (
	"cats-social/model/database"
	"cats-social/model/dto"
	"cats-social/src/repository"
	"context"
	"time"
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

func (u *CatUsecase) AddCat(request dto.RequestCreateCat) (int64, error) {
	data := database.Cat{
		Name: request.Name,
		UserId: request.UserId,
    Race: request.Race,
    Sex: request.Sex,
    AgeInMonth: request.AgeInMonth,
    Description: request.Description,
    ImageUrls: request.ImageUrls,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	 // Call the CreateCat method of the repository to add the cat
	 id, err := u.iCatRepository.CreateCat(context.TODO(), data)
	 if err != nil {
			 return 0, err // Return 0 as the ID and the error
	 }

	return id, err
}
