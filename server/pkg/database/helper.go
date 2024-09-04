package database

import (
	"context"
	"jatin/pkg/constants"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindById[T any](ctx context.Context, db *mongo.Database, coll constants.CollectionNames, id string) (*T, error) {
	var obj T
	err := db.Collection(string(coll)).FindOne(ctx, bson.D{
		{
			Key:   "_id",
			Value: id,
		},
	}).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
