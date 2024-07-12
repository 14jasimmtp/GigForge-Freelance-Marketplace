package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/notification-svc/pb"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/notification-svc/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	Coll *mongo.Collection
}

func NewNotificationRepo(Coll *mongo.Collection) *Repo {
	return &Repo{Coll: Coll}
}

func (r *Repo) CreateNewNotification(body []byte) error {
	var notification models.NotificationModel
	err := json.Unmarshal(body, &notification)
	if err != nil {
		return err
	}
	notification.CreatedAt = time.Now()
	_, err = r.Coll.InsertOne(context.TODO(), &notification)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetNotificationsForUser(userId int32) ([]*pb.Notification, error) {
	filter := bson.M{"user_id": bson.M{"$in": bson.A{userId}}}

	cursor, err := r.Coll.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{"CreatedAt", -1}}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var notifications []*pb.Notification
	for cursor.Next(context.TODO()) {
		var notification models.NotificationModel
		if err := cursor.Decode(&notification); err != nil {
			return nil, err
		}
		n := &pb.Notification{
			Content: notification.CommentText,
			UserID:  strconv.Itoa(int(notification.UserID)),
			Type:    "notification",
		}
		fmt.Println(notification, "hi")

		notifications = append(notifications, n)
	}
	return notifications, nil
}
