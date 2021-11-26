package aadmin

import(
	"os"
	"log"
	"io/ioutil"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ForAdmin struct {
	Name, Password string
}

func Login(ctx *gin.Context) {

	err := godotenv.Load()
	
	if err != nil {
		log.Print(err)
	}

	NAME		:= os.Getenv("ADMIN_NAME")
	PASSWORD	:= os.Getenv("ADMIN_PASSWORD")

	body, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {panic(err)}

	var data ForAdmin

	json.Unmarshal(body, &data)

	if data.Name == NAME && data.Password == PASSWORD {

		ctx.Writer.WriteHeader(200)

	} else {

		ctx.Writer.WriteHeader(404)
	}
}