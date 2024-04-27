package domain

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Title       string
	Description string
	Category    int
	User_id     int
}

type TierProject struct {
	gorm.Model
	ProjectID   int
	Title       string
	Description string
	Price       float64
	DeliverDays       int
	NumberOfRevisions int
}

type SingleProject struct {
	gorm.Model
	ProjectID         int
	Price             float64
	DeliverDays       int
	NumberOfRevisions int
}

type ProjectOrders struct{
	gorm.Model
	ClientID int
	ProjectID int
	Payment_status string
	Delivery_status string
}
