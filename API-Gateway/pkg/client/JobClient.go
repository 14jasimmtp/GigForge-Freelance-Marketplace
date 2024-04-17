package client

import (
	"log"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/Job"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitJobClient() Job.JobServiceClient {
	conn, err := grpc.Dial(viper.GetString("Job_SVC"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error while connecting job client : ", err)
	}

	client := Job.NewJobServiceClient(conn)

	return client
}
