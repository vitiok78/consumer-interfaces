package main

import (
	"consumer-interfaces/db"
	"consumer-interfaces/decorator"
	"consumer-interfaces/rabbit"
	"consumer-interfaces/services"
	"fmt"
	"log"
)

func main() {

	repo := db.NewUserRepo()
	queue := rabbit.NewRabbit()

	repoQueue := decorator.NewUserRepoWithQueue(repo, queue)
	userService := services.NewUserService(repoQueue)

	fmt.Println("Start")

	username := "john"
	password := "123456"

	err := userService.CreateNewUser(username, password)
	if err != nil {
		log.Println(fmt.Sprintf("Error creating new user: %s", err))
	}

	fmt.Println("End")

}
