package client

import (
	"log"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitChatClient() chat.ChatServiceClient {
	conn, err := grpc.Dial("localhost:30005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error while connecting job client : ", err)
	}

	client := chat.NewChatServiceClient(conn)

	return client
}
