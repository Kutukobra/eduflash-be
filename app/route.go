package app

import (
	"github.com/gin-gonic/gin"
)

func (a *App) Routes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			//?email
			user.GET("/", a.userHandler.GetUserByEmail)
			//BODY: {email, username, password, role}
			user.POST("/register", a.userHandler.RegisterUser)
			//BODY: {email, password}
			user.POST("/login", a.userHandler.LoginUser)
			//
			user.GET("/:ownerId/rooms", a.userHandler.GetRoomsByOwnerId)
		}

		room := api.Group("/room")
		{
			//BODY: {roomName, ownerId}
			room.POST("/create", a.roomHandler.CreateRoom)
			//?studentName
			room.POST("/:roomId/join", a.roomHandler.JoinRoom)
			//
			room.GET("/:roomId/students", a.roomHandler.GetStudentsByRoomId)
		}
	}
}
