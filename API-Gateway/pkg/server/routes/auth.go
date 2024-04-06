package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func Auth(auth fiber.Router, handler *handler.Handler) {
	auth.Post("/login", handler.Login)
	auth.Post("/signup", handler.Signup)
	auth.Post("/verify", handler.Verify)
	auth.Post("/forgot-password", handler.ForgotPassword)
	auth.Post("/reset-password", handler.ResetPassword)
}
