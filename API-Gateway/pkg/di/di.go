package di

import (
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/client"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/config"
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/routes"
	"github.com/gofiber/fiber/v2"
)

func InitializeAPI(app *fiber.App, cfg *config.Config) {
	//clients
	Authclient := client.InitAuthClient(cfg)
	JobsClient := client.InitJobClient()
	adminClient:=client.InitAdminClient()
	ProjectClient:=client.InitProjectClient()

	//handler
	auth := handler.NewAuthHandler(Authclient)
	profile := handler.NewProfilehandler(Authclient)
	jobs:= handler.NewJobsHandler(JobsClient)
	admins:=handler.NewAdminHandler(adminClient,Authclient)
	project:=handler.NewProjectHandler(ProjectClient)

	//routes
	routes.Freelancer(app.Group("/freelancer"),profile,project,jobs)
	routes.Admin(app.Group("/admin"),&admins)
	routes.Client(app.Group("/client"),profile,project,jobs)
	routes.Auth(app.Group("/auth"),auth)
}
