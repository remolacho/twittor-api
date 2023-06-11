package user

import (
	"encoding/json"
	"net/http"
	"twittor-api/app/middleware"
	responseService "twittor-api/app/services/response.service"
	userService "twittor-api/app/services/user.service"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
)

// Profile GET route /v1/users/profile?userId
func Profile(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	claim, _ := middleware.AuthClaim(r)
	repositoryUser := repositoryFactoryUser.Build()
	profile := userService.NewProfile(repositoryUser)

	if len(userId) == 0 {
		userId = claim.ID.Hex()
	}

	response, err := profile.Get(userId)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseService.Call(true, "", response))
}
