package main

import (
	"fmt"
	"log"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/config"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/di"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Print("error loading configurations : ", err)
	}

	app := fiber.New()
	app.Static("/template","./template")
    app.Use(logger.New())

	di.InitializeAPI(app)
	fmt.Println("listening on port 3000")
	if err :=app.Listen(viper.GetString("PORT"));err != nil {
		log.Fatal(err)
	}
	fmt.Println("started")
}
