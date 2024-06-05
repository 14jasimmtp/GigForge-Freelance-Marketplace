package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func Common(api fiber.Router, job *handler.JobsHandler, project *handler.ProjectHandler, profile *handler.ProfileHandler) {
	jobs := api.Group("/jobs")
	jobs.Get("/search", job.Search)
	jobs.Get("", job.GetJobs)
	api.Get("/projects", project.ListProjects)
	api.Get("/talents", profile.GetTalents)
	api.Get("/categories", job.GetCategories)
	api.Get("/reviews/:Fid",profile.GetFreelancerReviews)
	api.Get("/notifications",profile.GetNotifications)
}
