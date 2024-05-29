package client

import (
	"fmt"
	"log"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/auth"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/config"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitAuthClient(c *config.Config) auth.AuthServiceClient {
	fmt.Println("workig new image ")
	conn, err := grpc.Dial(viper.GetString("AUTH_SVC"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connecting auth svc : ", err)
	}

	client := auth.NewAuthServiceClient(conn)

	return client
}
