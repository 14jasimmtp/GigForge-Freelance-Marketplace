package main

import (
	"fmt"
	"log"
	"net"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/auth"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/job"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/project"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/config"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/di"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("new image with certificat")
	err := config.LoadConfig()
	if err != nil {
		log.Println("error loading configs")
	}
	lis, err := net.Listen("tcp", viper.GetString("PORT"))
	if err != nil {
		log.Fatal("error", err)
	}

	svc,repo := di.InitializeAPI()
	grpcServer := grpc.NewServer()

	auth.RegisterAuthServiceServer(grpcServer, svc)
	job.RegisterJobserviceServer(grpcServer,repo)
	project.RegisterUserServiceServer(grpcServer,repo)

	fmt.Println(lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
