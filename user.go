package eplaza

import (
	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	Created_at uuid.Time `json:"created_at"`
	Updated_at uuid.Time `json:"updated_at"`
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
