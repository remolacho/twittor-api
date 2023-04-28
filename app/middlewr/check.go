package middlewr

import (
	"net/http"
	"twittor-api/infraestructure/db/mongoDB"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if mongoDB.CurrentSession().Check() == 0 {
			http.Error(w, "the connection with DB is lose!!!!", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
