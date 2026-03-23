package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/twonkista/PikePlace/internal/models"
)

// type Pool struct {
// 	ID          uint    "json:\"id\""
// 	Title       string  "json:\"title\""
// 	Description string  "json:\"description\""
// 	CreatorID   uint    "json:\"creator_id\""
// 	Status      string  "json:\"status\""
// 	SLTotal     float64 "json:\"sl_total\""
// 	SWTotal     float64 "json:\"sw_total\""
// 	Outcome     bool    "json:\"outcome\""
// }

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

	userID, ok := r.Context().Value("userID").(uint)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	result := app.db.Model(&models.User{}).Where("id = ?", userID).Update("balance", balance)
	if result.Error != nil {
		http.Error(w, "Failed to update balance", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
