package client

import (
	"log"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pb/user"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitProjectClient() user.UserServiceClient {
	conn, err := grpc.Dial(viper.GetString("USER_SVC"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error while connecting job client : ", err)
	}

	client := user.NewUserServiceClient(conn)

	return client
}
