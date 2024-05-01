package mapper

import (
	"cats-social/model/database"
	"cats-social/model/dto"
)

type CatMapper struct{}

func NewCatMapper() CatMapperInterface {
	return &CatMapper{}
}

func (m *CatMapper) ToResponseGetCatByID(data database.Cat) (response dto.ToResponseGetCat) {
	return response
}
