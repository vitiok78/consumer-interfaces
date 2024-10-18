package services

import (
	"consumer-interfaces/db"
	"fmt"
	"log"
	"math/rand"
)

type UserService struct {
	repo *db.UserRepo
}

func NewUserService(repo *db.UserRepo) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateNewUser(username, password string) error {
	id := rand.Int()

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

	log.Printf("Creating new user with id: %d, username: %s, password: %s\n", id, username, password)
	s.repo.CreateUser(id, username, password)

	return nil
}
