package di

import (
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/db"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/infrastructure/repository"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/infrastructure/service"
	"github.com/robfig/cron/v3"
)

func InitializeAPI() (*service.Service, *repository.Repo) {
	db := db.ConnectToDB()
	repo := repository.NewRepo(db)
	service := service.NewService(repo)
	cron := cron.New()
	_, err := cron.AddFunc("@every 2m", func() { repo.DeleteOTP() })
	if err != nil {
		panic(err)
	}

	cron.Start()
	return service, repo
}
