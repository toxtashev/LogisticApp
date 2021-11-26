package controllers

import(
	"io/ioutil"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/toxtashev/LogisticApp/utils"
	model "github.com/toxtashev/LogisticApp/modules/neworder/models"
)

func Delete(ctx *gin.Context) {

	body, err := ioutil.ReadAll(ctx.Request.Body)

	defer utils.DB.Close()

	if err != nil {err.Error()}

	var delete model.OnlyId

	json.Unmarshal(body, &delete)

	utils.DB.Exec(model.DELETED, delete.Id)

	defer utils.DB.Close()

	ctx.Writer.WriteHeader(200)
}