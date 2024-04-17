package di

import (
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/admin-svc/pkg/db"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/admin-svc/pkg/repository"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/admin-svc/pkg/service"
)

func InitializeAPI() service.AdminService{
	db := db.ConnectToDB()
	repo := repository.NewAdminRepo(db)
	service := service.NewAdminService(repo)

	return service
}
