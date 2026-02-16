package main

import (
	"encoding/json"
	"net/http"
)

type Pool struct {
	ID          uint    "json:\"id\""
	Title       string  "json:\"title\""
	Description string  "json:\"description\""
	CreatorID   uint    "json:\"creator_id\""
	Status      string  "json:\"status\""
	SLTotal     float64 "json:\"sl_total\""
	SWTotal     float64 "json:\"sw_total\""
	Outcome     bool    "json:\"outcome\""
}

func (app *application) listPoolsHandler(w http.ResponseWriter, r *http.Request) {
	var pools []Pool
	// Fetch all pools from the database
	app.db.Find(&pools)

	jsonData, err := json.Marshal(pools)
	if err != nil {
		http.Error(w, "Error serializing pools", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (app *application) openPoolsHandler(w http.ResponseWriter, r *http.Request) {
	var pools []Pool
	// Fetch all pools from the database
	app.db.Where("status = ?", "open").Find(&pools)

	jsonData, err := json.Marshal(pools)
	if err != nil {
		http.Error(w, "Error serializing pools", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (app *application) resolvedPoolsHandler(w http.ResponseWriter, r *http.Request) {
	var pools []Pool
	// Fetch all pools from the database
	app.db.Where("status = ?", "resolved").Find(&pools)

	// Serialize the pools to JSON and write to response
	// (Assuming you have a function toJSON for serialization)
	jsonData, err := json.Marshal(pools)
	if err != nil {
		http.Error(w, "Error serializing pools", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// func (app *application) createNewPoolHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("shi"))
// }
