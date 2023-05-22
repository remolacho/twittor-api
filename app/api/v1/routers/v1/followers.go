package v1

import (
	"github.com/gorilla/mux"
	follow "twittor-api/app/api/v1/handlers/followers"
	"twittor-api/app/middleware"
)

func FollowRoutes(router *mux.Router) {
	router.HandleFunc("/v1/followers", middleware.Authorization(follow.Create)).Methods("POST")
	router.HandleFunc("/v1/followers", middleware.Authorization(follow.Destroy)).Methods("DELETE")
	router.HandleFunc("/v1/followed", middleware.Authorization(follow.Followed)).Methods("GET")
}
