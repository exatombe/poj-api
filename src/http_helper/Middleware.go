package httphelper

import (
	"net/http"
)

// Middleware is a function that checks if the request has a valid Authorization header
// If it doesn't, it returns a 401 Unauthorized status
// If it does, it calls the next handler in the chain
func Middleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if r.Header.Get("Authorization") == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	next(w, r)
}
