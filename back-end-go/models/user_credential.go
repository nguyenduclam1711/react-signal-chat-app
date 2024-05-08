package models

import "time"

type UserCredentialFrontendPayload struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func NewUserCredentialFrontendPayload() UserCredentialFrontendPayload {
	return UserCredentialFrontendPayload{}
}

type CreatedUserCredentialResponse struct {
	Username string `json:"username" bson:"username"`
	Id       string `json:"id" bson:"_id"`
}

func NewCreatedUserCredentialResponse() CreatedUserCredentialResponse {
	return CreatedUserCredentialResponse{}
}

type UserCredentialDatabaseStruct struct {
	Username    string    `bson:"username"`
	Password    string    `bson:"password"`
	CreatedTime time.Time `bson:"createdTime"`
}
