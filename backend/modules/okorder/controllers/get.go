package controllers

import(
	"github.com/gin-gonic/gin"

	model "github.com/toxtashev/LogisticApp/modules/okorder/models"
)

func Get (ctx *gin.Context) {

	ctx.JSON(200, model.Get())
}