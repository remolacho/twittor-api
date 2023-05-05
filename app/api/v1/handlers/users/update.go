package user

import (
	"net/http"
	"twittor-api/app/middleware"
)

func Update(w http.ResponseWriter, r *http.Request) {
	claim, _ := middleware.AuthClaim(r)
	print(claim)
}
