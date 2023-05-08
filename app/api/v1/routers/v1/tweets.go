package v1

import (
	"github.com/gorilla/mux"
	tweet "twittor-api/app/api/v1/handlers/tweets"
	"twittor-api/app/middleware"
)

func TweetRoutes(router *mux.Router) {
	router.HandleFunc("/v1/tweet", middleware.Authorization(tweet.Create)).Methods("POST")
	router.HandleFunc("/v1/tweets", middleware.Authorization(tweet.Index)).Methods("GET")
	router.HandleFunc("/v1/tweet", middleware.Authorization(tweet.Delete)).Methods("DELETE")
}
