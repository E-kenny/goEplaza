package eplaza

import (
	"github.com/google/uuid"
)

type User struct {
	Id                          uuid.UUID
	FirstName, LastName         string
	Phone                       int
	Email, ProfilePhoto, Status string

	Created_at, Updated_at uuid.Time
}

type UserService interface {
	//create user
	CreateUser(u *User) error
	//get user
	GetUser(id int) User
	//Get all users
	GetAllUsers() []User
	//update user
	UpdateUser(id int) error
	//delete user
	DeleteUser(id int) error
}
