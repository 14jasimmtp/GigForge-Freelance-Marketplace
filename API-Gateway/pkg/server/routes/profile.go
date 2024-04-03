package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers/auth"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Profile(profile fiber.Router,h *handler.ProfileHandler){
	profile.Post("/Education",middlewares.AuthFreelancer,h.AddEducationDetails)
	profile.Put("/Education/:id",middlewares.AuthFreelancer,h.UpdateEducation)
	profile.Delete("/Education/:id",middlewares.AuthFreelancer,h.DeleteEducation)
	profile.Post("/description",middlewares.AuthFreelancer,h.AddProfileDescription)
	profile.Put("/description",middlewares.AuthFreelancer,h.EditProfileDescription)
	// profile.Get("/freelancer/profile",middlewares.AuthFreelancer,h.GetProfile)
	// profile.Get("/client/profile",middlewares.AuthClient,h.GetProfileClient)
}