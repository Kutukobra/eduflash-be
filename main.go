package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Kutukobra/eduflash-be/app"
	"github.com/Kutukobra/eduflash-be/app/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	app, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	app.Routes(router)

	router.Run(":" + cfg.AppPort)
	fmt.Println("Pagerank Running on Port :" + cfg.AppPort)
}
