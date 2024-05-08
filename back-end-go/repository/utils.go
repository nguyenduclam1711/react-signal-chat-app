package repository

import (
	"errors"
	"strings"

	"github.com/nguyenduclam1711/react-signal-chat-app/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type CoreRepositoryInsertOne[M interface{}] func(payload M) (*mongo.InsertOneResult, error)

type CoreRepository[InsertOnePayload interface{}] struct {
	Coll           *mongo.Collection
	CollectionName string
	InsertOne      CoreRepositoryInsertOne[InsertOnePayload]
}

type CreateNewRepositoryParams[InsertOnePayload interface{}] struct {
	CollectionName    string
	CreateInitIndexes func(coll *mongo.Collection) ([]string, error)
	InsertOne         func(coll *mongo.Collection) CoreRepositoryInsertOne[InsertOnePayload]
}

func createNewRepository[InsertOnePayload interface{}](params CreateNewRepositoryParams[InsertOnePayload]) (CoreRepository[InsertOnePayload], error) {
	collectionName := params.CollectionName
	result := CoreRepository[InsertOnePayload]{}
	if collectionName == "" {
		return result, errors.New("Collection name can't be empty")
	}
	result.CollectionName = collectionName
	if params.InsertOne == nil {
		return result, errors.New("Insert One function can't be empty")
	}
	result.Coll = database.MongoDatabase.Collection(collectionName)
	result.InsertOne = params.InsertOne(result.Coll)

	// the collection doesnt exist so we're gonna create one with index
	if params.CreateInitIndexes != nil && !database.ExistCollectionNames[collectionName] {
		_, createIndexesErr := params.CreateInitIndexes(result.Coll)
		if createIndexesErr != nil {
			return result, createIndexesErr
		}
	}
	return result, nil
}

func HandleCreateNewRepository[InsertOnePayload interface{}](params CreateNewRepositoryParams[InsertOnePayload]) CoreRepository[InsertOnePayload] {
	collectionName := strings.TrimSpace(params.CollectionName)
	params.CollectionName = collectionName
	res, err := createNewRepository(params)
	if err != nil {
		database.MissingCollections = append(database.MissingCollections, database.MissingCollectionsReport{
			CollectionName: params.CollectionName,
			Error:          err,
		})
	}
	return res
}
