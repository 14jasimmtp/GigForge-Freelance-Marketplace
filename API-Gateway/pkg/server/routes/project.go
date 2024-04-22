package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func Project(project fiber.Router, h *handler.ProjectHandler) {
	project.Post("/add",h.AddProject)
}
