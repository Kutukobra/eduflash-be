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
			//BODY: {email, username, password}
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
			//
			room.GET("/:roomId", a.roomHandler.GetRoomById)
			//?studentName
			room.POST("/:roomId/join", a.roomHandler.JoinRoom)
			//
			room.GET("/:roomId/students", a.roomHandler.GetStudentsByRoomId)
			//
			room.GET("/:roomId/quizzes", a.roomHandler.GetQuizzesByRoomId)
		}

		quiz := api.Group("/quiz")
		{
			//BODY {content}
			quiz.POST("/create", a.quizHandler.CreateQuiz)

			quiz.GET("/:quizId", a.quizHandler.GetQuizById)
			//BODY {studentName, score}
			quiz.POST("/:quizId/submit", a.quizHandler.SubmitScore)
		}
	}
}
