package main

import (
	"fmt"
	"log"
	"net"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pb/job"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/config"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/di"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.Listen("tcp", c.PORT)
	if err != nil {
		log.Fatal("error", err)
	}

	svc := di.InitializeAPI()
	grpcServer := grpc.NewServer()
	job.RegisterJobServiceServer(grpcServer, &svc)
	fmt.Println(lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
