package handler

import (
	"cats-social/model/dto"
	"cats-social/src/usecase"
	"database/sql"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
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

	userId, _ := c.Get("user_id")
	fmt.Println("Type of userId.(int): ", reflect.TypeOf(userId))
	request.UserId = userId.(int)
	id, err := h.iCatUsecase.AddCat(request)
	if err != nil {
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
    return
	}

	// Mocking the response
	response := gin.H{
		"message": "success",
		"data": gin.H{
			"id":        id, // use whatever id
			"createdAt": time.Now().Format(time.RFC3339), // in ISO 8601 format
		},
	}

	c.JSON(http.StatusCreated, response)
}

func (h *CatHandler) GetCat(c *gin.Context) {
	// Initialize an empty map to store valid query parameters
	queryParams := make(map[string]interface{})

	// Parse query parameters
	if id := c.Query("id"); id != "" {
			queryParams["id"] = id
	}
	if limitStr := c.Query("limit"); limitStr != "" {
			limit, err := strconv.Atoi(limitStr)
			if err == nil {
					queryParams["limit"] = limit
			}
	}
	if offsetStr := c.Query("offset"); offsetStr != "" {
			offset, err := strconv.Atoi(offsetStr)
			if err == nil {
					queryParams["offset"] = offset
			}
	}
	if race := c.Query("race"); race != "" {
			queryParams["race"] = race
	}
	if sex := c.Query("sex"); sex != "" {
			queryParams["sex"] = sex
	}
	if hasMatchedStr := c.Query("hasMatched"); hasMatchedStr != "" {
			hasMatched, err := strconv.ParseBool(hasMatchedStr)
			if err == nil {
					queryParams["hasMatched"] = hasMatched
			}
	}
	if ageInMonth := c.Query("ageInMonth"); ageInMonth != "" {
			queryParams["ageInMonth"] = ageInMonth
	}
	if ownedStr := c.Query("owned"); ownedStr != "" {
			owned, err := strconv.ParseBool(ownedStr)
			if err == nil {
					queryParams["owned"] = owned
			}
	}
	if search := c.Query("search"); search != "" {
			queryParams["search"] = search
	}

	// Create a new struct that matches the keys and types of the queryParams map
	var request dto.RequestGetCat;
	if len(queryParams) > 0 {
    request = dto.RequestGetCat{}

    if id, ok := queryParams["id"].(string); ok {
        request.Id = id
    }

    if limit, ok := queryParams["limit"].(int); ok {
        request.Limit = limit
    }

    if offset, ok := queryParams["offset"].(int); ok {
        request.Offset = offset
    }

		if race, ok := queryParams["race"].(string); ok {
        request.Race = race
    }

		if sex, ok := queryParams["sex"].(string); ok {
        request.Sex = sex
    }

		if hasMatched, ok := queryParams["hasMatched"].(bool); ok {
        request.HasMatched = hasMatched
    }

		if ageInMonth, ok := queryParams["ageInMonth"].(string); ok {
			request.AgeInMonth = ageInMonth
		}

		if owned, ok := queryParams["owned"].(bool); ok {
        request.Owned = owned
    }

		if search, ok := queryParams["search"].(string); ok {
        request.Search = search
    }
	}
	cats, err := h.iCatUsecase.GetCat(request)
	if err != nil {
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": cats})
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

func (h *CatHandler) UpdateCat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
	}

	// Parse request body
	var cat dto.RequestCreateCat
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.iCatUsecase.GetCatById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	// Check if the cat has already been requested to match
	hasMatched, err := h.iCatUsecase.CheckHasMatch(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check cat match status"})
		return
	}
	if hasMatched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cat is already requested to match"})
		return
	}

	// Validate the updated cat data
	if !validateCat(&cat) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request doesn’t pass validation"})
		return
	}

	err = h.iCatUsecase.UpdateCat(cat, int64(id))
    if err != nil {
        c.JSON(500, gin.H{"status": "internal server error", "message": err})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "Successfully add cat"})

}

func (h *CatHandler) DeleteCat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
	}

	err = h.iCatUsecase.GetCatById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	err = h.iCatUsecase.DeleteCat(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete cat from the database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully delete cat"})
}