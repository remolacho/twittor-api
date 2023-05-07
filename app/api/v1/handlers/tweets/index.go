package tweets

import (
	"encoding/json"
	"net/http"
)

// Index Get route /v1/tweets
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	var response = struct{ test string }{test: "all tweets"}
	json.NewEncoder(w).Encode(&response)
}
