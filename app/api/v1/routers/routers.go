package routers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"twittor-api/app/api/v1/handlers/users"
	"twittor-api/app/middleware"
)

func Handler() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/users/sign-up", user.SignUp).Methods("POST")
	router.HandleFunc("/v1/users/sign-in", user.SignIn).Methods("POST")
	router.HandleFunc("/v1/users/profile", middleware.Authorization(user.Profile)).Methods("GET")
	router.HandleFunc("/v1/users/update", middleware.Authorization(user.Update)).Methods("PUT")

	port := os.Getenv("PORT")
	corsAllowed := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, corsAllowed))
}
