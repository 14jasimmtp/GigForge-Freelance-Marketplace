package main

import (
	"log"
	"net"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pb"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pkg/di"
	"google.golang.org/grpc"
)

func main() {
	// c, err := config.LoadConfig()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	lis, err := net.Listen("tcp", ":30005")
	if err != nil {
		log.Fatal("error", err)
	}

	svc := di.InjectDependencies()

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer,svc)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
