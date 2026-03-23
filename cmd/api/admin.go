package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/twonkista/PikePlace/internal/models"
)

func (app *application) approvePoolHandler(w http.ResponseWriter, r *http.Request) {
	verdict := r.FormValue("verdict")
	poolID := r.FormValue("pool")

	if verdict != "approved" && verdict != "rejected" {
		http.Error(w, "Invalid verdict (use 'approved' or 'rejected')", http.StatusBadRequest)
		return
	}

	// Get admin username from context (set by adminHandler middleware)
	username, ok := r.Context().Value("username").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	result := app.db.Model(&models.Pool{}).Where("id = ?", poolID).Updates(map[string]interface{}{
		"approved":     verdict,
		"approved_by":  username,
		"approved_at":  time.Now(),
	})
	if result.Error != nil {
		http.Error(w, "Failed to update approval status", http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Pool not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pool " + verdict + " by " + username))
}

func (app *application) listPendingPoolsHandler(w http.ResponseWriter, r *http.Request) {
	var pendingPools []models.Pool
	app.db.Where("approved = ?", "pending").Find(&pendingPools)
	jsonData, err := json.Marshal(pendingPools)
	if err != nil {
		http.Error(w, "Error serializing pending pools", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
