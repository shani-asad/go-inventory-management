package handler

import (
	"fmt"
	"inventory-management/model/dto"
	"inventory-management/src/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ProductHandler struct {
	iProductUsecase usecase.ProductUsecaseInterface
}

func NewProductHandler(iProductUsecase usecase.ProductUsecaseInterface) ProductHandlerInterface {
	return &ProductHandler{iProductUsecase}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var request dto.RequestUpsertProduct
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Register bad request (ShouldBindJSON) >> ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	err = validateProduct(request)
	if err != nil {
		log.Println("Create Product bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	response, err := h.iProductUsecase.CreateProduct(request)
	if err != nil {
		c.JSON(500, gin.H{"status": "internal server error", "message": err.Error()})
	}

	c.JSON(201, response)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	param := dto.RequestGetProduct{}

	err := c.ShouldBind(&param)
	if err != nil {
		log.Println("Product bad request (ShouldBindJSON) >> ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	response, err := h.iProductUsecase.GetProduct(param)
	if err != nil {
		fmt.Println("error get product", err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err.Error()})
		return
	}

	c.JSON(200, response)

}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {

}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productID, _ := strconv.Atoi(id)
	statusCode := h.iProductUsecase.DeleteProduct(productID)

	c.JSON(statusCode, dto.ResponseSuccess{
		Status:  "success",
		Message: fmt.Sprintf("product with id %s successfull deleted", id),
	})
}

func (h *ProductHandler) CheckoutProduct(c *gin.Context) {
	var request dto.CheckoutProductRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Checkout product bad request (ShouldBindJSON) >> ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	// Validate the request
	if err := validateCheckoutProductRequest(request); err != nil {
		log.Println("Checkout product validation failed >> ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "request doesn’t pass validation"})
		return
	}

	// Call the usecase layer
	if err := h.iProductUsecase.CheckoutProduct(request); err != nil {
		log.Println("Failed to checkout product >> ", err)
    switch err.Error() {
    case usecase.ErrCustomerNotFound:
        c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Customer not found"})
    case usecase.ErrProductNotFound:
        c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "One of the products not found"})
    case usecase.ErrValidation:
        c.JSON(http.StatusBadRequest, gin.H{"status": "bad request", "message": "Request doesn’t pass validation"})
    case usecase.ErrPaidNotEnough:
        c.JSON(http.StatusBadRequest, gin.H{"status": "bad request", "message": "Paid amount is not enough based on all bought products"})
    case usecase.ErrChangeIncorrect:
        c.JSON(http.StatusBadRequest, gin.H{"status": "bad request", "message": "Change is incorrect based on all bought products and the amount paid"})
    case usecase.ErrProductStockInsufficient:
        c.JSON(http.StatusBadRequest, gin.H{"status": "bad request", "message": "One of the product's stock is not enough"})
    case usecase.ErrProductNotAvailable:
        c.JSON(http.StatusBadRequest, gin.H{"status": "bad request", "message": "One of the products is not available"})
    default:
        c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal server error"})
    }
    return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Successfully checked out product"})
}

func validateProduct(product dto.RequestUpsertProduct) error {
	validate := validator.New()
	if err := validate.RegisterValidation("categoryEnum", categoryEnum); err != nil {
		return err
	}
	// Perform validation
	if err := validate.Struct(product); err != nil {
		return err
	}
	return nil
}

func categoryEnum(fl validator.FieldLevel) bool {
	category := fl.Field().String()
	validCategories := map[string]bool{
		"Clothing":    true,
		"Accessories": true,
		"Footwear":    true,
		"Beverages":   true,
	}
	return validCategories[category]
}

func validateCheckoutProductRequest(req dto.CheckoutProductRequest) error {
	validate := validator.New()

	// Perform validation
	if err := validate.Struct(req); err != nil {
		return err
	}

	return nil
}

