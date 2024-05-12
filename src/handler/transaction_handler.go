package handler

import (
	"fmt"
	"inventory-management/model/dto"
	"inventory-management/src/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	iTransactionUsecase usecase.TransactionUsecaseInterface
}

func NewTransactionHandler(iTransactionUsecase usecase.TransactionUsecaseInterface) TransactionHandlerInterface {
	return &TransactionHandler{iTransactionUsecase}
}

func (h *TransactionHandler) GetTransactions(c *gin.Context){
	var customerId string
	
	if _, ok := c.Request.URL.Query()["customerId"]; ok{
		customerId = c.Query("customerId")
	}

	var limit, offset int
	if _, ok := c.Request.URL.Query()["limit"]; ok && c.Query("limit") != "" {
		val, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
				limit = 5
		} else {
			limit = val
		}
	} else {
		limit = 5
	}

	if _, ok := c.Request.URL.Query()["offset"]; ok && c.Query("offset") != "" {
		val, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
				offset = 0
		} else {
			offset = val
		}
	}

	var createdAt  string
	
	if _, ok := c.Request.URL.Query()["createdAt"]; ok && c.Query("createdAt") != "" {
		createdAt  = c.Query("createdAt")
	}

	params := dto.GetTransactionRequest{
		CustomerId	: customerId,
		Limit	: limit,
		Offset	: offset,
		CreatedAt	: createdAt,
	}

	fmt.Println("paramsCustomerId>>>>>", params.CustomerId)
	fmt.Println("paramsLimit>>>>>", params.Limit)
	fmt.Println("paramsOffset>>>>>", params.Offset)
	fmt.Println("paramsCreatedAt>>>>>", params.CreatedAt)

	transactionList, err := h.iTransactionUsecase.GetTransactions(params)

	if err != nil {
		log.Println("get sku server error ", err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	if(len(transactionList) < 1) { transactionList = []dto.TransactionData{}}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": transactionList})

}

