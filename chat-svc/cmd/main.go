package main

import (
	"fmt"
	"log"
	"net"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pb"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pkg/config"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pkg/di"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(viper.GetString("PORT"), "port")
	lis, err := net.Listen("tcp", viper.GetString("PORT"))
	if err != nil {
		log.Fatal("error", err)
	}

	svc := di.InjectDependencies()

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, svc)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
