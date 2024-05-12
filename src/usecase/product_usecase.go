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
		IsAvailable: data.IsAvailable,
	}

	db, err := u.iProductRepository.CreateProduct(context.TODO(), product)
	if err != nil {
		return response, err
	}

	response.Message = "success"
	response.Data.ID = strconv.Itoa(db.ID)
	response.Data.CreatedAt = db.CreatedAt
	return response, nil
}

func (u *ProductUsecase) GetProduct(param dto.RequestGetProduct) (response dto.ResponseGetProduct, err error) {
	products, err := u.iProductRepository.GetProduct(context.TODO(), param)
	if err != nil {
		return response, err
	}

	response.Message = "success"
	for _, v := range products {
		product := dto.Product{
			ID:          strconv.Itoa(v.ID),
			Name:        v.Name,
			SKU:         v.SKU,
			Category:    v.Category,
			ImageURL:    v.ImageURL,
			Stock:       v.Stock,
			Notes:       v.Notes,
			Price:       v.Price,
			Location:    v.Location,
			IsAvailable: v.IsAvailable,
			CreatedAt:   v.CreatedAt,
		}
		response.Data = append(response.Data, product)
	}

	return response, err
}

func (u *ProductUsecase) UpdateProduct(dto.RequestUpsertProduct) (statusCode int) {
	return 0
}

func (u *ProductUsecase) DeleteProduct(id int) (statusCode int) {
	return u.iProductRepository.DeleteProduct(context.TODO(), id)
}
