package dao

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	client *mongo.Client
	db     *mongo.Database
}

var mg *MongoInstance

const (
	mongoURI = "mongodb://localhost:27017"
	dbName   = "jobs"
)

func Connect() {
	// create db connection here
	opts := options.Client().ApplyURI(mongoURI)
	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, _ := mongo.Connect(context, opts)
	defer cancel()

	mg = &MongoInstance{
		client: client,
		db:     client.Database(dbName),
	}
}
