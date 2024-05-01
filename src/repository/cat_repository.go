package repository

import "cats-social/model/database"

type CatRepository struct {
}

func NewCatRepository() CatRepositoryInterface {
	return &CatRepository{}
}

func (r *CatRepository) GetCatById(id int) (response database.Cat, err error) {
	return response, err
}
