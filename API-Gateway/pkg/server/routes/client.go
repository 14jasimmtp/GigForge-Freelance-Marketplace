package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Client(api fiber.Router, profile *handler.ProfileHandler,
	project *handler.ProjectHandler,
	job *handler.JobsHandler) {

	profiles:= api.Group("/profile")	

	profiles.Get("",middlewares.AuthClient,profile.GetClientProfile)
	profiles.Put("/company-details",middlewares.AuthClient,profile.UpdateCompanyDetails)
	profiles.Put("/company-contacts",middlewares.AuthClient,profile.UpdateCompanyContacts)
	jobs := api.Group("/job")

	jobs.Post("", middlewares.AuthClient, job.PostJob)
	jobs.Put("/:jobID", middlewares.AuthClient, job.EditJob)
	jobs.Get("", middlewares.AuthClient, job.GetMyJobs)
	jobs.Post("/send-offer", middlewares.AuthClient, job.SendOffer)

	jobs.Get("/proposals/:id",middlewares.AuthClient,job.GetProposalsOfJob)

	contract:=api.Group("/contracts")
	contract.Get("",middlewares.AuthClient,job.GetAllContractsForClient)
	contract.Get("/:id",middlewares.AuthClient,job.GetOneContract)
	contract.Get("/invoices/:contractID",middlewares.AuthClient,middlewares.AuthClient,job.GetAllInvoicesOfContract)
	contract.Post("/attachment",middlewares.AuthChat,job.AddContractAttachment)
	contract.Get("/attachment/:contractID",middlewares.AuthChat,job.GetAttachments)
	

	projects:=api.Group("/projects")
	projects.Get("",project.ListProjects)
	projects.Get("/:id",project.ListProjectWithID)
	projects.Post("/order/:id",middlewares.AuthClient,project.BuyProject)
	// projects.Get("/orders",middlewares.AuthClient,project.GetProjectOrdersForClient)
	
	
	api.Post("/review-freelancer",middlewares.AuthClient,profile.ReviewFreelancer)

	api.Post("/payment/project/execute",project.ExecutePaymentProject)
	api.Get("/payment/project/order",project.GetPaymentProject)
	api.Post("/payment/project/capture",project.CapturePaymentProject)
	api.Get("/payment/contract",job.GetPaymentContract)
	api.Post("/payment/contract/execute",job.ExecutePaymentContract)
	api.Post("/payment/contract/capture",job.CapturePaymentContract)
}
