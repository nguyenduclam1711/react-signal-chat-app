package models

import "time"

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
