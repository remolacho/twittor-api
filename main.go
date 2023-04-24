package main

import (
	"log"
	"twittor-api/db"
	"twittor-api/handlers"
)

func main() {
	if db.CheckConn() == 0 {
		log.Fatal("error connection to DB")
	}

	handlers.Handler()
}
