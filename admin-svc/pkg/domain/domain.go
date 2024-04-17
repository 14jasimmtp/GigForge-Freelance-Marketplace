package domain

import "gorm.io/gorm"

type Admin struct{
	gorm.Model
	Email string 
	Password string
}