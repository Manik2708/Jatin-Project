package factory

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"jatin/pkg/constants"
)

type Factory interface {
	CreateGinClient() (*gin.Engine, error)
	GetGinClient() *gin.Engine
	CreateDatabase() *mongo.Database
	GetDatabase() *mongo.Database
	CreateMongoClient() *mongo.Client
	GetMongoClient() *mongo.Client
	GetCollection(name constants.CollectionNames) *mongo.Collection
}