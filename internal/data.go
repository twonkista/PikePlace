// we are literally making users out of thin fucking air
package main

import (
	"fmt"
)

func createUsers() {
	user := User{
		UserName: "The Big Stein",
		Password: "financier123",
		Balance:  600000000,
		Strikes:  0,
	}
	result := db.Create(&user)

	if result.Error != nil {
		fmt.Printf("Something went wrong creating user: %v\n", result.Error)
	}
}
