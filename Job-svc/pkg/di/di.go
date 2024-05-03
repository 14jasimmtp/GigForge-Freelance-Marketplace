package di

import (
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/client"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/db"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/repository"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/service"
)

func InitializeAPI() service.Service {
	jobClient:=client.InitJobClient()
	db := db.ConnectToDB()
	repo := repository.NewJobRepo(db,jobClient)
	service := service.NewJobService(repo)

	return service
}
