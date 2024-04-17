package di

import (
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/client"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/config"
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/routes"
	"github.com/gofiber/fiber/v2"
)

func InitializeAPI(app *fiber.App, cfg *config.Config) {
	Authclient := client.InitAuthClient(cfg)

	user := handler.NewAuthHandler(Authclient)
	routes.Auth(app.Group("/auth"), user)
	profile := handler.NewProfilehandler(Authclient)
	routes.Profile(app.Group("/profile"), profile)

	JobsClient := client.InitJobClient()
	jobs:= handler.NewJobsHandler(JobsClient)
	routes.Job(app.Group("/job"),jobs)

	adminClient:=client.InitAdminClient()

	admin:=handler.NewAdminHandler(adminClient,Authclient)
	routes.Admin(app.Group("/admin"),&admin)
}
