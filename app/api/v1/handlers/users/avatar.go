package user

import (
	"net/http"
	"twittor-api/app/middleware"
	uploadService "twittor-api/app/services/upload.service"
	userService "twittor-api/app/services/user.service"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
)

func Avatar(w http.ResponseWriter, r *http.Request) {
	claim, _ := middleware.AuthClaim(r)
	file, handler, err := r.FormFile("avatar")

	if err != nil {
		http.Error(w, "Error upload.service avatar "+err.Error(), http.StatusBadRequest)
		return
	}

	serviceUpload := uploadService.New(file, handler, claim.ID.Hex())
	avatar, errOS := serviceUpload.Call("avatars")

	if errOS != nil {
		http.Error(w, "Error upload.service avatar "+errOS.Error(), http.StatusBadRequest)
		return
	}

	repository := repositoryFactoryUser.Build()
	serviceAvatar := userService.NewAvatar(repository, avatar)
	err = serviceAvatar.Upload(claim.ID.Hex())

	if err != nil {
		http.Error(w, "Error upload.service avatar "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
