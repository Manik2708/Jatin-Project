package middleware

import (
	"jatin/pkg/constants"
	"jatin/pkg/database"
	"jatin/pkg/errors"
	"jatin/pkg/factory"
	"jatin/pkg/schemas"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService struct {
	s factory.Factory
}

func (auth *AuthService) CustomerAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get(constants.MIDDLEWARE_TOKEN)
		decodedDetails, err := jwt.ParseWithClaims(token, &AuthTokenPayload{}, func(token *jwt.Token) (interface{}, error) {
			return "", nil
		})
		if err != nil {
			errors.ThrowBadRequestError(ctx, errors.INVALID_TOKEN, err)
			return
		}
		parsedUserDetails := decodedDetails.Claims.(*AuthTokenPayload)
		customer, err := database.FindById[schemas.Customer](
			auth.s.GetMongoContext(),
			auth.s.GetDatabase(),
			constants.CUSTOMER_COLLECTION,
			parsedUserDetails.Id,
		)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				errors.ThrowBadRequestError(ctx, errors.USER_RELATED_TOKEN_NOT_FOUND, err)
			} else {
				errors.ThrowInternalServerError(ctx, errors.DEFAULT_INTERNAL_SERVER_ERROR, err)
			}
			return
		}
		ctx.Set(constants.MIDDLEWARE_USER_ID, customer.Id.Hex())
		ctx.Next()
	}
}

type AuthTokenPayload struct {
	Id   string `json:"Id"`
	Type string `json:"Type"`
	jwt.RegisteredClaims
}
