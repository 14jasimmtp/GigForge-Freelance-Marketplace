package di

import (
	"log"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pkg/Infrastructure/repository"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pkg/Infrastructure/service"
	broker "github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pkg/consumer"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pkg/db"
)

func InjectDependencies() *service.Service {
	mongoColl, err := db.ConnectMongoDB()
	if err != nil {
		log.Fatal("mongo", err)
	}
	chatRepo := repository.NewRepository(mongoColl)
	AMQPConn := broker.ConnectAMQP()
	chatService := service.NewChatService(chatRepo, AMQPConn)

	go chatService.ChatReciever()

	return chatService
}
