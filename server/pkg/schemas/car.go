package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	Id     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId string             `json:"user_id,omitempty" bson:"user_id"`
	Name   string             `bson:"name" json:"name"`
	Make   string             `json:"make" bson:"make"`   // Manufacturer of the car (e.g., Toyota, Ford)
	Model  string             `json:"model" bson:"model"` // Model of the car (e.g., Corolla, Mustang)
	Year   int                `json:"year" bson:"year"`   // Manufacturing year
}
