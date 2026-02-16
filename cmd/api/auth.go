package main

//fuck clerk

import (
	"fmt"
	"net/http"

	"github.com/twonkista/PikePlace/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	HashPass     string
	CSRF         string
	SessionToken string
}

func (app *application) loginHandeler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login handler called")
}

func (app *application) registrationHandeler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	hashpass := r.FormValue("password")

	user := models.User{
		UserName:       username, // Uses the variable declared above
		HashedPassword: hashpass, // Logic check: You'll hash this soon
	}

	if len(user.UserName) < 5 || (len(user.HashedPassword) < 8) {
		http.Error(w, "Invalid Username/Password", http.StatusBadRequest)
		return
	}

	if err := app.db.Where("user_name = ?", user.UserName).First(&models.User{}).Error; err == nil {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := hashPassword(user.HashedPassword)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.HashedPassword = hashedPassword

	if err := app.db.Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}

func (app *application) logoutHandeler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logout handler called")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
