package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateAppointment struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Appointment string             `bson:"appointment" json:"appointment"`
	Message     string             `bson:"message" json:"message"`
	CreatedAt   primitive.DateTime `bson:"created_at" json:"created_at,omitempty"`
}
