package user

import (
	"net/http"
	"twittor-api/app/middleware"
	userService "twittor-api/app/services/user.service"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
	uploadFile "twittor-api/infraestructure/upload.file"
)

// Banner POST route /v1/users/banner-upload
func Banner(w http.ResponseWriter, r *http.Request) {
	claim, _ := middleware.AuthClaim(r)
	file, handler, err := r.FormFile("banner")

	if err != nil {
		http.Error(w, "Error upload banner "+err.Error(), http.StatusBadRequest)
		return
	}

	serviceUpload := uploadFile.New(file, handler, claim.ID.Hex())
	avatar, errOS := serviceUpload.Call("banners")

	if errOS != nil {
		http.Error(w, "Error upload banner "+errOS.Error(), http.StatusBadRequest)
		return
	}

	repository := repositoryFactoryUser.Build()
	serviceAvatar := userService.NewBanner(repository, avatar)
	_, err = serviceAvatar.Upload(claim.ID.Hex())

	if err != nil {
		http.Error(w, "Error upload banner "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
