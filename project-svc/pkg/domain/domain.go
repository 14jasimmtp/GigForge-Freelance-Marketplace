package domain

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Title       string
	Description string
	CategoryID  int
}

type Tiers struct {
	PID          int
	Title        string
	Description  string
	DeliveryDays int
	Price        float64
}

