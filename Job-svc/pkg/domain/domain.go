package domain

import "gorm.io/gorm"

type Jobs struct{
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"Description"`
	TimePeriod  string  `json:"time-period"`
	Type       	string  `json:"type"`
	Category    int64   `json:"category"`
	Budget      float32  `json:"budget"`
	Client_id int
}

type JobSkills struct{
	ID uint `gorm:"primarykey"`
	Job_id int
	Skill_id int
}

type Proposals struct{
	gorm.Model
	Cover_letter string `json:"cover_letter"`
	Budget float32
	JobId int
	User_id int
}