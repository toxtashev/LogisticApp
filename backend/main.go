package main

import (
	"os"
	"log"

	// "github.com/toxtashev/LogisticApp/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main () {

	gin.SetMode(gin.ReleaseMode)
	
	err := godotenv.Load()
	
	if err != nil {
		log.Print(err)
	}

	HOST := os.Getenv("APP_PG_HOST")
	USER := os.Getenv("APP_PG_USER")
	PASSWORD := os.Getenv("APP_PG_PASSWORD")
	PORT := os.Getenv("APP_PG_PORT")
	DBNAME := os.Getenv("APP_PG_DBNAME")

	InitDatabase(
		DatabaseConfig{
			Host: HOST,
			User: USER,
			Password: PASSWORD,
			Port: PORT,
			Database: DBNAME,
		},
	)

	r := gin.Default()

	InitRoutes(r)

	log.Println(":8080")

	r.Run()
}
