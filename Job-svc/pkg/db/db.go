package db

import (
	"fmt"
	"log"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/domain"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(viper.GetString("DB_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	dbName := "gigforge_jobs_db"

	var exists bool
	err = db.Raw("SELECT EXISTS (SELECT FROM pg_database WHERE datname = ?)", dbName).Scan(&exists).Error
	if err != nil {
		fmt.Println(err)
	}

	if !exists {
		err = db.Exec("CREATE DATABASE gigforge_jobs_db").Error
		if err != nil {
			log.Fatal(err)
		}
		log.Println("created database " + dbName)
	}

	db, err = gorm.Open(postgres.Open(viper.GetString("DB_URL")+"/"+dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain.Jobs{})
	db.AutoMigrate(&domain.Proposals{})
	db.AutoMigrate(&domain.JobSkills{})

	return db
}
