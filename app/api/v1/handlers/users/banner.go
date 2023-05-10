package user

import (
	"encoding/json"
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
	banner, errOS := serviceUpload.Call("banners")

	if errOS != nil {
		http.Error(w, "Error upload banner "+errOS.Error(), http.StatusBadRequest)
		return
	}

	repository := repositoryFactoryUser.Build()
	serviceBanner := userService.NewBanner(repository, banner)
	_, err = serviceBanner.Upload(claim.ID.Hex())

	if err != nil {
		http.Error(w, "Error upload banner "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetBanner GET route /v1/users/banner?userId
func GetBanner(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	repository := repositoryFactoryUser.Build()
	serviceBanner := userService.NewBanner(repository)

	_, response, err := serviceBanner.Get(userID)

	if err != nil {
		http.Error(w, "Error banner "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
