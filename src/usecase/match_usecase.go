package usecase

import (
	"cats-social/helpers"
	"cats-social/model/database"
	"cats-social/model/dto"
	"cats-social/src/repository"
	"context"
	// "errors"
	// "fmt"
	"time"
)

type MatchUsecase struct {
	iMatchRepository repository.MatchRepositoryInterface
}

func NewMatchUsecase(
	iMatchRepository repository.MatchRepositoryInterface,
	helper helpers.HelperInterface) MatchUsecaseInterface {
	return &MatchUsecase{iMatchRepository}
}

func (u *MatchUsecase) CreateMatch(request dto.RequestCreateMatch) error {

	data := database.Match{
		MatchCatId:	request.MatchCatId,
		UserCatId:	request.UserCatId,
		Message:	request.Message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := u.iMatchRepository.CreateMatch(context.TODO(), data)
	return err
}

func (u *MatchUsecase) GetMatch(userId int) ([]dto.ResponseGetMatch, error) {

	response, err := u.iMatchRepository.GetMatch(context.TODO(), userId)
	return response, err
}