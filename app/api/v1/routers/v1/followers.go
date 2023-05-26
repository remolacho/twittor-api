package v1

import (
	"github.com/gorilla/mux"
	follower "twittor-api/app/api/v1/handlers/followers"
	"twittor-api/app/middleware"
)

func FollowRoutes(router *mux.Router) {
	router.HandleFunc("/v1/followers", middleware.Authorization(follower.Create)).Methods("POST")
	router.HandleFunc("/v1/followers", middleware.Authorization(follower.Destroy)).Methods("DELETE")
	router.HandleFunc("/v1/followed", middleware.Authorization(follower.Followed)).Methods("GET")
	router.HandleFunc("/v1/followers/tweets", middleware.Authorization(follower.Tweets)).Methods("GET")

}
