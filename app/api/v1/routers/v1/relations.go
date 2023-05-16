package v1

import (
	"github.com/gorilla/mux"
	relation "twittor-api/app/api/v1/handlers/relations"
	"twittor-api/app/middleware"
)

func RelationRoutes(router *mux.Router) {
	router.HandleFunc("/v1/relations/high", middleware.Authorization(relation.High)).Methods("POST")
}
