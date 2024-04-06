package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func Job(job fiber.Router, handler *handler.JobsHandler) {
	job.Post("",handler.PostJob)
	job.Put("")
	job.Post("/propose",handler.SendProposal)
	job.Post("/hire")
	job.Post("/send-offer")
	job.Post("/accept-offer",)
}
