package tweets

import (
	"encoding/json"
	"net/http"
	"twittor-api/app/middleware"
	tweetService "twittor-api/app/services/tweet.service"
	"twittor-api/domain/models/tweet"
	repositoryFactoryTweet "twittor-api/infraestructure/repositories/factories/repository.factory.tweet"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
)

// Create POST route /v1/tweet
func Create(w http.ResponseWriter, r *http.Request) {
	claim, _ := middleware.AuthClaim(r)
	t := tweet.New(claim.ID.Hex())
	err := json.NewDecoder(r.Body).Decode(t)

	if err != nil {
		http.Error(w, "error to create tweet "+err.Error(), 400)
		return
	}

	repositoryTweet := repositoryFactoryTweet.Build()
	repositoryUser := repositoryFactoryUser.Build()
	service := tweetService.NewTweet(repositoryTweet, repositoryUser)

	_, err = service.Create(t)

	if err != nil {
		http.Error(w, "error to create tweet "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
