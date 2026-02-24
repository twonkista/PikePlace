package main

import (
	"net/http"
)

func (app *application) protectedCheckHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check session/csrf token here
		// if invalid:
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		next.ServeHTTP(w, r)
	})
}
