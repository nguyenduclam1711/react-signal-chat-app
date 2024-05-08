package database

import (
	"context"
	"log"

	"github.com/nguyenduclam1711/react-signal-chat-app/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MissingCollectionsReport struct {
	CollectionName string
	Error          error
}

var (
	MongoDatabase        *mongo.Database
	MongoCollectionNames = make([]string, 0)
	ExistCollectionNames = make(map[string]bool)
	MissingCollections   = make([]MissingCollectionsReport, 0)
)

func ConnectToDatabase() {
	uri := env.EnvData["MONGODB_URI"]
	cl, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	MongoDatabase = cl.Database(env.EnvData["MONGODB_DATABASE"])
	collections, collectionsErr := MongoDatabase.ListCollectionNames(
		context.TODO(),
		bson.D{},
	)
	if collectionsErr == nil {
		MongoCollectionNames = collections
		for _, name := range MongoCollectionNames {
			ExistCollectionNames[name] = true
		}
	}
}
