package routers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"twittor-api/app/api/v1/routers/v1"
)

func Handler() {
	router := mux.NewRouter()

	v1.UserRoutes(router)
	v1.TweetRoutes(router)

	port := os.Getenv("PORT")
	corsAllowed := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+port, corsAllowed))
}
