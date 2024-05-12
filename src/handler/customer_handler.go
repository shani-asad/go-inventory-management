// handler.go

package handler

import (
	"log"
	"net/http"

	"inventory-management/model/dto"
	"inventory-management/src/usecase"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
    iCustomerUsecase usecase.CustomerUsecaseInterface
}

func NewCustomerHandler(iCustomerUsecase usecase.CustomerUsecaseInterface) CustomerHandlerInterface {
    return &CustomerHandler{iCustomerUsecase}
}

func (h *CustomerHandler) RegisterCustomer(c *gin.Context) {
		var request dto.RegisterCustomerRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			log.Println("register customer bad request ", err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "bad request", "message": err})
			return
		}

		if !validateCustomer(&request) {
			log.Println("register customer bad request")
			c.JSON(http.StatusBadRequest, gin.H{"error": "request doesn’t pass validation"})
			return
		}

		// Check if phoneNumber already exists
		exists, _ := h.iCustomerUsecase.GetCustomerByPhoneNumber(request.PhoneNumber)
		if exists {
			log.Println("phoneNumber already exists ")
			c.JSON(409, gin.H{"status": "bad request", "message": "phoneNumber already exists"})
			return
		}

    userID, err := h.iCustomerUsecase.RegisterCustomer(request)
    if err != nil {
				log.Println("register customer internal server error:", err)
				c.JSON(500, gin.H{"status": "internal server error", "message": err})
        return
    }

    log.Println("register customer success")
		// Mocking the response
		response := gin.H{
			"message": "success",
			"data": gin.H{
				"userId":        userID,
				"phoneNumber": request.PhoneNumber,
				"name": request.Name,
			},
		}

		c.JSON(http.StatusCreated, response)
}

func (h *CustomerHandler) SearchCustomers(c *gin.Context) {
		var request dto.SearchCustomersRequest;
    phoneNumber := c.Query("phoneNumber")
    name := c.Query("name")

		request.PhoneNumber = phoneNumber
		request.Name = name

    customers, err := h.iCustomerUsecase.SearchCustomers(request)
		if err != nil {
			log.Println("search customer server error ", err)
			c.JSON(500, gin.H{"status": "internal server error", "message": err})
			return
		}

		log.Println("search customer successful")
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": customers})
}

func validateCustomer(customer *dto.RegisterCustomerRequest) bool {
	// Validate phoneNumber format
	if !isValidPhoneNumber(customer.PhoneNumber) {
		return false
	}

	// Validate name length
	if len(customer.Name) < 5 || len(customer.Name) > 50 {
		return false
	}

	return true
}

