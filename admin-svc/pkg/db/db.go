package db

import (
	"fmt"
	"log"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/admin-svc/pkg/domain"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB{
	db, err := gorm.Open(postgres.Open(viper.GetString("DB_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	dbName := "gigforge_admin_db"

	var exists bool
	err = db.Raw("SELECT EXISTS (SELECT FROM pg_database WHERE datname = ?)",dbName).Scan(&exists).Error
	if err != nil {
		fmt.Println(err)
	}

	if !exists {
		err = db.Exec("CREATE DATABASE gigforge_admin_db").Error
		if err != nil {
			log.Fatal(err)
		}
		log.Println("created database " + dbName)
	}

	db, err = gorm.Open(postgres.Open(viper.GetString("DB_URL")+"/"+dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain.Admin{})
	log.Println("migrated table")
	createAdmin(db)
	return db
}

func createAdmin(db *gorm.DB){
	email:="admin@gigforge.com"
	password:="admin@123"

	hashed,_:=bcrypt.GenerateFromPassword([]byte(password),10)

	err:=db.FirstOrCreate(&domain.Admin{Email: email,Password: string(hashed)}).Error
	if err != nil {
		fmt.Print("failed to create admin")
	}
}

