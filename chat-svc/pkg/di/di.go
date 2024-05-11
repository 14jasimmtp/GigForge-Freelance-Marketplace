package di

import (
	"log"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pkg/Infrastructure/repository"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pkg/Infrastructure/service"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pkg/db"
)


func InjectDependencies() *service.Service{
	mongoColl,err:=db.ConnectMongoDB()
	if err != nil {
		log.Fatal("mongo",err)
	}
	chatRepo:=repository.NewRepository(mongoColl)
	chatService:=service.NewChatService(chatRepo)

	go chatService.ChatReciever()

	return chatService
}