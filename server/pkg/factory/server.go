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

func (s *Server) New() (*Server, error) {
	svr := &Server{}
	svr.db = svr.cl.Database("jatin")
	svr.gn = gin.New()
	err := svr.gn.Run("0.0.0.0:8000")
	if err != nil {
		return nil, err
	}
	return svr, nil
}

func (s *Server) GetGinClient() *gin.Engine {
	return s.gn
}

func (s *Server) GetDatabase() *mongo.Database {
	return s.db
}

func (s *Server) GetCollection(name constants.CollectionNames) *mongo.Collection {
	return s.db.Collection(string(name))
}

func (s *Server) CreateIndexes(ctx context.Context) error {
	cst := s.GetCollection(constants.CUSTOMER_COLLECTION)
	apt := s.GetCollection(constants.APPOINTMENTS_COLLECTION)
	car := s.GetCollection(constants.CAR_COLLECTION)
	adr := s.GetCollection(constants.ADDRESS_COLLECTION)
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
	indexAppointmentStatus := mongo.IndexModel{
		Keys: bson.D{
			{
				Key:   "status",
				Value: 1,
			},
		},
		Options: options.Index().SetUnique(false),
	}
	indexAppointmentRequestedBy := mongo.IndexModel{
		Keys: bson.D{
			{
				Key:   "requested_by",
				Value: 1,
			},
		},
		Options: options.Index().SetUnique(true),
	}
	indexAppointmentAcceptedBy := mongo.IndexModel{
		Keys: bson.D{
			{
				Key:   "accepted_by",
				Value: 1,
			},
		},
		Options: options.Index().SetUnique(true),
	}
	indexCarAddressUserId := mongo.IndexModel{
		Keys: bson.D{
			{
				Key:   "user_id",
				Value: 1,
			},
		},
		Options: options.Index().SetUnique(false),
	}
	_, err := cst.Indexes().CreateMany(ctx, []mongo.IndexModel{
		indexEmail,
		indexUserName,
	})
	if err != nil {
		return err
	}
	_, err = apt.Indexes().CreateMany(ctx, []mongo.IndexModel{
		indexAppointmentStatus,
		indexAppointmentRequestedBy,
		indexAppointmentAcceptedBy,
	})
	if err != nil {
		return err
	}
	_, err = car.Indexes().CreateMany(ctx, []mongo.IndexModel{
		indexCarAddressUserId,
	})
	if err != nil {
		return err
	}
	_, err = adr.Indexes().CreateMany(ctx, []mongo.IndexModel{
		indexCarAddressUserId,
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
