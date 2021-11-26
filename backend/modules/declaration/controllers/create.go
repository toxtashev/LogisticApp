package controllers

import(
	"io/ioutil"
	"encoding/json"
	model "github.com/toxtashev/LogisticApp/modules/declaration/models"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {

	body , err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {err.Error()}

	var Data model.NewOrder

	json.Unmarshal(body, &Data)

	model.Create(Data)

	ctx.Writer.WriteHeader(201)
}