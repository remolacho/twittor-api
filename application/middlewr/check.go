package middlewr

import (
	"net/http"
	"twittor-api/infraestructure/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CurrentSession().Check() == 0 {
			http.Error(w, "the connection with DB is lose!!!!", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
