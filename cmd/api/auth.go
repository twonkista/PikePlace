package main

//fuck clerk

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"crypto/rand"
	"encoding/base64"

	"github.com/twonkista/PikePlace/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	HashPass     string
	CSRF         string
	SessionToken string
}

func (app *application) loginHandeler(w http.ResponseWriter, r *http.Request) {
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

	if err := app.db.Where("user_name = ?", user.UserName).First(&user).Error; err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if !comparehashpassword(user.HashedPassword, hashpass) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	sessionToken := generateToken(24)
	csrfToken := generateToken(24)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(12 * time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(12 * time.Hour),
		HttpOnly: true,
	})

	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	if err := app.db.Save(&user).Error; err != nil {
		log.Printf("Failed to save tokens for user %s: %v", user.UserName, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Login successful")
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
func comparehashpassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func generateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Error generating session token %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
