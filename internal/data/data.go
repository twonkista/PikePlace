package data

import (
	"fmt"

	"github.com/twonkista/PikePlace/internal/models"

	"gorm.io/gorm"
)

func CreateUsers(db *gorm.DB, username string, password string, balance float64, strikes int) error {
	user := models.User{
		UserName:       username,
		HashedPassword: password,
		Balance:        balance,
		Strikes:        strikes,
	}
	result := db.Create(&user)

	if result.Error != nil {
		fmt.Printf("Something went wrong creating user: %v\n", result.Error)
	}

	return nil
}
