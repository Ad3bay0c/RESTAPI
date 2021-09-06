package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var Client *mongo.Client

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)

	Client, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	err := Client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}else {
		log.Println("Database Connected...")
	}
}
