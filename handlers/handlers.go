package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Handler() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")
	corsAllowed := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, corsAllowed))
}
