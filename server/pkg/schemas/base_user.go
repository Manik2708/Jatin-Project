package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	constants "jatin/pkg/constants"
)

type BaseUser struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserName string             `bson:"user_name" json:"user_name,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Phone    string             `bson:"phone" json:"phone"`
	UserType constants.UserType `bson:"user_type" json:"user_type,omitempty"`
	Password *string            `bson:"password" json:"password,omitempty"`
}
