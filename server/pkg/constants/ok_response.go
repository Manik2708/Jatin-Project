package constants

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendOkResponse(ctx *gin.Context, obj any){
	ctx.JSON(http.StatusOK, gin.H{
		"main_data":obj,
	})
}