package main

import (
	"fmt"
	"log"
	"net"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/auth"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/config"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/di"
	"google.golang.org/grpc"
)

func main() {
	c, _ := config.LoadConfig()
	lis, err := net.Listen("tcp", c.PORT)
	if err != nil {
		log.Fatal("error", err)
	}

	svc := di.InitializeAPI(c)
	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, svc)
	fmt.Println(lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
