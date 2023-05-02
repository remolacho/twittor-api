package middlewr

import (
	"net/http"
	jwtService "twittor-api/app/services/jwt.service"
)

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service := jwtService.NewJwt()
		_, err := service.Decode(w.Header().Get("Authorization"))

		if err != nil {
			http.Error(w, "the JWT is not valid !!!! "+err.Error(), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
}
