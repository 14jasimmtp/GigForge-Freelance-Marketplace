package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(c *config.Config) *gorm.DB {
	sql, err := sql.Open("postgres", c.DB_URL)
	if err != nil {
		log.Println("Error : ", err)
		return nil
	}

	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE datname = '" + "gigforge_job_svc_db" + "'")
	if err != nil {
		log.Println("something went wrong", err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("Database" + "gigforge_user_db" + " already exists.")
	} else {
		_, err = sql.Exec("CREATE DATABASE " + "gigforge_job_svc_db")
		if err != nil {
			fmt.Println("Error creating database:", err)
		}
	}

	db, err := gorm.Open(postgres.Open(c.DB_URL+"/gigforge_job_svc_db"), &gorm.Config{})
	if err != nil {
		log.Fatal("error while connecting to db : ", err)
	}

	// db.AutoMigrate(&domain.User{})
	// db.AutoMigrate(&domain.Client{})
	// db.AutoMigrate(&domain.Freelancer{})
	return db
}
