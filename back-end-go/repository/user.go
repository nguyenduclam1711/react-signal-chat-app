package repository

import (
	"context"
	"time"

	"github.com/nguyenduclam1711/react-signal-chat-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserRepository CoreRepository[models.UserDatabaseStruct]

func NewUserRepository() {
	collectionName := "user"
	UserRepository = HandleCreateNewRepository[models.UserDatabaseStruct](CreateNewRepositoryParams[models.UserDatabaseStruct]{
		CollectionName: collectionName,
		InsertOne: func(coll *mongo.Collection) CoreRepositoryInsertOne[models.UserDatabaseStruct] {
			return func(payload models.UserDatabaseStruct) (*mongo.InsertOneResult, error) {
				payload.CreatedTime = time.Now()
				return coll.InsertOne(context.TODO(), payload)
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
