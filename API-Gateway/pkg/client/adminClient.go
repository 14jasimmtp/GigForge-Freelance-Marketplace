package client

import (
	"log"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/admin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitAdminClient() admin.AdminServiceClient {
	conn, err := grpc.Dial(viper.GetString("ADMIN_SVC"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error while connecting job client : ", err)
	}

	client := admin.NewAdminServiceClient(conn)

	return client
}
