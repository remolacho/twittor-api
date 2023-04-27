package user

import (
	"encoding/json"
	"net/http"
	"twittor-api/application/services/userService"
	"twittor-api/domain/models/user"
)

// SignUp POST route /v1/users/sign-up
func SignUp(w http.ResponseWriter, r *http.Request) {
	u := user.New()
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "error to create user "+err.Error(), 400)
		return
	}

	service := userService.NewUser()
	errUserService, _ := service.Create(u)

	if errUserService != nil {
		http.Error(w, errUserService.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
