package routers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"twittor-api/application/api/v1/handlers/users"
	"twittor-api/application/middlewr"
)

func Handler() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/users/sign-up", middlewr.CheckDB(user.SignUp)).Methods("POST")

	port := os.Getenv("PORT")
	corsAllowed := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, corsAllowed))
}
