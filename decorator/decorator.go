package decorator

import (
	"consumer-interfaces/db"
	"consumer-interfaces/rabbit"
	"fmt"
)

type UserRepoWithQueue struct {
	dbRepo *db.UserRepo
	queue  *rabbit.Rabbit
}

func NewUserRepoWithQueue(dbRepo *db.UserRepo, queue *rabbit.Rabbit) *UserRepoWithQueue {
	return &UserRepoWithQueue{dbRepo, queue}
}

func (u *UserRepoWithQueue) CreateUser(username, password string) *db.User {
	user := u.dbRepo.CreateUser(username, password)
	u.queue.PublishMessage(fmt.Sprintf("User created with ID: %d, username: %s, password: %s", user.ID, user.Username, user.Password))
	return user
}
