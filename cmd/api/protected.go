package main

import (
	"context"
	"log"
	"net/http"

	"github.com/twonkista/PikePlace/internal/models"
)

func (app *application) protectedCheckHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		csrfToken := r.Header.Get("csrf-token")
		log.Printf("Token received: %s", csrfToken)

		if csrfToken == "" {
			log.Printf("Token is empty")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var user models.User
		if err := app.db.Where("csrf_token = ?", csrfToken).First(&user).Error; err != nil {
			log.Printf("DB lookup failed: %v", err)
			http.Error(w, "Invalid CSRF token", http.StatusUnauthorized)
			return
		}

		log.Printf("User found: %d", user.ID)
		ctx := context.WithValue(r.Context(), "userID", user.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func (app *application) adminHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		csrfToken := r.Header.Get("csrf-token")
		log.Printf("Token received: %s", csrfToken)

		if csrfToken == "" {
			log.Printf("Token is empty")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var user models.User
		if err := app.db.Where("csrf_token = ?", csrfToken).First(&user).Error; err != nil {
			log.Printf("DB lookup failed: %v", err)
			http.Error(w, "Invalid CSRF token", http.StatusUnauthorized)
			return
		}

		if user.Role != "admin" {
			log.Printf("User %d is not an admin", user.ID)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		log.Printf("User found: %d", user.ID)
		ctx := context.WithValue(r.Context(), "userID", user.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
