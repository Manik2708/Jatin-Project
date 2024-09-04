package constants

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var (
	EMAIL_REGEX = regexp.MustCompile(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`)
)

type CollectionNames string

const (
	CUSTOMER_COLLECTION CollectionNames = "customer"
	ADMIN_COLLECTION    CollectionNames = "admin"
)

type UserType string

const (
	ADMIN_USER_TYPE      UserType = "ADMIN"
	CUSTOMER_USER_TYPE   UserType = "CUSTOMER"
	SUPERADMIN_USER_TYPE UserType = "SUPERADMIN"
)

const (
	MIDDLEWARE_TOKEN   = "token"
	MIDDLEWARE_USER_ID = "id"
)

func SendOkResponse(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusOK, gin.H{
		"main_data": obj,
	})
}
