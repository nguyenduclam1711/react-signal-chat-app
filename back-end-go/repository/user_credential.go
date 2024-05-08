package repository

import (
	"context"
	"errors"
	"time"

	"github.com/nguyenduclam1711/react-signal-chat-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var UserCredentialRepository CoreRepository[models.RegisteredUserPayload]

func NewUserCredentialRepository() {
	collectionName := "user_credential"
	UserCredentialRepository = HandleCreateNewRepository[models.RegisteredUserPayload](CreateNewRepositoryParams[models.RegisteredUserPayload]{
		CollectionName: collectionName,
		InsertOne: func(coll *mongo.Collection) CoreRepositoryInsertOne[models.RegisteredUserPayload] {
			return func(payload models.RegisteredUserPayload) (*mongo.InsertOneResult, error) {
				if payload.Password == "" || payload.Username == "" {
					return nil, errors.New("Bad Request")
				}
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
				if err != nil {
					return nil, err
				}
				createPayload := models.UserCredentialDatabaseStruct{
					Username:    payload.Username,
					CreatedTime: time.Now(),
					Password:    string(hashedPassword),
				}
				return coll.InsertOne(context.TODO(), createPayload)
			}
		},
		CreateInitIndexes: func(coll *mongo.Collection) ([]string, error) {
			return coll.Indexes().CreateMany(
				context.TODO(),
				[]mongo.IndexModel{
					{
						Keys: bson.D{
							{
								Key:   "username",
								Value: 1,
							},
						},
						Options: options.Index().SetUnique(true),
					},
				},
			)
		},
	})
}
