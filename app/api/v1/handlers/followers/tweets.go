package followers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"twittor-api/app/middleware"
	followService "twittor-api/app/services/follower.service"
	responseService "twittor-api/app/services/response.service"
	repositoryFactoryFollow "twittor-api/infraestructure/repositories/factories/repository.factory.follower"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
)

// Tweets GET route  /v1/followers/tweets
func Tweets(w http.ResponseWriter, r *http.Request) {
	claim, _ := middleware.AuthClaim(r)
	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)

	if err != nil {
		http.Error(w, "the param page error:  "+err.Error(), http.StatusBadRequest)
	}

	repositoryFollow := repositoryFactoryFollow.Build()
	repositoryUser := repositoryFactoryUser.Build()
	service := followService.NewListTweets(repositoryFollow, repositoryUser)
	response, _, errList := service.ListTweets(claim.ID.Hex(), page)

	if errList != nil {
		http.Error(w, "the list tweets error:  "+errList.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseService.Call(true, "", response))

}
