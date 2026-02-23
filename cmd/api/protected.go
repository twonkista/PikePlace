package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) protectedHandler(w http.ResponseWriter, r *http.Request) {
    ctx := context.WithValue(r.Context(), "userID", user.ID)
    next.ServeHTTP(w, r.WithContext(ctx))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // check session/csrf token here
        // if invalid:
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
        // if valid:
        next.ServeHTTP(w, r)
    })
}