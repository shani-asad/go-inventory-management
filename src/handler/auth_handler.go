package handler

import (
	"errors"
	"inventory-management/model/dto"
	"inventory-management/src/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	iAuthUsecase usecase.AuthUsecaseInterface
}

func NewAuthHandler(iAuthUsecase usecase.AuthUsecaseInterface) AuthHandlerInterface {
	return &AuthHandler{iAuthUsecase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request dto.RequestCreateStaff
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Register bad request")
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	// Validate request payload
	err = ValidateRegisterRequest(request.PhoneNumber, request.Name, request.Password)
	if err != nil {
		log.Println("Register bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	// Check if phoneNumber already exists
	exists, _ := h.iAuthUsecase.GetStaffByPhoneNumber(request.PhoneNumber)
	if exists {
		log.Println("Register bad request ", err)
		c.JSON(409, gin.H{"status": "bad request", "message": "phoneNumber already exists"})
		return
	}

	token, err := h.iAuthUsecase.Register(request)
	if err != nil {
		log.Println("Register bad request ", err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	log.Println("Register successful")
	c.JSON(201, gin.H{
		"message": "Staff registered successfully",
		"data": gin.H{
			"phoneNumber":       request.PhoneNumber,
			"name":        request.Name,
			"accessToken": token,
		},
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request dto.RequestAuth
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Login bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	err = ValidateLoginRequest(request.PhoneNumber, request.Password)
	if err != nil {
		log.Println("Login bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	token, userData, err := h.iAuthUsecase.Login(request)
	if err != nil {
		log.Println("Login bad request ", err)
		if err.Error() == "user not found" {
			c.JSON(404, gin.H{"status": "bad request", "message": "user not found"})
			return
		}
		if err.Error() == "wrong password" {
			c.JSON(400, gin.H{"status": "bad request", "message": "wrong password"})
			return
		}
	}

	log.Println("Login successful")
	c.JSON(200, gin.H{
		"message": "Staff logged successfully",
		"data": gin.H{
			"phoneNumber":       userData.PhoneNumber,
			"name":        userData.Name,
			"accessToken": token,
		},
	})
}

// ValidateRegisterRequest validates the register user request payload
func ValidateRegisterRequest(phoneNumber, name, password string) error {
	// Validate phoneNumber format
	if !isValidPhoneNumber(phoneNumber) {
		return errors.New("phoneNumber must be in valid phoneNumber format")
	}

	// Validate name length
	if len(name) < 5 || len(name) > 50 {
		return errors.New("name length must be between 5 and 50 characters")
	}

	// Validate password length
	if len(password) < 5 || len(password) > 15 {
		return errors.New("password length must be between 5 and 15 characters")
	}

	return nil
}

func ValidateLoginRequest(phoneNumber, password string) error {
	// Validate phoneNumber format
	if !isValidPhoneNumber(phoneNumber) {
		return errors.New("phoneNumber must be in valid phoneNumber format")
	}

	// Validate password length
	if len(password) < 5 || len(password) > 15 {
		return errors.New("password length must be between 5 and 15 characters")
	}

	return nil
}

// Helper function to validate phoneNumber format
func isValidPhoneNumber(phoneNumber string) bool {
	// Regular expression pattern for phoneNumber format
	// TODO - implement
	return false
}
