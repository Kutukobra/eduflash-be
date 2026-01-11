package main

import (
	"fmt"
	"log"

	"github.com/Kutukobra/eduflash-be/app"
	"github.com/Kutukobra/eduflash-be/app/config"
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
	app.Routes(router)

	router.Run(":" + cfg.appPort)
	fmt.Println("Pagerank Running on Port :" + cfg.appPort)
}
