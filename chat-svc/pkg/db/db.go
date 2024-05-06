package db

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() (*mongo.Collection, error) {

	clientOptions := options.Client().ApplyURI(viper.GetString("Mongo_Url"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	// message := domain.Messages{
	// 	Sender_id:     1,
	// 	Reciepient_id: 2,
	// 	Message:       "Hello, how are you?",
	// 	Timestamp:     time.Now(),
	// 	Read:          false,
	// }

	coll := client.Database("Gigforge_chat_svc").Collection("Messages_user")
	// coll.InsertOne(context.TODO(),&message)
	fmt.Println("Connected to MongoDB!")

	return coll, nil
}
