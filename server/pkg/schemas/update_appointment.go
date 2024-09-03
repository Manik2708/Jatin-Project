package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateAppointment struct{
		Id primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
		Appointment primitive.ObjectID `bson:"appointment" json:"appointment"`
		CreatedAt primitive.DateTime `bson:"created_at" json:"created_at,omitempty"`
}