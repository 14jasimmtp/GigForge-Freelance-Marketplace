package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Project(project fiber.Router, h *handler.ProjectHandler) {
	project.Post("/add",middlewares.AuthFreelancer,h.AddSingleProject)
	project.Post("/addtiered",middlewares.AuthFreelancer,h.AddTieredProject)
}
