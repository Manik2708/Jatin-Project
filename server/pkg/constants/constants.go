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
	APPOINTMENTS_COLLECTION CollectionNames = "appointments"
	CAR_COLLECTION CollectionNames = "car"
	ADDRESS_COLLECTION CollectionNames = "address"
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

type AppointmentStatus string

const (
	PENDING_APPOINTMENT_STATUS = "PENDING"
	ACCEPTED_APPOINTMENT_STATUS = "ACCEPTED"
	REJECTED_APPOINTMENT_STATUS = "REJECTED"
	CLOSED_APPOINTMENT_STATUS = "CLOSED"
)

func SendOkResponse(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusOK, gin.H{
		"main_data": obj,
	})
}
