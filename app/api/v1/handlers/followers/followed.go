package followers

import (
	"encoding/json"
	"net/http"
	"twittor-api/app/middleware"
	followService "twittor-api/app/services/follower.service"
	responseService "twittor-api/app/services/response.service"
	repositoryFactoryFollow "twittor-api/infraestructure/repositories/factories/repository.factory.follower"
)

// Followed GET route /v1/followed
func Followed(w http.ResponseWriter, r *http.Request) {
	claim, _ := middleware.AuthClaim(r)
	followerID := r.URL.Query().Get("followUserId")

	repoRel := repositoryFactoryFollow.Build()
	service := followService.NewFollower(repoRel)
	response, _ := service.Followed(claim.ID.Hex(), followerID)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseService.Call(true, "", response))
}
