package v1

import (
	"github.com/gorilla/mux"
	user "twittor-api/app/api/v1/handlers/users"
	"twittor-api/app/middleware"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/v1/users/sign-up", user.SignUp).Methods("POST")
	router.HandleFunc("/v1/users/sign-in", user.SignIn).Methods("POST")
	router.HandleFunc("/v1/users/profile", middleware.Authorization(user.Profile)).Methods("GET")
	router.HandleFunc("/v1/users/update", middleware.Authorization(user.Update)).Methods("PUT")
	router.HandleFunc("/v1/users/avatar", middleware.Authorization(user.Avatar)).Methods("PUT")
}
