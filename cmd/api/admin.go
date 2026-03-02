package main

import (
	"encoding/json"
	"net/http"

	"github.com/twonkista/PikePlace/internal/models"
)

func (app *application) approvePoolHandler(w http.ResponseWriter, r *http.Request) {
	verdict := r.FormValue("verdict")
	pool := r.FormValue("pool")
	if verdict != "approve" && verdict != "reject" {
		http.Error(w, "Invalid verdict", http.StatusBadRequest)
		return
	}

	result := app.db.Model(&models.User{}).Where("id = ?", pool).Update("approve", verdict)
	if result.Error != nil {
		http.Error(w, "Failed to update approval status", http.StatusInternalServerError)
		return
	}
}

func (app *application) listPendingPoolsHandler(w http.ResponseWriter, r *http.Request) {
	var pendingPools []models.User
	// Fetch all pending pools from the database
	app.db.Where("approve = ?", "pending").Find(&pendingPools)
	jsonData, err := json.Marshal(pendingPools)
	if err != nil {
		http.Error(w, "Error serializing pending pools", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
