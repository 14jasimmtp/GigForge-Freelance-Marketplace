package di

import (
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/client"
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/routes"
	"github.com/gofiber/fiber/v2"
)

func InitializeAPI(app *fiber.App) {
	//clients
	Authclient := client.InitAuthClient()
	JobsClient := client.InitJobClient()
	ProjectClient := client.InitProjectClient()
	chatClient := client.InitChatClient()

	//handler
	auth := handler.NewAuthHandler(Authclient)
	profile := handler.NewProfilehandler(Authclient)
	jobs := handler.NewJobsHandler(JobsClient)
	admins := handler.NewAdminHandler(JobsClient, Authclient)
	project := handler.NewProjectHandler(ProjectClient)
	chat := handler.NewChatHandler(chatClient)

	//routes
	routes.Freelancer(app.Group("/freelancer"), profile, project, jobs)
	routes.Admin(app.Group("/admin"), &admins)
	routes.Client(app.Group("/client"), profile, project, jobs)
	routes.Auth(app.Group("/auth"), auth)
	routes.Chat(app.Group("/chat"), &chat)
	routes.Common(app.Group(""), jobs, project, profile)
}
