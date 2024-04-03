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
	Is_active bool   `gorm:"is_active" default:"true"`
}

type Freelancer_Description struct {
	gorm.Model
	User_id     int
	Title       string
	Description string
	Hourly_rate string
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

type Client struct {
	gorm.Model
	User_id         int
	Company_name    string
	Industry        string
	Phone           string
	Billing_address string
	Contact_person  string
}

type Skills struct {
	gorm.Model
	Name string
}

type Freelancer_skills struct {
	Freelancer_id int
	Skill_id      int
}

type OtpInfo struct {
	ID         uint `gorm:"primaryKey"`
	Email      string
	OTP        int
	Expiration time.Time
}
