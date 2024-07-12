package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Admin(admin fiber.Router, handler *handler.AdminHandler) {
	admin.Post("/login", handler.AdminLogin)
	admin.Post("/skills", middlewares.AuthAdmin, handler.AddSkill)
	admin.Put("/user/block", middlewares.AuthAdmin, handler.BlockUser)
	admin.Put("/user/unblock", middlewares.AuthAdmin, handler.UnBlockUser)
	admin.Post("/category", middlewares.AuthAdmin, middlewares.AuthAdmin, handler.AddCategory)
	admin.Get("/Contract-dashboard", middlewares.AuthAdmin, handler.AdminContractDashboard)
}
