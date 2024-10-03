package factory

import (
	"context"
	"jatin/pkg/constants"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Factory interface {
	GetGinClient() *gin.Engine
	GetDatabase() *mongo.Database
	GetMongoClient() *mongo.Client
	GetCollection(name constants.CollectionNames) *mongo.Collection
	GetMongoContext() context.Context
}
