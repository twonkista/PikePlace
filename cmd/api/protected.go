package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) protectedHandler(w http.ResponseWeiter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // check session/csrf token here
        // if invalid:
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
        // if valid:
        next.ServeHTTP(w, r)
    })
}