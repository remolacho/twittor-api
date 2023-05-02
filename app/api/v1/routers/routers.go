package routers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"twittor-api/app/api/v1/handlers/users"
)

func Handler() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/users/sign-up", user.SignUp).Methods("POST")
	router.HandleFunc("/v1/users/sign-in", user.SignIn).Methods("POST")

	port := os.Getenv("PORT")
	corsAllowed := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, corsAllowed))
}
