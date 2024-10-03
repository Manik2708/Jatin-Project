package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type Appointment struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RequestedBy string             `bson:"requested_by" json:"requested_by"`
	AcceptedBy  string             `bson:"accepted_by,omitempty" json:"accepted_by,omitempty"`
	Address     Address            `bson:"address" json:"address"`
	CreatedAt   primitive.DateTime `bson:"created_at" json:"created_at,omitempty"`
	Car         Car                `bson:"car" json:"car"`
	UpdatedAt   primitive.DateTime `bson:"updated_at" json:"updated_at,omitempty"`
	Status      string             `bson:"status" json:"status"`
}
