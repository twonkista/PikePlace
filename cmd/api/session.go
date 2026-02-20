package main

import (
	"errors"
	"net/http"

	"github.com/twonkista/PikePlace/internal/models"
)

var autherror = errors.New("authentication error")

func (app *application) Authorize(r *http.Request) error {
	username := r.FormValue("username")

	user := models.User{
		UserName: username, // Uses the variable declared above
	}

	if err := app.db.Where("user_name = ?", user.UserName).First(&user).Error; err != nil {
		return autherror
	}

	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		return autherror
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf == "" {
		return autherror
	}
	return nil
}
