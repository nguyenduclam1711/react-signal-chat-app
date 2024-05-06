package repository

import (
	"context"

	"github.com/nguyenduclam1711/react-signal-chat-app/database"
	"github.com/nguyenduclam1711/react-signal-chat-app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCredentialRepositoryStruct struct {
	Coll *mongo.Collection
}

var UserCredentialRepository UserCredentialRepositoryStruct

func NewUserCredentialRepository() {
	coll := database.MongoDatabase.Collection("user_credential")
	UserCredentialRepository = UserCredentialRepositoryStruct{
		Coll: coll,
	}
}

func (this UserCredentialRepositoryStruct) Create(payload models.UserCredentialFrontendPayload) (*mongo.InsertOneResult, error) {
	return this.Coll.InsertOne(context.TODO(), payload)
}
