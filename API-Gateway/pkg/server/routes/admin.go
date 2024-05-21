package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Admin(admin fiber.Router, handler *handler.AdminHandler) {
	admin.Post("/login", handler.AdminLogin)
	admin.Post("/skills", handler.AddSkill)
	admin.Post("/user/block", handler.BlockUser)
	admin.Post("/user/unblock", handler.UnBlockUser)
	admin.Post("/category",middlewares.AuthAdmin, handler.AddCategory)
	// admin.Get("/dashboard",handler.AdminDashboard)
}
