package eplaza

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func (a User) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required),
		// Password cannot be empty, and the length must between 5 and 50

		validation.Field(&a.Password, validation.Required, validation.Length(7, 50)),
	)
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
