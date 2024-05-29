package main

import (
	"fmt"
	"log"
	"net"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/di"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pb"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pkg/config"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Println(err)
	}
	lis, err := net.Listen("tcp", viper.GetString("PORT"))
	if err != nil {
		log.Fatal("error", err)
	}

	svc := di.InitializeAPI()
	grpcServer := grpc.NewServer()
	pb.RegisterProjectServiceServer(grpcServer, &svc)
	fmt.Println(lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
