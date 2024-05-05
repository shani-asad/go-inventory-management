package handler

import (
	"cats-social/model/dto"
	"cats-social/src/usecase"
	"errors"
	"log"
	"regexp"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	iAuthUsecase usecase.AuthUsecaseInterface
}

func NewAuthHandler(iAuthUsecase usecase.AuthUsecaseInterface) AuthHandlerInterface {
	return &AuthHandler{iAuthUsecase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request dto.RequestCreateUser
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Register bad request")
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	// Validate request payload
	err = ValidateRegisterRequest(request.Email, request.Name, request.Password)
	if err != nil {
		log.Println("Register bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	// Check if email already exists
	exists, _ := h.iAuthUsecase.GetUserByEmail(request.Email)
	if exists {
		log.Println("Register bad request ", err)
		c.JSON(409, gin.H{"status": "bad request", "message": "email already exists"})
		return
	}

	err = h.iAuthUsecase.Register(request)
	if err != nil {
		log.Println("Register bad request ", err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	log.Println("Register successful")
	c.JSON(200, gin.H{
		"message": "user register success",
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

	err = ValidateLoginRequest(request.Email, request.Password)
	if err != nil {
		log.Println("Login bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	token, err := h.iAuthUsecase.Login(request)
	if err != nil {
		log.Println("Login bad request ", err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	log.Println("Login successful")
	c.JSON(200, gin.H{
		"token": token,
	})
}

// ValidateRegisterRequest validates the register user request payload
func ValidateRegisterRequest(email, name, password string) error {
	// Validate email format
	if !isValidEmail(email) {
		return errors.New("email must be in valid email format")
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

func ValidateLoginRequest(email, password string) error {
	// Validate email format
	if !isValidEmail(email) {
		return errors.New("email must be in valid email format")
	}

	// Validate password length
	if len(password) < 5 || len(password) > 15 {
		return errors.New("password length must be between 5 and 15 characters")
	}

	return nil
}

// Helper function to validate email format
func isValidEmail(email string) bool {
	// Regular expression pattern for email format
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}
