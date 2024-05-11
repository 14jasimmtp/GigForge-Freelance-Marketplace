package domain

import (
	"time"

)

type Message struct {
	SenderID    string    `json:"SenderID" validate:"required"`
	RecipientID string    `json:"RecipientID" validate:"required"`
	Content     string    `json:"Content" validate:"required"`
	Timestamp   time.Time `json:"TimeStamp" validate:"required"`
	Type        string    `json:"Type" validate:"required"`
	Tag         string    `json:"Tag"`
	Status      string    `json:"Status"`
}

