package repository

import (
	"github.com/nguyenduclam1711/react-signal-chat-app/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryStruct struct {
	Coll *mongo.Collection
}

var UserRepository UserRepositoryStruct

func NewUserRepository() {
	coll := database.MongoDatabase.Collection("user")
	UserRepository = UserRepositoryStruct{
		Coll: coll,
	}
}

func (this UserRepositoryStruct) Create() {
}
