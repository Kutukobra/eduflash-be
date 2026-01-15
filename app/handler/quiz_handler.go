package handler

import (
	"errors"
	"net/http"

	"github.com/Kutukobra/eduflash-be/app/model"
	"github.com/Kutukobra/eduflash-be/app/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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
		RoomId  string              `json:"roomId" binding:"required"`
		Content []model.QuizContent `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return
	}

	quizId, err := h.serv.CreateQuiz(ctx, requestData.RoomId, requestData.Content)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": quizId})
}

func (h *QuizHandler) GetQuizById(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("quizId")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must have ID."})
		return
	}

	content, err := h.serv.GetQuizById(ctx, id)
	if err != nil {
		c.Error(err)
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID."})
			return
		}

		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"quiz": content})
}

func (h *QuizHandler) SubmitScore(c *gin.Context) {
	ctx := c.Request.Context()

	var requestData struct {
		StudentName string  `json:"studentName" binding:"required"`
		Score       float32 `json:"score" binding:"required"`
	}

	quizId := c.Param("quizId")

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}

	err := h.serv.SubmitScore(ctx, quizId, requestData.StudentName, requestData.Score)

	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
