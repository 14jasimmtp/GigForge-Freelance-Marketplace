package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Messages struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Sender_id     int                `bson:"sender_id"`
	Reciepient_id int                `bson:"recepient_id"`
	Message       string             `bson:"message"`
	Timestamp     time.Time          `bson:"timestamp"`
	Read          bool               `bson:"read"`
}
