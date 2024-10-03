package database

import (
	"context"
	"jatin/pkg/constants"
	"jatin/pkg/factory"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseHelper[T any] struct {
	ctx context.Context
	db *mongo.Database
	coll_name constants.CollectionNames
}

func (dh DatabaseHelper[T]) New(ft factory.Factory, coll_name constants.CollectionNames) *DatabaseHelper[T] {
 return &DatabaseHelper[T]{
	ctx: ft.GetMongoContext(),
	db: ft.GetDatabase(),
	coll_name: coll_name,
 }
}

func (dh *DatabaseHelper[T])FindById(id string) (*T, error) {
	var obj T
	err := dh.db.Collection(string(dh.coll_name)).FindOne(dh.ctx, bson.D{
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

func (dh *DatabaseHelper[T]) FindByIdAndUpdateOne(id string, updateInterface T) (*T, error){
	var obj T
	err := dh.db.Collection(string(dh.coll_name)).FindOneAndUpdate(dh.ctx, bson.D{
		{
			Key: "_id",
			Value: id,
		},
	},
	updateInterface,
	).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func (dh *DatabaseHelper[T]) FindByIdAndDelete(id string) (*T, error){
	var obj T
	err := dh.db.Collection(string(dh.coll_name)).FindOneAndDelete(dh.ctx, bson.D{
		{
			Key: "_id",
			Value: id,
		},
	}).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func (dh *DatabaseHelper[T]) GetPaginatedFindOptions(limit int64, page int64) *options.FindOptions {
	options := &options.FindOptions{}
	options.SetLimit(limit)
	options.SetSkip((page-1)*limit)
	return options
}
