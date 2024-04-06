package di

import (
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/client"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/config"
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/routes"
	"github.com/gofiber/fiber/v2"
)

func InitializeAuth(app *fiber.App, cfg *config.Config) {
	client := client.InitAuthClient(cfg)

	user := handler.NewAuthHandler(client)
	routes.Auth(app.Group("/auth"), user)
	profile := handler.NewProfilehandler(client)
	routes.Profile(app.Group("/profile"), profile)
}
