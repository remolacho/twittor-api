package user

import (
	"net/http"
	"twittor-api/app/middleware"
	userService "twittor-api/app/services/user.service"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
	uploadFile "twittor-api/infraestructure/upload.file"
)

// Avatar POST route /v1/users/avatar-upload
func Avatar(w http.ResponseWriter, r *http.Request) {
	claim, _ := middleware.AuthClaim(r)
	file, handler, err := r.FormFile("avatar")

	if err != nil {
		http.Error(w, "Error upload avatar "+err.Error(), http.StatusBadRequest)
		return
	}

	serviceUpload := uploadFile.New(file, handler, claim.ID.Hex())
	avatar, errOS := serviceUpload.Call("avatars")

	if errOS != nil {
		http.Error(w, "Error upload avatar "+errOS.Error(), http.StatusBadRequest)
		return
	}

	repository := repositoryFactoryUser.Build()
	serviceAvatar := userService.NewAvatar(repository, avatar)
	_, err = serviceAvatar.Upload(claim.ID.Hex())

	if err != nil {
		http.Error(w, "Error upload avatar "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAvatar GET route /v1/users/avatar?userId
func GetAvatar(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	repository := repositoryFactoryUser.Build()
	serviceAvatar := userService.NewAvatar(repository)

	response, hasAvatar := serviceAvatar.Get(userID)

	if !hasAvatar {
		http.Error(w, "The user not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "image/*")
	http.ServeFile(w, r, response)
}
