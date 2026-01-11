package app

import (
	"github.com/gin-gonic/gin"
)

func (a *App) Routes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.GET("/")
		}
	}
}
