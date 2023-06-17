package user

import (
	"encoding/json"
	"net/http"
	responseService "twittor-api/app/services/response.service"
	userService "twittor-api/app/services/user.service"
	"twittor-api/domain/models/user"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
)

// SignIn POST route /v1/users/sign-in
func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	u := user.New()
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "error to user login "+err.Error(), 403)
		return
	}

	repositoryUser := repositoryFactoryUser.Build()
	service := userService.NewLogin(repositoryUser)
	response, isValid := service.Login(u.Email, u.Password)

	if !isValid {
		http.Error(w, "Thw email or password is not valid!!!", 403)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseService.Call(true, "", response))
}
