package domain

import "gorm.io/gorm"

type Jobs struct{
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"Description"`
	TimePeriod  string  `json:"time-period"`
	Level       string  `json:"freelancer_level"`
	Category    int64   `json:"category"`
	Budget      string  `json:"budget"`
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
	Budget string
	JobId int
	User_id int
}