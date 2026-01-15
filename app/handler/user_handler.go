package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/Kutukobra/eduflash-be/app/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	serv *service.UserService
}

func NewUserHandler(serv *service.UserService) *UserHandler {
	return &UserHandler{
		serv: serv,
	}
}

// /user?email=email
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	ctx := c.Request.Context()

	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email."})
		return
	}

	userData, err := h.serv.GetUserByEmail(ctx, email)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userData})
}

func (h *UserHandler) GetRoomsByOwnerId(c *gin.Context) {
	ctx := c.Request.Context()

	owner_id := c.Param("ownerId")
	if owner_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid owner ID."})
		return
	}

	rooms, err := h.serv.GetRoomsByOwnerId(ctx, owner_id)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Owner."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	ctx := c.Request.Context()

	var requestData struct {
		Email    string `json:"email" binding:"required,email"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.Error(err)
		log.Println(requestData)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return
	}

	userData, err := h.serv.RegisterUser(ctx, requestData.Email, requestData.Username, requestData.Password, requestData.Role)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": userData.ID})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	ctx := c.Request.Context()

	var requestData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return
	}

	userData, err := h.serv.LoginUser(ctx, requestData.Email, requestData.Password)
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password."})
		return
	} else if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found."})
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userData})
}
