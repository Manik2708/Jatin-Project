package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct{
	BaseUser
	Appointments []Appointment `bson:"appointments,omitempty" json:"appointments,omitempty"`
	Address Address `bson:"address,omitempty" json:"address,omitempty"`
	Cars []primitive.ObjectID `bson:"cars,omitempty" json:"cars,omitempty"`
}