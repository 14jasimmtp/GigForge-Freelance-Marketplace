package db

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() (*mongo.Collection, error) {
	fmt.Println(viper.GetString("mongoURL"))
	clientOptions := options.Client().ApplyURI(viper.GetString("mongoURL"))

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
	

	coll := client.Database("Gigforge_chat_svc").Collection("Messages_user")
	fmt.Println("Connected to MongoDB!")

	return coll, nil
}
