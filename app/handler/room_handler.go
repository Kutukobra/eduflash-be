package handler

import (
	"errors"
	"net/http"

	"github.com/Kutukobra/eduflash-be/app/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type RoomHandler struct {
	serv *service.RoomService
}

func NewRoomHandler(serv *service.RoomService) *RoomHandler {
	return &RoomHandler{
		serv: serv,
	}
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	ctx := c.Request.Context()

	var requestData struct {
		RoomName string `json:"roomName" binding:"required"`
		OwnerId  string `json:"ownerId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return
	}

	roomData, err := h.serv.CreateRoom(ctx, requestData.RoomName, requestData.OwnerId)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"room": roomData})
}

func (h *RoomHandler) JoinRoom(c *gin.Context) {
	ctx := c.Request.Context()

	room_id := c.Param("roomId")
	student_name := c.Query("studentName")

	if room_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID."})
		return
	}

	roomData, err := h.serv.JoinRoom(ctx, room_id, student_name)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name already taken."})
			return
		} else if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Room not found taken."})
			return
		}
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"room": roomData})
}

func (h *RoomHandler) GetRoomById(c *gin.Context) {
	ctx := c.Request.Context()

	room_id := c.Param("roomId")

	if room_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID."})
		return
	}

	roomData, err := h.serv.GetRoomById(ctx, room_id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Room not found"})
			return
		}
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get room data."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"room": roomData})
}

func (h *RoomHandler) GetStudentsByRoomId(c *gin.Context) {
	ctx := c.Request.Context()

	room_id := c.Param("roomId")
	if room_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID."})
		return
	}

	students, err := h.serv.GetStudentsByRoomId(ctx, room_id)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"students": students})
}

func (h *RoomHandler) GetQuizzesByRoomId(c *gin.Context) {
	ctx := c.Request.Context()

	room_id := c.Param("roomId")
	if room_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID."})
		return
	}

	quizzes, err := h.serv.GetQuizzesByRoomId(ctx, room_id)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"quizzes": quizzes})
}

func (h *RoomHandler) AddQuiz(c *gin.Context) {
	ctx := c.Request.Context()

	var requestData struct {
		QuizId string `json:"quizId" binding:"required"`
	}

	roomId := c.Param("roomId")

	if err := c.ShouldBindJSON(&requestData); err != nil || roomId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return
	}

	err := h.serv.AddQuiz(ctx, roomId, requestData.QuizId)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
