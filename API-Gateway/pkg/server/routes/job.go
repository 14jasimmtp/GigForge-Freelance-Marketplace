package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Job(job fiber.Router, handler *handler.JobsHandler) {
	job.Post("",middlewares.AuthClient,handler.PostJob)
	// job.Put("")
	job.Post("/proposal",middlewares.AuthFreelancer,handler.SendProposal)
	// job.Post("/hire")
	// job.Post("/send-offer")
	// job.Post("/accept-offer",)
}
