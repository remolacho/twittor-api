package main

import (
	"log"
	"twittor-api/app/api/v1/routers"
	"twittor-api/infraestructure/db/mongoDB"
	"twittor-api/infraestructure/db/postgresSql"
	"twittor-api/infraestructure/environment"
)

func main() {
	if err := environment.Load(); err != nil {
		log.Fatal("Error loading .env upload")
	}

	if mongoDB.CurrentSession().Check() == 0 {
		log.Fatal("error connection to DB")
	}

	postgresSql.CurrentSession()

	routers.Handler()
}
