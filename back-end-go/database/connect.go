package database

import (
	"context"
	"log"

	"github.com/nguyenduclam1711/react-signal-chat-app/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDatabase *mongo.Database

func ConnectToDatabase() {
	uri := env.EnvData["MONGODB_URI"]
	cl, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	MongoDatabase = cl.Database(env.EnvData["MONGODB_DATABASE"])
}
