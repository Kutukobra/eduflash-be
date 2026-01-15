package handler

import (
	"net/http"

	"github.com/Kutukobra/eduflash-be/app/model"
	"github.com/Kutukobra/eduflash-be/app/service"
	"github.com/gin-gonic/gin"
)

type QuizHandler struct {
	serv *service.QuizService
}

func NewQuizHandler(serv *service.QuizService) *QuizHandler {
	return &QuizHandler{serv: serv}
}

func (h *QuizHandler) CreateQuiz(c *gin.Context) {
	ctx := c.Request.Context()

	var requestData struct {
		Content []model.QuizContent `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return
	}

	quizId, err := h.serv.CreateQuiz(ctx, requestData.Content)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create quiz."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": quizId})
}
