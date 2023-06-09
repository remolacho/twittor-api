package followers

import (
	"encoding/json"
	"net/http"
	"twittor-api/app/middleware"
	followService "twittor-api/app/services/follower.service"
	responseService "twittor-api/app/services/response.service"
	repositoryFactoryFollow "twittor-api/infraestructure/repositories/factories/repository.factory.follower"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
)

// Create POST route /v1/followers
func Create(w http.ResponseWriter, r *http.Request) {
	claim, _ := middleware.AuthClaim(r)
	userID := r.URL.Query().Get("followUserId")

	repoUser := repositoryFactoryUser.Build()
	repoRel := repositoryFactoryFollow.Build()
	service := followService.NewCreate(repoUser, repoRel)

	_, err := service.Create(claim.ID.Hex(), userID)

	if err != nil {
		http.Error(w, "error to create follow "+err.Error(), 404)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseService.Call(true, "", nil))
}
