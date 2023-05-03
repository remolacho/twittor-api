package user

import (
	"encoding/json"
	"net/http"
	"twittor-api/app/middleware"
	userService "twittor-api/app/services/user.service"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
)

// Profile GET route /v1/users/profile
func Profile(w http.ResponseWriter, r *http.Request) {
	claim, _ := middleware.AuthClaim(r)
	repositoryUser := repositoryFactoryUser.Build()
	profile := userService.NewProfile(repositoryUser)
	response, err := profile.Get(claim.ID.Hex())

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
