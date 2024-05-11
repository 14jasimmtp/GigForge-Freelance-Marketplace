package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pb"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/chat-svc/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	Coll *mongo.Collection
}

func NewRepository(coll *mongo.Collection) *Repo {
	return &Repo{Coll: coll}
}

func (r *Repo) SaveMessage(msg []byte) error {
	var message domain.Message
	err := json.Unmarshal(msg, &message)
	if err != nil {
		return err
	}
	message.Timestamp = time.Now()
	_, err = r.Coll.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetChats(req *pb.GetChatReq) ([]*pb.Message, error) {

	var messages []*pb.Message

	filter := bson.M{"senderid": bson.M{"$in": bson.A{req.SenderId, req.RecieverId}}, "recipientid": bson.M{"$in": bson.A{req.RecieverId, req.SenderId}}}

	cursor, err := r.Coll.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{"timestamp", -1}}))
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var message domain.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		m:=&pb.Message{
			SenderId: message.SenderID,
			RecipientId: message.RecipientID,
			Content: message.Content,
			Timestamp: message.Timestamp.String(),
			Type: message.Type,
			Status: message.Status,
		}

		messages = append(messages, m)
	}
	return messages, nil
}
