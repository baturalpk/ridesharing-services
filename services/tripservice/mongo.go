package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgc *mongo.Client
var mdb *mongo.Database

func InitMongoConnection(uri, dbname string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	if err = client.Connect(ctx); err != nil {
		cancel()
		return err
	}
	mdb = client.Database(dbname)
	cancel()
	return nil
}

func GetCollection(name string) *mongo.Collection {
	if mdb == nil {
		panic("need to initialize mongo connection beforehand")
	}
	return mdb.Collection(name)
}

func CloseMongoConnection() {
	if mgc != nil {
		_ = mgc.Disconnect(context.Background())
	}
}
