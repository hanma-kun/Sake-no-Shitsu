package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() *mongo.Client {

	uri := "mongodb+srv://development:sake@sake.mqjgftx.mongodb.net/?retryWrites=true&w=majority" // change it in future (use ENV)
	fmt.Println(uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Connected to the mongodb")
	return client
}

var ClientConnection *mongo.Client = MongoConnection()

func CreateContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 100*time.Second)
}
