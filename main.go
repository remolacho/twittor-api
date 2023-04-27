package main

import (
	"log"
	"twittor-api/application/api/v1/routers"
	"twittor-api/infraestructure/db"
)

func main() {
	if db.CurrentSession().Check() == 0 {
		log.Fatal("error connection to DB")
	}

	routers.Handler()
}
