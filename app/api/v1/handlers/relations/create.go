package relations

import (
	"net/http"
	"twittor-api/app/middleware"
	relationService "twittor-api/app/services/relation.service"
	repositoryFactoryRelation "twittor-api/infraestructure/repositories/factories/repository.factory.relation"
	repositoryFactoryUser "twittor-api/infraestructure/repositories/factories/repository.factory.user"
)

// High POST route /v1/relations/high
func High(w http.ResponseWriter, r *http.Request) {
	claim, _ := middleware.AuthClaim(r)
	userID := r.URL.Query().Get("userId")

	repoUser := repositoryFactoryUser.Build()
	repoRel := repositoryFactoryRelation.Build()
	service := relationService.NewCreate(repoUser, repoRel)

	flag, err := service.Create(claim.ID.Hex(), userID)

	if !flag {
		http.Error(w, "error to create relation following "+err.Error(), 404)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
