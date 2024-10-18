package main

import (
	"consumer-interfaces/db"
	"consumer-interfaces/rabbit"
	"consumer-interfaces/services"
	"fmt"
	"log"
)

func main() {

	repo := db.NewUserRepo()
	queue := rabbit.NewRabbit()
	userService := services.NewUserService(repo, queue)

	fmt.Println("Start")

	username := "john"
	password := "123456"

	err := userService.CreateNewUser(username, password)
	if err != nil {
		log.Println(fmt.Sprintf("Error creating new user: %s", err))
	}

	fmt.Println("End")

}
