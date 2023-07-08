package tweets

import (
	"encoding/json"
	"net/http"
	"twittor-api/app/middleware"
	responseService "twittor-api/app/services/response.service"
	tweetService "twittor-api/app/services/tweet.service"
	repositoryFactoryTweet "twittor-api/infraestructure/repositories/factories/repository.factory.tweet"
)

// Delete DELETE route /v1/tweet?tweetID=
func Delete(w http.ResponseWriter, r *http.Request) {
	tweetID := r.URL.Query().Get("tweetId")
	claim, _ := middleware.AuthClaim(r)

	repositoryTweet := repositoryFactoryTweet.Build()
	service := tweetService.NewDelete(repositoryTweet)

	_, err := service.Delete(tweetID, claim.ID.Hex())

	if err != nil {
		http.Error(w, "error to delete tweet "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseService.Call(true, "", nil))
}
