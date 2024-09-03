package middleware

import (
	"jatin/pkg/constants"
	"jatin/pkg/factory"
	"jatin/pkg/schemas"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

type AuthService struct {
	s factory.Factory
}

func (auth *AuthService) CustomerAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		decodedDetails, err := jwt.ParseWithClaims(token, &AuthTokenPayload{}, func(token *jwt.Token) (interface{}, error) {
			return "", nil
		})
		parsedUserDetails := decodedDetails.Claims.(*AuthTokenPayload)
		if err != nil {

		}
		var user *schemas.Customer
		err = auth.s.GetCollection(constants.CUSTOMER_COLLECTION).FindOne(ctx, bson.D{
			{
				Key:   "_id",
				Value: parsedUserDetails.ID,
			},
		}).Decode(&user)
		if err != nil{
			
		}
		
	}
}

type AuthTokenPayload struct {
	Id   string `json:"id"`
	Type string `json:"Type"`
	jwt.RegisteredClaims
}
