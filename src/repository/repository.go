package repository

import "cats-social/model/database"

type CatRepositoryInterface interface {
	GetCatById(id int) (response database.Cat, err error)
}
