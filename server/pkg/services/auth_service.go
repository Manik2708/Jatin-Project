package services

import (
	"jatin/pkg/constants"
	ct_errors "jatin/pkg/errors"
	"jatin/pkg/middleware"
	"jatin/pkg/schemas"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServiceTemplate interface {
	Login(indentify string, password string, userType constants.UserType) (*schemas.CustomerLoginOutput, error)
}

func (as *GlobalService) Login(identify string, password string, userType constants.UserType) (*schemas.CustomerLoginOutput, error) {
	var coll_name constants.CollectionNames
	if userType == constants.CUSTOMER_USER_TYPE {
		coll_name = constants.CUSTOMER_COLLECTION
	} else if userType == constants.ADMIN_USER_TYPE {
		coll_name = constants.ADMIN_COLLECTION
	} else {
		return nil, ct_errors.ErrInsertIdNotGenerated
	}
	isEmail := constants.EMAIL_REGEX.MatchString(identify)
	var find_result *mongo.SingleResult
	if isEmail {
		find_result = as.ft.GetCollection(coll_name).FindOne(as.ft.GetMongoContext(), bson.D{
			{
				Key:   "email",
				Value: identify,
			},
		})
	} else {
		find_result = as.ft.GetCollection(coll_name).FindOne(as.ft.GetMongoContext(), bson.D{
			{
				Key:   "user_name",
				Value: identify,
			},
		})
	}
	output := &schemas.CustomerLoginOutput{}
	payload := &middleware.AuthTokenPayload{}
	if coll_name == constants.CUSTOMER_COLLECTION {
		var customer *schemas.Customer
		err := find_result.Decode(&customer)
		if err != nil {
			return nil, err
		}
		output.Id = customer.Id
		output.Email = customer.Email
		output.UserName = customer.UserName
		output.UserType = customer.UserType
		output.Phone = customer.Phone
		output.Name = customer.Name
		payload.Type = string(customer.UserType)
	} else {
		var admin *schemas.Admin
		err := find_result.Decode(&admin)
		if err != nil {
			return nil, err
		}
		output := &schemas.CustomerLoginOutput{}
		output.Id = admin.Id
		output.Email = admin.Email
		output.UserName = admin.UserName
		output.UserType = admin.UserType
		output.Phone = admin.Phone
		output.Name = admin.Name
		payload.Type = string(admin.UserType)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token_output, err := token.SignedString("")
	if err != nil {
		return nil, err
	}
	output.Token = token_output
	return output, nil
}
