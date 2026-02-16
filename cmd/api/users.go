package main

import (
	"encoding/json"
	"net/http"

	"github.com/twonkista/PikePlace/internal/models"
)

func (app *application) listUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	// Fetch all users from the database
	app.db.Find(&users)

	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error serializing users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
