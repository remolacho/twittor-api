package tweets

import (
	"net/http"
)

// Delete DELETE route /v1/tweet?tweetID=
func Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
