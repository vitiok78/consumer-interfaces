package services

import (
	"consumer-interfaces/db"
	"consumer-interfaces/rabbit"
	"fmt"
	"log"
)

type UserService struct {
	repo  *db.UserRepo
	queue *rabbit.Rabbit
}

func NewUserService(repo *db.UserRepo, queue *rabbit.Rabbit) *UserService {
	return &UserService{repo, queue}
}

func (s *UserService) CreateNewUser(username, password string) error {
	if len(username) == 0 {
		return fmt.Errorf("username cannot be empty")
	}

	if len(password) == 0 {
		return fmt.Errorf("password cannot be empty")
	}

	if len(username) < 3 {
		return fmt.Errorf("username must be at least 3 characters long")
	}

	if len(password) < 6 {
		return fmt.Errorf("password must be at least 6 characters long")
	}

	log.Printf("Starting Create new user with username: %s, password: %s\n", username, password)
	user := s.repo.CreateUser(username, password)
	log.Printf("Completed Create new user with ID: %d, username: %s, password: %s\n", user.ID, user.Username, user.Password)

	s.queue.PublishMessage(fmt.Sprintf("event: user-created ID: %d, username: %s", user.ID, user.Username))

	return nil
}
