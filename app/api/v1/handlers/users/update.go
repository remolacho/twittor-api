package user

import (
	"encoding/json"
	"net/http"
	"twittor-api/app/middleware"
	responseService "twittor-api/app/services/response.service"
	userService "twittor-api/app/services/user.service"
	"twittor-api/domain/models/user"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
)

// Update PUT route /v1/users/update
func Update(w http.ResponseWriter, r *http.Request) {
	claim, _ := middleware.AuthClaim(r)

	u := user.New()
	err := json.NewDecoder(r.Body).Decode(u)

	if err != nil {
		http.Error(w, "error to update user "+err.Error(), 400)
		return
	}

	repositoryUser := repositoryFactoryUser.Build()
	service := userService.NewUpdate(repositoryUser)
	_, err = service.Update(claim.ID.Hex(), u)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseService.Call(true, "", nil))
}
