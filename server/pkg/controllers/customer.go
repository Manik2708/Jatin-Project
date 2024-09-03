package controllers

import (
	"jatin/pkg/constants"
	ct_errors "jatin/pkg/errors"
	"jatin/pkg/schemas"
	"jatin/pkg/services"

	"github.com/gin-gonic/gin"
)

type UserController struct{
	svc services.CustomerServiceTemplate
}

func (uc *UserController) CreateCustomer(ctx *gin.Context){
	var customer *schemas.Customer
	err := ctx.ShouldBindJSON(&customer)
	if err != nil{
		ct_errors.ThrowBadRequestError(
			ctx,
			ct_errors.BINDING_ERROR,
			err,
		)
		return
	}
	customer, err = uc.svc.CreateCustomer(customer)
	if err !=nil{
		ct_errors.HandleServicesError(ctx, err)
		return
	}
	constants.SendOkResponse(ctx, *customer)
}