package usecase

import (
	"context"
	"inventory-management/model/dto"
	"inventory-management/src/repository"
)

type SkuUsecase struct {
	iProductRepository repository.ProductRepositoryInterface
}

func NewSkuUsecase(
		iProductRepository repository.ProductRepositoryInterface,
	) SkuUsecaseInterface {
	return &SkuUsecase{iProductRepository}
}

func (u *SkuUsecase) Search(request dto.SearchSkuParams) ([]dto.SearchSkuResponse, error) {
	params := dto.SearchSkuParams{
		Limit:    validateLimit(request.Limit), // TODO - validate user input is integer. Do this in handler when doing ShouldBindJSON()
		Offset:   validateOffset(request.Offset),
		Name:     validateName(request.Name),
		Category: validateCategory(request.Category),
		Sku:      validateSKU(request.Sku),
		Price:    validatePrice(request.Price),
		InStock:  validateInStock(request.InStock),
		IsInstockValid: request.IsInstockValid,
	}
	
	response, err := u.iProductRepository.SearchSku(context.TODO(), params)
	
	return response, err
}

func validateLimit(limit int) int {
    if limit >= 0 {
        return limit
    }

}

func validateOffset(offset int) int {
    if offset >= 0 {
        return offset
    }

}


func validateName(name string) string {
	return name
}

func validateCategory(category string) string {
	allowedCategories := map[string]bool{"Clothing": true, "Accessories": true, "Footwear": true, "Beverages": true}
	if allowedCategories[category] {
		return category
	} else {
		return ""
	}

}

func validateSKU(sku string) string {
	return sku
}

func validatePrice(price string) string {
	if price == "asc" || price == "desc" {
		return price
	}

}

func validateInStock(inStock bool) bool {
    return inStock
}