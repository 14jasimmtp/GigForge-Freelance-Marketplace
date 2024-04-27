package di

import (
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pkg/db"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pkg/repository"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pkg/service"
)

func InitializeAPI() service.ProjectService {
	db := db.ConnectToDB()
	repo := repository.NewRepo(db)
	service := service.NewService(repo)

	return service
}
