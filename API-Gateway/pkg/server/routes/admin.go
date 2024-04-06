package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func Admin(admin fiber.Router, handler *handler.AdminHandler) {
	admin.Post("/login", handler.AdminLogin)
	admin.Post("/profile/skills", handler.AddSkill)
	admin.Post("/user/block", handler.BlockUser)
	admin.Post("/user/unblock", handler.UnBlockUser)
}
