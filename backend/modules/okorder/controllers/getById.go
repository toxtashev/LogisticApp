package controllers

import(
	"github.com/gin-gonic/gin"

	model "github.com/toxtashev/LogisticApp/modules/okorder/models"
)

func GetById(ctx *gin.Context) {

	order := model.GetById(ctx.Param("itemId"))

	ctx.JSON(200, order)
}