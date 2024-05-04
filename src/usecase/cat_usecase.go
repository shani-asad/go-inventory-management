package usecase

import (
	"cats-social/model/database"
	"cats-social/model/dto"
	"cats-social/src/repository"
	"context"
	"fmt"
	"time"
)

type CatUsecase struct {
	iCatRepository repository.CatRepositoryInterface
}

func NewCatUsecase(
	iCatRepository repository.CatRepositoryInterface) CatUsecaseInterface {
  return &CatUsecase{iCatRepository}
}

func (u *CatUsecase) GetCatById(id int) (error) {
	err := u.iCatRepository.GetCatById(context.TODO(), id)
	fmt.Println(err)

  return err
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

func (u *CatUsecase) GetCat(request dto.RequestGetCat) ([]dto.CatDetail, error) {
	cats, err := u.iCatRepository.GetCat(context.TODO(), request)
	fmt.Println(err)
  if err != nil {
    return nil, err
  }

  return cats, err
}

func (u *CatUsecase) UpdateCat(request dto.RequestCreateCat, id int64) (error) {
  cat := database.Cat{
    Id:        int(id),
    Name:      request.Name,
    UserId:    request.UserId,
    Race:      request.Race,
    Sex:       request.Sex,
    AgeInMonth: request.AgeInMonth,
    Description: request.Description,
    ImageUrls: request.ImageUrls,
    UpdatedAt: time.Now(),
  }

  err := u.iCatRepository.UpdateCat(context.TODO(), cat)
  return err
}

func (u *CatUsecase) CheckHasMatch(id int) (bool, error) {
	hasMatched, err := u.iCatRepository.CheckHasMatch(context.TODO(), id)
	fmt.Println(err)
	if err != nil {
    return false, err
  }

	return hasMatched, err
}

func (u *CatUsecase) DeleteCat(id int) (error) {
	err := u.iCatRepository.DeleteCat(context.TODO(), id)
	fmt.Println(err)

	return err
}
