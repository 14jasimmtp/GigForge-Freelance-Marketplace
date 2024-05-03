package client

import (
	"log"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pb/user"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitJobClient() user.JobserviceClient {
	conn, err := grpc.Dial(viper.GetString("USER_SVC"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error while connecting job client : ", err)
	}

	client := user.NewJobserviceClient(conn)

	return client
}
