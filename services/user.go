package services

import (
	"consumer-interfaces/db"
	"fmt"
	"log"
)

type userRepo interface {
	CreateUser(username, password string) *db.User
}

type UserService struct {
	repo userRepo
}

func NewUserService(repo userRepo) *UserService {
	return &UserService{repo}
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

	return nil
}
