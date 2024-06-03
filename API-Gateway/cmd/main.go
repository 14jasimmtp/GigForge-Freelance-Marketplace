package main

import (
	"fmt"
	"log"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/docs"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/config"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/di"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
)

// @title          GigForge
// @version        1.0
// @description    Freelance Marketplace.
// @termsOfService http://swagger.io/terms/

// @host     gigforge.jasim.online
// @BasePath /

// @securityDefinitions.apikey FreelancerAccessToken
// @in                         header
// @name                       Authorization
// @securityDefinitions.apikey ClientAccessToken
// @in                         header
// @name                       Authorization
// @securityDefinitions.apikey AdminAccessToken
// @in                         header
// @name                       Authorization

func main() {

	docs.SwaggerInfo.Host = "gigforge.jasim.online"
	err := config.LoadConfig()
	if err != nil {
		log.Print("error loading configurations : ", err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Static("/template", "./template")
	app.Use(logger.New())

	di.InitializeAPI(app)
	fmt.Println("listening on port ")
	if err := app.Listen(viper.GetString("PORT")); err != nil {
		log.Fatal(err)
	}
	fmt.Println("started")
}
