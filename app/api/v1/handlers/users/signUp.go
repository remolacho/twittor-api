package user

import (
	"encoding/json"
	"net/http"
	"twittor-api/app/services/userService"
	"twittor-api/domain/models/user"
	"twittor-api/infraestructure/repositories/factories/repositoryFactoryUser"
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
