package db

import (
	"fmt"
	"log"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/domain"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(viper.GetString("DB_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	dbName := "gigforge_user_db"

	var exists bool
	err = db.Raw("SELECT EXISTS (SELECT FROM pg_database WHERE datname = ?)", dbName).Scan(&exists).Error
	if err != nil {
		fmt.Println(err)
	}

	if !exists {
		err = db.Exec("CREATE DATABASE gigforge_user_db").Error
		if err != nil {
			log.Fatal(err)
		}
		log.Println("created database " + dbName)
	}

	db, err = gorm.Open(postgres.Open(viper.GetString("DB_URL")+"/"+dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Client{})
	db.AutoMigrate(&domain.Freelancer_Description{})
	db.AutoMigrate(&domain.OtpInfo{})
	db.AutoMigrate(&domain.Freelancer_Education{})
	db.AutoMigrate(&domain.Freelancer_skills{})
	log.Println("migrated table")
	return db
}
