package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/gofiber/websocket/v2"
	"github.com/gofiber/fiber/v2"
)

func Chat(api fiber.Router,chat *handler.ChatHandler){
	api.Get("/chat/:chatid",websocket.New(chat.Chat))
	
}
