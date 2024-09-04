package services

import (
	"jatin/pkg/constants"
	"jatin/pkg/errors"
	"jatin/pkg/factory"
	"jatin/pkg/schemas"
	"unicode"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type CustomerServiceTemplate interface {
	CreateCustomer(ct *schemas.Customer) (*schemas.Customer, error)
}

type CustomerService struct {
	ctx *gin.Context
	ft  factory.Factory
}

func (cs *CustomerService) CreateCustomer(ct *schemas.Customer) (*schemas.Customer, error) {
	if ct.Password == nil {
		return nil, errors.ErrPasswordNotEntered
	}
	if !constants.EMAIL_REGEX.MatchString(ct.Email) {
		return nil, errors.ErrInvalidEmailAddress
	}
	err := isPasswordStrong(*ct.Password)
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

func isPasswordStrong(p string) error {
	errs := []error{
		errors.ErrPasswordIsShort,
		errors.ErrPasswordNotHaveUpperCase,
		errors.ErrPasswordNotHaveLowerCase,
		errors.ErrPasswordNotHaveNumber,
		errors.ErrPasswordNotHaveSpecialChar,
	}
	validation := []bool{
		false,
		false,
		false,
		false,
		false,
	}
	s := []rune(p)
	if len(s) >= 7 {
		validation[0] = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			validation[1] = true
		case unicode.IsLower(char):
			validation[2] = true
		case unicode.IsNumber(char):
			validation[3] = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			validation[4] = true
		}
	}
	for i, vdt := range validation {
		if !vdt {
			return errs[i]
		}
	}
	return nil
}
