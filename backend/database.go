package main

import (
	"fmt"
	"log"

	"database/sql"
	_"github.com/lib/pq"

	"github.com/toxtashev/LogisticApp/utils"
)

type DatabaseConfig struct {
	Host, User, Password, Port, Database string
}

func InitDatabase (config DatabaseConfig) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s",
		config.Host,
		config.User,
		config.Password,
		config.Port,
		config.Database,
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err.Error())
	}

	utils.DB = db
}