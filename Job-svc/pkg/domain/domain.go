package domain

import (
	"time"

	"gorm.io/gorm"
)

type Jobs struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"Description"`
	TimePeriod  string  `json:"time-period"`
	Type        string  `json:"type"`
	Category    int64   `json:"category"`
	Budget      float32 `json:"budget"`
	Client_id   int
}

type Category struct{
	gorm.Model
	Category string
}


type JobSkills struct {
	ID       uint `gorm:"primarykey"`
	Job_id   int
	Skill_id int
}

type Proposals struct {
	gorm.Model
	Cover_letter string `json:"cover_letter"`
	Budget       float32
	JobId        int
	User_id      int
}

type Offer struct{
	gorm.Model
	Budget        float32
	Offer_letter  string
	Starting_time string
	Job_id        int
	Freelancer_id int
	Client_id int
	Status string
}

type Contract struct{
	gorm.Model
	Start_date time.Time
	Client_id int
	Freelancer_id int
	Type string
	Budget float32
	Job_id int
	Paid_amount int
	Pending_amount int
	Status string
	AttachmentUrl string 
}

type Invoice struct{
	gorm.Model
	Freelancer_fee float64
	MarketPlace_fee float64
	Start_date time.Time
	End_date   time.Time
	Status string
	ContractID int
}

type Attachment struct{
	gorm.Model
	ContractID int
	UploadTime time.Time
	UploadedBy int
	Description string
	AttachmentURL string
}