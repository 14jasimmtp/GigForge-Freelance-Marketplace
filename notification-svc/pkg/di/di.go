package di

import (
	"log"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/notification-svc/pkg/config"
	broker "github.com/14jasimmtp/GigForge-Freelance-Marketplace/notification-svc/pkg/consumer"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/notification-svc/pkg/db"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/notification-svc/pkg/repository"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/notification-svc/pkg/service"
)

func InjectDependencies() *service.Service{
	err:=config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	Coll,err:=db.ConnectMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	conn:=broker.ConnectAMQP()
	repo:=repository.NewNotificationRepo(Coll)
	svc:=service.NewNotificationService(repo,conn)
	return svc
}