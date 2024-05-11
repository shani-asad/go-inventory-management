package handler

import (
	"errors"
	"inventory-management/model/dto"
	"inventory-management/src/usecase"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type SkuHandler struct {
	iSkuUsecase usecase.SkuUsecaseInterface
}

func NewSkuHandler(iSkuUsecase usecase.SkuUsecaseInterface) SkuHandlerInterface {
	return &SkuHandler{iSkuUsecase}
}

func (h *SkuHandler) Search(c *gin.Context) {	
	var limit, offset int
	
	if _, ok := c.Request.URL.Query()["limit"]; ok && c.Query("limit") != "" {
		_, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
				limit = 5
		}
	}

	if _, ok := c.Request.URL.Query()["offset"]; ok && c.Query("offset") != "" {
		_, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
				offset = 0
		}
	}
	
	var name string
	
	if _, ok := c.Request.URL.Query()["name"]; !ok{
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": []string{} })
		return
	} else {
		name = c.Query("offset")
	}
	
	var category, sku, price string

	if _, ok := c.Request.URL.Query()["category"]; ok && c.Query("category") != "" {
		category = c.Query("category")
	}

	if _, ok := c.Request.URL.Query()["sku"]; ok && c.Query("sku") != "" {
		sku = c.Query("sku")
	}

	if _, ok := c.Request.URL.Query()["price"]; ok && c.Query("price") != "" {
		price = c.Query("price")
	}

	var inStock bool
	isInstockValid := true

	if _, ok := c.Request.URL.Query()["inStock"]; ok && c.Query("inStock") != "" {
		if(c.Query("inStock") == "true") {
			inStock = true
		} else if(c.Query("inStock") == "false") {
			inStock = false
		} else {
			isInstockValid = false
		}
		
	}

	params := dto.SearchSkuParams{
		Limit	: limit,
		Offset	: offset,
		Name	: name,
		Category	: category,
		Sku	: sku,
		Price	: price,
		InStock	: inStock,
		IsInstockValid: isInstockValid,
	}


	skuList, err := h.iSkuUsecase.Search(params)
	if err != nil {
		log.Println("get sku server error ", err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	log.Println("get cat successful")
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": skuList})
}
