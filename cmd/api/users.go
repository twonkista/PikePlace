package main

import (
	"net/http"
	"os/user"

	"github.com/twonkista/PikePlace/internal/data"
)

type struct UserSchema {
	UserName string `json:"username"`
	Password string `json:"password"`
	Balance  float64 `json:"balance"`
	Strikes  int     `json:"strikes"`
}

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request, ) {
	// Parse the request body to get user details
	
}
