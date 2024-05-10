package usecase

import (
	"inventory-management/helpers"
	"inventory-management/model/dto"
	"inventory-management/src/repository"
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

func (u *ProductUsecase) CreateProduct(dto.RequestUpsertProduct) (dto.ResponseCreateProduct, error) {
	return dto.ResponseCreateProduct{}, nil
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
