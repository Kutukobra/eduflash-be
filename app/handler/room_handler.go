package handler

import (
	"net/http"

	"github.com/Kutukobra/eduflash-be/app/service"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roomData})
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
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room or name."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roomData})
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

	c.JSON(http.StatusOK, gin.H{"data": students})
}
