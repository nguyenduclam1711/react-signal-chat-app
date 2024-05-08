package models

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertObjectIdToString(id interface{}) string {
	if oid, ok := id.(primitive.ObjectID); ok {
		return oid.Hex()
	}
	return "Cannot parse object id to string"
}
