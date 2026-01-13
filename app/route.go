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
		}

		room := api.Group("/room")
		{
			//?ownerId
			room.POST("/create", a.roomHandler.CreateRoom)
		}
	}
}
