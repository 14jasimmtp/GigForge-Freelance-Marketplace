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
	{
		jobs.Post("", middlewares.AuthAdmin, job.AcceptOffer)
	}
	jobs.Post("", middlewares.AuthClient, job.PostJob)
	jobs.Put("", middlewares.AuthClient, job.EditJob)
	jobs.Get("", middlewares.AuthClient, job.GetMyJobs)
	jobs.Post("send-offer", middlewares.AuthClient, job.SendOffer)
	// jobs.Get("/proposal/:jobID",job.GetProposals)
	// jobs.Get("",job.)

	// profiles :=api.Group("/profile")
	// // profiles.Post("/",profile.)

	// profile
}
