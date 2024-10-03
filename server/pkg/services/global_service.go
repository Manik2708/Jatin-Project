package services

import (
	"jatin/pkg/constants"
	"jatin/pkg/database"
	"jatin/pkg/factory"
	"jatin/pkg/schemas"
)

type GlobalService struct {
	ft    factory.Factory
	cardh *database.DatabaseHelper[schemas.Car]
}

func (gs *GlobalService) New(server *factory.Server) *GlobalService {
	gblsvc := &GlobalService{}
	gblsvc.ft = server
	gblsvc.cardh = database.DatabaseHelper[schemas.Car]{}.New(gblsvc.ft, constants.CAR_COLLECTION)
	return gblsvc
}
