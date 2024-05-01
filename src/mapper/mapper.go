package mapper

import (
	"cats-social/model/database"
	"cats-social/model/dto"
)

type CatMapperInterface interface {
	ToResponseGetCatByID(data database.Cat) (response dto.ToResponseGetCat)
}