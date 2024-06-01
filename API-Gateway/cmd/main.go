package main

import (
	"fmt"
	"log"

	_"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/docs"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/config"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/di"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title          GigForge
// @version        1.0
// @description    Freelance Marketplace.
// @termsOfService http://swagger.io/terms/

// // @host     gigforge.jasim.online
// @BasePath /

// @securityDefinitions.apikey UserAuthorization
// @in                         header
// @name                       AccessToken
// @securityDefinitions.apikey UserConfirmToken
// @in                         header
// @name                       ConfirmToken
// @securityDefinitions.apikey AdminAutherisation
// @in                         header
// @name                       AccessToken

func main() {
	// docs.SwaggerInfo.Host = "gigforge.jasim.online"
	err := config.LoadConfig()
	if err != nil {
		log.Print("error loading configurations : ", err)
	}

	app := fiber.New()


	// engin.Use(middleware.CORS())

	app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Static("/template", "./template")
	app.Use(logger.New())

	di.InitializeAPI(app)
	fmt.Println("listening on port ")
	if err := app.Listen(viper.GetString("PORT")); err != nil {
		log.Fatal(err)
	}
	fmt.Println("started")
}
