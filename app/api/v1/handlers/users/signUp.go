package user

import (
	"encoding/json"
	"net/http"
	userService "twittor-api/app/services/user.service"
	"twittor-api/domain/models/user"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
)

// SignUp POST route /v1/users/sign-up
func SignUp(w http.ResponseWriter, r *http.Request) {
	u := user.New()
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "error to create user "+err.Error(), 400)
		return
	}

	repositoryUser := repositoryFactoryUser.Build()
	service := userService.NewUser(repositoryUser)
	_, errService := service.Create(u)

	if errService != nil {
		http.Error(w, errService.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
