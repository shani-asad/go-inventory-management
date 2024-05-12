package usecase

import (
	"context"
	"errors"
	"fmt"
	"inventory-management/helpers"
	"inventory-management/model/database"
	"inventory-management/model/dto"
	"inventory-management/src/repository"
	"strconv"
)

const (
	ErrCustomerNotFound = "customer_not_found"
	ErrProductNotFound = "product_not_found"
	ErrValidation        = "validation_error"
	ErrPaidNotEnough     = "paid_not_enough"
	ErrChangeIncorrect   = "change_incorrect"
	ErrProductStockInsufficient = "product_stock_insufficient"
	ErrProductNotAvailable = "product_not_available"
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
		IsAvailable: *data.IsAvailable,
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
	param.Category = validateCategory(param.Category)
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

	if len(response.Data) == 0 {
		response.Data = []dto.Product{};
	}
	return response, err
}

func (u *ProductUsecase) UpdateProduct(data dto.RequestUpsertProduct) (statusCode int) {
	product := database.Product{
		ID:          data.ID,
		Name:        data.Name,
		SKU:         data.SKU,
		Category:    data.Category,
		ImageURL:    data.ImageURL,
		Notes:       data.Notes,
		Price:       data.Price,
		Stock:       data.Stock,
		Location:    data.Location, 
		IsAvailable: *data.IsAvailable,
	}
	_, err := u.iProductRepository.UpdateProduct(context.TODO(), product)
	if err != nil {
		return 500
	}

	return 200
}

func (u *ProductUsecase) DeleteProduct(id int) (statusCode int) {
	return u.iProductRepository.DeleteProduct(context.TODO(), id)
}

func (u *ProductUsecase) CheckoutProduct(req dto.CheckoutProductRequest) error {
	// Validate input parameters
	if req.CustomerID == "" {
		return errors.New(ErrCustomerNotFound)
	}
	if len(req.ProductDetails) == 0 {
		return errors.New(ErrProductNotFound) 
	}
	if req.Paid <= 0 {
		return errors.New(ErrValidation)
	}

	// Calculate total price of all products
	totalPrice := 0.0
	for _, pd := range req.ProductDetails {
			if pd.Quantity <= 0 {
					return errors.New(ErrValidation)
			}
			// Retrieve product details from database to check stock
			productId, err := strconv.Atoi(pd.ProductID)
			if err != nil {
				fmt.Println("Error:", err)
			}
			product, err := u.iProductRepository.GetProductById(context.TODO(), productId)
			if err != nil {
					return err
			}
			if product.ID == 0 {
					return errors.New(ErrProductNotFound)
			}
			if !product.IsAvailable {
					return errors.New(ErrProductNotAvailable)
			}
			if product.Stock < pd.Quantity {
					return errors.New(ErrProductStockInsufficient)
			}
			totalPrice += float64(pd.Quantity) * product.Price
	}

	// Check if paid amount is enough
	if float64(req.Paid) < totalPrice {
			return errors.New(ErrPaidNotEnough)
	}

	// Check if change is correct
	if float64(req.Change) != (float64(req.Paid) - totalPrice) {
			return errors.New(ErrChangeIncorrect)
	}

	// Start a transaction
	tx, err := u.iProductRepository.BeginTransaction(context.TODO())
	if err != nil {
			return err
	}
	defer func() {
			// Rollback the transaction if there's an error
			if err != nil {
					tx.Rollback()
					return
			}
			// Commit the transaction if there's no error
			tx.Commit()
	}()

	// If all validations pass, update product stock and checkout
	for _, pd := range req.ProductDetails {
			err := u.iProductRepository.UpdateProductStockInTransaction(context.TODO(), tx, pd.ProductID, pd.Quantity)
			if err != nil {
					return err
			}
			// Create transaction record
			customerID, _ := strconv.Atoi(req.CustomerID)
			transactionID, err := u.iProductRepository.CreateTransaction(context.TODO(), tx, customerID, req.Paid, req.Change)
			if err != nil {
					return err
			}

			// Add transaction products
			productId, _ := strconv.Atoi(pd.ProductID)
			for _, pd := range req.ProductDetails {
					err := u.iProductRepository.AddTransactionProduct(context.TODO(), tx, transactionID, productId, pd.Quantity)
					if err != nil {
							return err
					}
			}
	}

	return nil
}
