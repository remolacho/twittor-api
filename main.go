package main

import (
	"github.com/joho/godotenv"
	"log"
	"twittor-api/application/api/v1/routers"
	"twittor-api/infraestructure/mongoDB"
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
