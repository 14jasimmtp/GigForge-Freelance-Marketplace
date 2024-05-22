package main

import (
	"fmt"
	"log"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/config"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/di"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Print("error loading configurations : ", err)
	}

	app := fiber.New()
	app.Static("/","/home/jasim/GigForge-Freelance-Marketplace/API-Gateway/template")
	di.InitializeAPI(app, cfg)
	fmt.Println("listening on port 3000")
	if err :=app.Listen(cfg.PORT);err != nil {
		log.Fatal(err)
	}
	fmt.Println("started")
}
