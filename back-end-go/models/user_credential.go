package models

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
