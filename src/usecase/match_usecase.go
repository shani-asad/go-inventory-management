package usecase

import (
	"cats-social/helpers"
	"cats-social/model/database"
	"cats-social/model/dto"
	"cats-social/src/repository"
	"context"
	"fmt"

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

func (u *MatchUsecase) CreateMatch(request dto.RequestCreateMatch, reqUserId int) error {

	data := database.Match{
		MatchCatId:	request.MatchCatId,
		UserCatId:	request.UserCatId,
		Message:	request.Message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := u.iMatchRepository.CreateMatch(context.TODO(), data, reqUserId)
	return err
}

func (u *MatchUsecase) GetMatch(userId int) ([]dto.ResponseGetMatch, error) {

	response, err := u.iMatchRepository.GetMatch(context.TODO(), userId)
	return response, err
}

func (u *MatchUsecase) GetMatchById(id int) (error) {
	err := u.iMatchRepository.GetMatchById(context.TODO(), id)
	fmt.Println(err)

  return err
}


func (u *MatchUsecase) DeleteMatch(id int) (error) {
	err := u.iMatchRepository.DeleteMatch(context.TODO(), id)
	fmt.Println(err)

	return err
}

func (u *MatchUsecase) GetCatIdByMatchId(id int) (int, int, error) {
	matchCatId, userCatId, err := u.iMatchRepository.GetCatIdByMatchId(context.TODO(), id)
	fmt.Println(err)

  return matchCatId, userCatId, err
}

func (u *MatchUsecase) ApproveMatch(id int, matchCatId int, userCatId int) error {
	err := u.iMatchRepository.ApproveMatch(context.TODO(), id, matchCatId, userCatId)

	return err
}

func (u *MatchUsecase) RejectMatch(id int) error {
	err := u.iMatchRepository.RejectMatch(context.TODO(), id)

	return err
}

