package client

import (
	"log"
	"fmt"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/auth"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitAuthClient(c *config.Config) auth.AuthServiceClient {
	fmt.Println("workig new image ")
	conn, err := grpc.Dial(c.AUTH_SVC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connecting auth svc : ", err)
	}

	client := auth.NewAuthServiceClient(conn)

	return client
}
