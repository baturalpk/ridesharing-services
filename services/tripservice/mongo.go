package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db *mongo.Database

func InitMongoConnection(uri, dbname string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if err = client.Connect(ctx); err != nil {
		return err
	}

	db = client.Database(dbname)
	return nil
}

func GetCollection(name string) *mongo.Collection {
	return db.Collection(name)
}

func CloseMongoConnection() {
	_ = client.Disconnect(context.Background())
}
