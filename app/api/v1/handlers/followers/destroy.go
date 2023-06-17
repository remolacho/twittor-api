package followers

import (
	"encoding/json"
	"net/http"
	"twittor-api/app/middleware"
	followService "twittor-api/app/services/follower.service"
	responseService "twittor-api/app/services/response.service"
	repositoryFactoryFollow "twittor-api/infraestructure/repositories/factories/repository.factory.follower"
)

// Destroy DELETE route /v1/followers
func Destroy(w http.ResponseWriter, r *http.Request) {
	claim, _ := middleware.AuthClaim(r)
	followID := r.URL.Query().Get("followId")

	repoRel := repositoryFactoryFollow.Build()
	service := followService.NewDelete(repoRel)

	_, err := service.Destroy(followID, claim.ID.Hex())

	if err != nil {
		http.Error(w, "error to delete follow "+err.Error(), 404)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseService.Call(true, "", nil))
}
