package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func Common(api fiber.Router,job *handler.JobsHandler){
	jobs:=api.Group("/jobs")
	jobs.Get("/search",job.Search)
	jobs.Get("",job.GetJobs)
}