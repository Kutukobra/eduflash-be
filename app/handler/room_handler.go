package handler

import (
	"net/http"
	"strings"

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

	owner_id := strings.ToUpper(c.Query("ownerId"))
	if owner_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid owner ID."})
		return
	}

	roomData, err := h.serv.CreateRoom(ctx, owner_id)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roomData})
}

func (h *RoomHandler) JoinRoom(c *gin.Context) {
	ctx := c.Request.Context()

	room_id := c.Query("roomId")
	if room_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID."})
		return
	}

	err := h.serv.JoinRoom(ctx, room_id)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}
