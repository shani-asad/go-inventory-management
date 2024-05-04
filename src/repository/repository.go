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
	Getmatch(ctx context.Context) (response database.Match, err error)
}
