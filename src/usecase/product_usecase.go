package usecase

import (
	"context"
	"inventory-management/helpers"
	"inventory-management/model/database"
	"inventory-management/model/dto"
	"inventory-management/src/repository"
	"strconv"
)

type ProductUsecase struct {
	iProductRepository repository.ProductRepositoryInterface
	helper             helpers.HelperInterface
}

func NewProductUsecase(
	iProductRepository repository.ProductRepositoryInterface,
	helper helpers.HelperInterface) ProductUsecaseInterface {
	return &ProductUsecase{iProductRepository, helper}
}

func (u *ProductUsecase) CreateProduct(data dto.RequestUpsertProduct) (response dto.ResponseCreateProduct, err error) {
	product := database.Product{
		Name:        data.Name,
		SKU:         data.SKU,
		Category:    data.Category,
		ImageURL:    data.ImageURL,
		Notes:       data.Notes,
		Price:       data.Price,
		Stock:       data.Stock,
		Location:    data.Location,
		IsAvailable: false,
	}

	db, err := u.iProductRepository.CreateProduct(context.TODO(), product)
	if err != nil {
		return response, err
	}

	response.Data.ID = strconv.Itoa(db.ID)
	return response, nil
}

func (u *ProductUsecase) GetProduct(dto.RequestGetProduct) (dto.ResponseGetProduct, error) {
	return dto.ResponseGetProduct{}, nil
}

func (u *ProductUsecase) UpdateProduct(dto.RequestUpsertProduct) (statusCode int) {
	return 0
}

func (u *ProductUsecase) DeleteProduct(id int) (statusCode int) {
	return 0
}
