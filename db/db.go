package db

import (
	"log"
	"math/rand"
)

type User struct {
	ID       int
	Username string
	Password string
}

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (u *UserRepo) CreateUser(username, password string) *User {
	id := rand.Int()

	log.Printf("Save user to database with id: %d, username: %s, password: %s\n", id, username, password)

	return &User{
		ID:       id,
		Username: username,
		Password: password,
	}
}

func (u *UserRepo) GetUser(id int) {
	log.Printf("Get user from database with id: %d\n", id)
}

func (u *UserRepo) UpdateUser(user User) {
	log.Printf("Update user in database with id: %d, username: %s, password: %s\n", user.ID, user.Username, user.Password)
}

func (u *UserRepo) DeleteUser(id int) {
	log.Printf("Delete user from database with id: %d\n", id)
}
