package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname" gorm:"firstname"`
	LastName  string `json:"lastname" gorm:"lastname"`
	Email     string `gorm:"email"`
	Phone     string `gorm:"phone"`
	Password  string `gorm:"password"`
	Country   string `gorm:"country"`
	Role      string `json:"role" gorm:"role"`
	Is_active bool   `gorm:"is_active" default:"true"`
	Profile_URL string `gorm:"profile_url"`
}

type UserModel struct {
	ID        uint   `gorm:"id"`
	FirstName string `json:"firstname" gorm:"firstname"`
	LastName  string `json:"lastname" gorm:"lastname"`
	Email     string `gorm:"email"`
	Phone     string `gorm:"phone"`
	Password  string `gorm:"password"`
	Country   string `gorm:"country"`
	Role      string `json:"role" gorm:"role"`
	Is_active bool   `gorm:"is_active"`
}

type Freelancer_Description struct {
	gorm.Model
	User_id     int
	Title       string
	Description string
	Hourly_rate int
}

type Freelancer_Education struct {
	gorm.Model
	User_id       int    `gorm:"user_id"`
	School        string `gorm:"school"`
	Course        string `gorm:"course"`
	Year_Started  string `gorm:"year_started"`
	Year_Ended    string `gorm:"year_ended"`
	Area_Of_Study string `gorm:"area_of_study"`
	Description   string `gorm:"description"`
}

type Freelancer_Experiences struct{
	gorm.Model
	Company string
	City string
	Country string
	Title string
	FromDate string
	ToDate string
	Description string
	User_id int64
}

type ClientCompany struct{
	gorm.Model
	ClientID int
	CompanyName string
	Website string
	NumberOfEmployees int
	Tagline string
	Industry string
}

type CompanyAddress struct{
	gorm.Model
	ClientID int
	OwnerName string
	Phone string
	Country string
	State string 
	District string
	City string 
	Pincode string
}

type Skill struct {
	gorm.Model
	Skill       string
	Description string
}

type Freelancer_skills struct {
	ID            uint `gorm:"primarykey"`
	Freelancer_id int
	Skill_id      int
}

type OtpInfo struct {
	ID         uint `gorm:"primaryKey"`
	Email      string
	OTP        int
	Expiration time.Time
}

type Admin struct{
	gorm.Model
	Email string 
	Password string
}

type Freelancer_paypal struct{
	UserID uint `gorm:"primarykey"`
	Email string `gorm:"email"`
}

type FreelancerReview struct{
	gorm.Model
	Review        string
    Rating        int 
    Freelancer_id int    
	Client_id int
}