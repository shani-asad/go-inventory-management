package handler

import (
	"cats-social/model/dto"
	"cats-social/src/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CatHandler struct {
	iCatUsecase usecase.CatUsecaseInterface
}

func NewCatHandler(iCatUsecase usecase.CatUsecaseInterface) CatHandlerInterface {
	return &CatHandler{iCatUsecase}
}

func (h *CatHandler) GetCatById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "cat by id",
	})
}

func (h *CatHandler) AddCat(c *gin.Context) {
	var request dto.RequestCreateCat
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request", "message": err})
		return
	}

	if !validateCat(&request) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request doesn’t pass validation"})
		return
	}

	err := h.iCatUsecase.AddCat(request)
	if err != nil {
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
    return
	}

	// Mocking the response
	response := gin.H{
		"message": "success",
		"data": gin.H{
			"id":        "some-id", // use whatever id
			"createdAt": time.Now().Format(time.RFC3339), // in ISO 8601 format
		},
	}

	c.JSON(http.StatusCreated, response)
}

func validateCat(cat *dto.RequestCreateCat) bool {
	if cat.Name == "" || len(cat.Name) > 30 {
		return false
	}

	validRaces := map[string]bool{
		"Persian":          true,
		"Maine Coon":       true,
		"Siamese":          true,
		"Ragdoll":          true,
		"Bengal":           true,
		"Sphynx":           true,
		"British Shorthair": true,
		"Abyssinian":       true,
		"Scottish Fold":    true,
		"Birman":           true,
	}
	if !validRaces[cat.Race] {
		return false
	}

	if cat.Sex != "male" && cat.Sex != "female" {
		return false
	}

	if cat.AgeInMonth < 1 || cat.AgeInMonth > 120082 {
		return false
	}

	if cat.Description == "" || len(cat.Description) > 200 {
		return false
	}

	if len(cat.ImageUrls) < 1 {
		return false
	}
	for _, url := range cat.ImageUrls {
		if url == "" {
			return false
		}
	}

	return true
}