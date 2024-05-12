package handler

import (
	"fmt"
	"inventory-management/model/dto"
	"inventory-management/src/usecase"
	"log"
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

	c.JSON(statusCode, nil)
}

func validateProduct(product dto.RequestUpsertProduct) error {
	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		return err
	}
	return nil
}
