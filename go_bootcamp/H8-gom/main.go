package main

import (
	"fmt"
	"go_bootcamp/H8-gom/database"
	"go_bootcamp/H8-gom/models"
)

func main() {
	database.StartDB()

	CreateUser("azwarrasyid@gmail.com")
}

func CreateUser(email string) {
	db := database.GetDB()

	var User = models.User{
		Email: email,
	}

	// fmt.Println(&User)

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New User Data:", User)
}
