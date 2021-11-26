package controllers

import(
	"io/ioutil"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/toxtashev/LogisticApp/utils"
	model "github.com/toxtashev/LogisticApp/modules/neworder/models"
)

func Complete(ctx *gin.Context) {

	body, err := ioutil.ReadAll(ctx.Request.Body)

	defer utils.DB.Close()

	if err != nil {err.Error()}

	var complete model.OnlyId

	json.Unmarshal(body, &complete)

	utils.DB.Exec(model.COMPLETED, complete.Id)

	defer utils.DB.Close()

	ctx.Writer.WriteHeader(200)
}