package errors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type BadRequestErrorMesssage string
type InternalServerErrorMessage string
type ForbiddenErrorMessage string

var (
	ErrInsertIdNotGenerated       = errors.New("insert id is not generated from database. please try again")
	ErrPasswordNotEntered         = errors.New("password field is nil")
	ErrPasswordIsShort            = errors.New("password is too-short, it's minimum length should be 8")
	ErrPasswordNotHaveUpperCase   = errors.New("password doesn't have any uppercase letter, it should contain atleast one")
	ErrPasswordNotHaveLowerCase   = errors.New("password doesn't have any lowercase letter, it should contain atleast one")
	ErrPasswordNotHaveNumber      = errors.New("password doesn't have any number, it should contain atleast one")
	ErrPasswordNotHaveSpecialChar = errors.New("password doesn't have any special character, it should contain atleast one")
	ErrInvalidEmailAddress        = errors.New("invalid email-address, please provide valid email")
	ErrInvalidUserType            = errors.New("invalid user type, it should be either CUSTOMER or ADMIN")
	ErrUserNotAllowedToChange     = errors.New("the user id is not the owner of the entity, hence can't update or delete the entity")
)

const (
	INSERT_ID_NOT_INSERTED        InternalServerErrorMessage = "Server is not able to retrieve the id from database, please try again later!"
	DEFAULT_INTERNAL_SERVER_ERROR InternalServerErrorMessage = "Internal Server Error"
)

const (
	BINDING_ERROR                BadRequestErrorMesssage = "Unable to bind the request body with customer"
	PASSWORD_TOO_LONG_ERROR      BadRequestErrorMesssage = "Password is too long, please consider shortening it"
	PASSWORD_IS_WEEK             BadRequestErrorMesssage = "Weak Password"
	INVALID_EMAIL                BadRequestErrorMesssage = "Invalid Email-Address!!"
	INVALID_TOKEN                BadRequestErrorMesssage = "There is some problem with token, please see the error!"
	USER_RELATED_TOKEN_NOT_FOUND BadRequestErrorMesssage = "No user found related to this token found."
	INVALID_USER_TYPE            BadRequestErrorMesssage = "Invalid user type"
)

func ThrowBadRequestError(ctx *gin.Context, message BadRequestErrorMesssage, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": message,
		"error":   err.Error(),
	})
}

func ThrowInternalServerError(ctx *gin.Context, message InternalServerErrorMessage, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": message,
		"error":   err.Error(),
	})
}

func HandleServicesError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, bcrypt.ErrPasswordTooLong):
		ThrowBadRequestError(
			ctx,
			PASSWORD_TOO_LONG_ERROR,
			err,
		)
	case errors.Is(err, ErrPasswordIsShort) ||
		errors.Is(err, ErrPasswordNotHaveLowerCase) ||
		errors.Is(err, ErrPasswordNotHaveNumber) ||
		errors.Is(err, ErrPasswordNotHaveSpecialChar) ||
		errors.Is(err, ErrPasswordNotHaveUpperCase):
		ThrowBadRequestError(
			ctx,
			PASSWORD_IS_WEEK,
			err,
		)
	case errors.Is(err, ErrInvalidEmailAddress):
		ThrowBadRequestError(
			ctx,
			INVALID_EMAIL,
			err,
		)
	case errors.Is(err, ErrInvalidUserType):
		ThrowBadRequestError(
			ctx,
			INVALID_USER_TYPE,
			err,
		)
		// --------------------------------------------------------------------------------------------------------------------------------------------------------------------
		// --------------------------------------------------------------------------------------------------------------------------------------------------------------------
	case errors.Is(err, ErrInsertIdNotGenerated):
		ThrowInternalServerError(
			ctx,
			INSERT_ID_NOT_INSERTED,
			err,
		)
	default:
		ThrowInternalServerError(
			ctx,
			DEFAULT_INTERNAL_SERVER_ERROR,
			err,
		)
	}
}
