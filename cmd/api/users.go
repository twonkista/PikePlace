package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func (app *application) updateBalanceHandler(w http.ResponseWriter, r *http.Request) {
	balance, err := strconv.ParseFloat(r.FormValue("balance"), 64)

	if err != nil || balance < 0 || balance > 1000 {
		http.Error(w, "Invalid balance", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID")
	app.db.Model(&models.User{}).Where("id = ?", userID).Update("balance", balance)
	return
}
