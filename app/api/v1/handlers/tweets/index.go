package tweets

import (
	"encoding/json"
	"net/http"
	"strconv"
	tweetService "twittor-api/app/services/tweet.service"
	repositoryFactoryTweet "twittor-api/infraestructure/repositories/factories/repository.factory.tweet"
)

// Index Get route /v1/tweets?userId=&page=1
func Index(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)

	if err != nil {
		http.Error(w, "the param page error:  "+err.Error(), http.StatusBadRequest)
	}

	repository := repositoryFactoryTweet.Build()
	service := tweetService.NewList(repository)
	tweets, errTweets := service.AllByPagedUser(userID, page)

	if errTweets != nil {
		http.Error(w, errTweets.Error(), http.StatusNotFound)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&tweets)
}
