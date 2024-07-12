package main

import (
	"log"
	"net"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/notification-svc/pb"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/notification-svc/pkg/di"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatal(err)
	}
	service:=di.InjectDependencies()
	RPCServer := grpc.NewServer()
	pb.RegisterNotificationServiceServer(RPCServer,service)
	if err := RPCServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
