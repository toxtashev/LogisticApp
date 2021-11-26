package controllers

import(
	"github.com/gin-gonic/gin"

	"github.com/toxtashev/LogisticApp/utils"
)

func Active(ctx *gin.Context) {

	var Total int

	err := utils.DB.QueryRow(
		`select
			count(order_id)
		from orders
		where
			isActivated = true AND isFinished = false
		`,
	).Scan(
		&Total,
	)

	if err != nil {err.Error()}

	ctx.JSON(200, Total)
}