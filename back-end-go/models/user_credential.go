package models

import "time"

type UserCredentialDatabaseStruct struct {
	Username    string    `bson:"username"`
	Password    string    `bson:"password"`
	CreatedTime time.Time `bson:"createdTime"`
}

type UserCredentialParsedFromDB struct {
	Id          string    `bson:"_id"`
	Username    string    `bson:"username"`
	Password    string    `bson:"password"`
	CreatedTime time.Time `bson:"createdTime"`
}
