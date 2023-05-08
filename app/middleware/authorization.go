package middleware

import (
	"net/http"
	"twittor-api/infraestructure/jwt"
)

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := AuthClaim(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func AuthClaim(r *http.Request) (*jwt.Claim, error) {
	tokenKey := r.Header.Get("Authorization")
	return jwt.NewJwt().Decode(tokenKey)
}
