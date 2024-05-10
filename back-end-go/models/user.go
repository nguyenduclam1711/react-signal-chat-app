package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisteredUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisteredUserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type UserDatabaseStruct struct {
	Username    string    `bson:"username"`
	CreatedTime time.Time `bson:"createdTime"`
}

type UserParseFromDB struct {
	Id          string    `bson:"_id" json:"id"`
	Username    string    `bson:"username" json:"username"`
	CreatedTime time.Time `bson:"createdTime" json:"createdTime"`
}

type CurrentUser struct {
	Id       string
	Username string
	MongoId  primitive.ObjectID
}
