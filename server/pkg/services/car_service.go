package services

import (
	"jatin/pkg/constants"
	"jatin/pkg/errors"
	"jatin/pkg/schemas"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CarServiceTemplate interface {
	CreateCar(userId string, name string, make string, model string, year int) (*schemas.Car, error)
	UpdateCar(updatedCar *schemas.Car) (*schemas.Car, error)
	DeleteCar(id string) (*schemas.Car, error)
	GetCarsByUserId(userId string) (*[]schemas.Car, error)
	GetCarByCarId(id string, limit int, page int) (*schemas.Car, error)
}

func (cs *GlobalService) CreateCar(userId string, name string, make string, model string, year int) (*schemas.Car, error) {
	car := &schemas.Car{}
	car.Make = make
	car.Model = model
	car.Name = name
	car.Year = year
	car.UserId = userId
	coll := cs.ft.GetCollection(constants.CAR_COLLECTION)
	res, err := coll.InsertOne(cs.ft.GetMongoContext(), car)
	if err != nil {
		return nil, err
	}
	if obid, ok := res.InsertedID.(primitive.ObjectID); ok {
		car.Id = obid
	} else {
		return nil, errors.ErrInsertIdNotGenerated
	}
	return car, nil
}

func (cs *GlobalService) UpdateCar(updatedCar *schemas.Car, userId string) (*schemas.Car, error) {
	car, err := cs.cardh.FindById(updatedCar.Id.String())
	if err != nil {
		return nil, err
	}
	if strings.Compare(car.UserId, updatedCar.UserId) != 0 || strings.Compare(car.UserId, userId) != 0 {
		return nil, errors.ErrUserNotAllowedToChange
	}
	return cs.cardh.FindByIdAndUpdateOne(updatedCar.Id.String(), *updatedCar)
}

func (cs *GlobalService) DeleteCar(id string, userId string) (*schemas.Car, error) {
	car, err := cs.cardh.FindById(id)
	if err != nil {
		return nil, err
	}
	if strings.Compare(car.UserId, userId) != 0 {
		return nil, errors.ErrUserNotAllowedToChange
	}
	return cs.cardh.FindByIdAndDelete(id)
}

func (cs *GlobalService) GetCarByCarId(id string) (*schemas.Car, error){
	return cs.cardh.FindById(id)
}

func (cs *GlobalService) GetCarsByUserId(id string, limit int64, page int64) (*[]schemas.Car, error){
	options := cs.cardh.GetPaginatedFindOptions(limit,page)
	var cars *[]schemas.Car
	res, err := cs.ft.GetCollection(constants.CAR_COLLECTION).Find(cs.ft.GetMongoContext(), bson.D{
		{
			Key: "user_id",
			Value: id,
		},
	},
	options,
)
	if err != nil {
	return nil, err
	}
	err = res.All(cs.ft.GetMongoContext(), cars)
	if err != nil {
		return nil, err
	}
	return cars, nil	
}
