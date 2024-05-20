package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Chat(api fiber.Router,chat *handler.ChatHandler){
	api.Get("",middlewares.AuthFreelancer,websocket.New(chat.Chat))
	api.Get("/messages/:receiver_id",middlewares.AuthChat,chat.GetMessages)
	// api.Delete("/messages/:message_id",chat.DeleteMessage)
}
