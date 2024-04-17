package di

import (
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/db"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/repository"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/service"
)

func InitializeAPI() service.Service {
	db := db.ConnectToDB()
	repo := repository.NewJobRepo(db)
	service := service.NewJobService(repo)

	return service
}
