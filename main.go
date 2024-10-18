package main

import (
	"consumer-interfaces/db"
	"consumer-interfaces/services"
	"fmt"
	"log"
)

func main() {

	repo := db.NewUserRepo()
	userService := services.NewUserService(repo)

	fmt.Println("Start")

	username := "john"
	password := "123456"

	err := userService.CreateNewUser(username, password)
	if err != nil {
		log.Println(fmt.Sprintf("Error creating new user: %s", err))
	}

	fmt.Println("End")

}
