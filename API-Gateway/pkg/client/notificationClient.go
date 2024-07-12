package client

import (
	"log"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/notification"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitNotificationClient() notification.NotificationServiceClient {
	conn, err := grpc.Dial(viper.GetString("Notification_SVC"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error while connecting job client : ", err)
	}

	client := notification.NewNotificationServiceClient(conn)

	return client
}