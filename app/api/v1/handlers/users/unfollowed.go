package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"twittor-api/app/middleware"
	responseService "twittor-api/app/services/response.service"
	usersService "twittor-api/app/services/users.service"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
)

// UnFollowed GET route /v1/users/unfollowed?page=&term=
func UnFollowed(w http.ResponseWriter, r *http.Request) {
	claim, _ := middleware.AuthClaim(r)
	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	term := r.URL.Query().Get("term")

	if err != nil {
		http.Error(w, "the param page error:  "+err.Error(), http.StatusBadRequest)
	}

	repositoryUser := repositoryFactoryUser.Build()
	service := usersService.New(repositoryUser)

	users, errUsers := service.UnFollowed(claim.ID.Hex(), page, term)

	if errUsers != nil {
		http.Error(w, "error to users unfollowed "+errUsers.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseService.Call(true, "", users))
}
