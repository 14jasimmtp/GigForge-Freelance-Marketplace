package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Client(api fiber.Router, profile *handler.ProfileHandler,
	project *handler.ProjectHandler,
	job *handler.JobsHandler) {
	jobs := api.Group("/job")

	jobs.Post("", middlewares.AuthAdmin, job.AcceptOffer)
	jobs.Post("", middlewares.AuthClient, job.PostJob)
	jobs.Put("", middlewares.AuthClient, job.EditJob)
	jobs.Get("", middlewares.AuthClient, job.GetMyJobs)
	jobs.Post("/send-offer", middlewares.AuthClient, job.SendOffer)

	jobs.Get("/proposals/:id",middlewares.AuthClient,job.GetProposalsOfJob)

	contract:=api.Group("/contracts")
	contract.Use(middlewares.AuthClient)
	// contract.Get("",job.GetAllContractsForClient)
	// contract.Get("/:id",job.GetOneContract)
	//contract.Get("/invoices/:job_id",job.GetAllInvoicesOfAJob)
	contract.Get("/payment/:invoice_id",job.GetPaymentForContractWithInvoiceID)

	projects:=api.Group("/project")
	projects.Use(middlewares.AuthClient)
	projects.Get("",project.ListProjects)
	projects.Get("/:id",project.ListProjectWithID)
	projects.Post("/buy/:id",project.BuyProject)
	// projects.Post("/payment/:order_id",project.ExecutePaymentForProject)

	// profiles:=api.Group("/profile")
	// profile
	// profiles.Get("",profile.GetProfile)
}
