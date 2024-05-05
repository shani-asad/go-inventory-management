package handler

import (
	"cats-social/model/dto"
	"cats-social/src/usecase"

	"github.com/gin-gonic/gin"
)

type MatchHandler struct {
	iMatchUsecase usecase.MatchUsecaseInterface
}

func NewMatchHandler(iMatchUsecase usecase.MatchUsecaseInterface) MatchHandlerInterface {
	return &MatchHandler{iMatchUsecase}
}

func (h *MatchHandler) CreateMatch(c *gin.Context) {
	var request dto.RequestCreateMatch
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	//TODO - validate request
	//userId, _ := c.Get("user_id")

	err = h.iMatchUsecase.CreateMatch(request)
	if err != nil {
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully send match request",
	})
}

func (h *MatchHandler) GetMatch(c *gin.Context) {

	//TODO - validate request
	//userId, _ := c.Get("user_id") ---> validate 

	userId, _  := c.Get("user_id")
	response, err := h.iMatchUsecase.GetMatch(userId.(int))
	if err != nil {
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully get match requests",
		"data": response,
	})
}
