package factory

import (
	"context"
	"jatin/pkg/constants"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	dbc *context.Context
	gn  *gin.Engine
	cl  *mongo.Client
	db  *mongo.Database
}

func (s *Server) CreateGinClient() (*gin.Engine, error) {
	s.gn = gin.New()
	err := s.gn.Run("0.0.0.0:8000")
	if err != nil {
		return nil, err
	}
	return s.gn, nil
}

func (s *Server) GetGinClient() *gin.Engine {
	return s.gn
}

func (s *Server) CreateDatabase() *mongo.Database {
	return s.cl.Database("jatin")
}

func (s *Server) GetDatabase() *mongo.Database {
	return s.db
}

func (s *Server) GetCollection(name constants.CollectionNames) *mongo.Collection {
	return s.db.Collection(string(name))
}

func (s *Server) CreateIndexes(ctx context.Context) error {
	cst := s.GetCollection(constants.CUSTOMER_COLLECTION)
	indexUserName := mongo.IndexModel{
		Keys: bson.D{
			{
				Key:   "user_name",
				Value: 1,
			},
		},
		Options: options.Index().SetUnique(true),
	}
	indexEmail := mongo.IndexModel{
		Keys: bson.D{
			{
				Key:   "email",
				Value: 1,
			},
		},
	}
	_, err := cst.Indexes().CreateMany(ctx, []mongo.IndexModel{
		indexEmail,
		indexUserName,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) GetMongoClient() *mongo.Client {
	return s.cl
}

func (s *Server) GetMongoContext() context.Context {
	return *s.dbc
}
