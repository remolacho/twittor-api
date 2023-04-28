package main

import (
	"github.com/joho/godotenv"
	"log"
	"twittor-api/app/api/v1/routers"
	"twittor-api/infraestructure/db/mongoDB"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if mongoDB.CurrentSession().Check() == 0 {
		log.Fatal("error connection to DB")
	}

	routers.Handler()
}
