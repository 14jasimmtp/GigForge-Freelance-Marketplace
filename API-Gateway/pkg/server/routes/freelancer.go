package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func Freelancer(
	api fiber.Router,
	profile *handler.ProfileHandler,
	project *handler.ProjectHandler,
	job *handler.JobsHandler,
){}

