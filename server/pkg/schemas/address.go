package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId        string             `json:"user_id,omitempty" bson:"user_id"`
	Name          string             `json:"name" bson:"name"`
	StreetAddress string             `json:"street_address" bson:"street_address"`
	City          string             `json:"city" bson:"city"`
	Province      string             `json:"province" bson:"province"`
	PostalCode    string             `json:"postal_code" bson:"postal_code"`
	Country       string             `json:"country" bson:"country"`
}
