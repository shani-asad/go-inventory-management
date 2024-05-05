package repository

import (
	"cats-social/model/database"
	"cats-social/model/dto"
	"context"
)

type CatRepositoryInterface interface {
	GetCatById(ctx context.Context, id int) (err error)
	CreateCat(ctx context.Context, data database.Cat) (id int64, err error)
	GetCat(ctx context.Context, data dto.RequestGetCat) (response []dto.CatDetail, err error)
	UpdateCat(ctx context.Context, data database.Cat) (err error)
	CheckHasMatch(ctx context.Context, id int) (hasMatched bool, err error)
	DeleteCat(ctx context.Context, id int) (err error)
}

type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (response database.User, err error)
	CreateUser(ctx context.Context, data database.User) (err error)
}

type MatchRepositoryInterface interface {
	CreateMatch(ctx context.Context, data database.Match) (err error)
	GetMatch(ctx context.Context, userId int) (response []dto.ResponseGetMatch, err error)
	GetMatchById(ctx context.Context, id int) (err error)
	DeleteMatch(ctx context.Context, id int) (err error)
	ApproveMatch(ctx context.Context, id int, matchCatId int, userCatId int) (error)
	RejectMatch(ctx context.Context, id int) (error)
	GetCatIdByMatchId(ctx context.Context, id int) (matchCatId int, userCatIs int, err error)
}
