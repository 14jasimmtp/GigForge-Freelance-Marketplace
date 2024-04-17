package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Admin(admin fiber.Router, handler *handler.AdminHandler) {
	admin.Post("/login", handler.AdminLogin)
	admin.Post("/profile/skills",middlewares.AuthAdmin, handler.AddSkill)
	admin.Post("/user/block",middlewares.AuthAdmin, handler.BlockUser)
	admin.Post("/user/unblock",middlewares.AuthAdmin, handler.UnBlockUser)
}
