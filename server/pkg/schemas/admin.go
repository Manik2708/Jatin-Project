package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct{
	BaseUser
	AppointmentRequests primitive.ObjectID `bson:"appointment_requests,omitempty" json:"appointment_requests,omitempty"`
}

type AdminOrUser interface{
	*Customer | *Admin
}