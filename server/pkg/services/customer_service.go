package services

import (
	"jatin/pkg/constants"
	"jatin/pkg/errors"
	"jatin/pkg/schemas"
	"jatin/pkg/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type CustomerServiceTemplate interface {
	CreateCustomer(userName string, name string, email string, phone string, password string) (*schemas.Customer, error)
}

func (cs *GlobalService) CreateCustomer(userName string, name string, email string, phone string, password string) (*schemas.Customer, error) {
	ct := &schemas.Customer{}
	ct.Password = &password
	ct.Phone = phone
	ct.Email = email
	ct.Name = name
	ct.UserName = userName
	ct.UserType = constants.CUSTOMER_USER_TYPE
	if ct.Password == nil {
		return nil, errors.ErrPasswordNotEntered
	}
	if !constants.EMAIL_REGEX.MatchString(ct.Email) {
		return nil, errors.ErrInvalidEmailAddress
	}
	err := utils.IsPasswordStrong(*ct.Password)
	if err != nil {
		return nil, err
	}
	col := cs.ft.GetCollection(constants.CUSTOMER_COLLECTION)
	hp, err := bcrypt.GenerateFromPassword([]byte(*ct.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	str := string(hp)
	ct.Password = &str
	res, err := col.InsertOne(cs.ft.GetMongoContext(), ct)
	if err != nil {
		return nil, err
	}
	if obid, ok := res.InsertedID.(primitive.ObjectID); ok {
		ct.Id = obid
	} else {
		return nil, errors.ErrInsertIdNotGenerated
	}
	ct.Password = nil
	return ct, nil
}
