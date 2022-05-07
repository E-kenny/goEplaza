package eplaza

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	Id         string `json:"id" db:"id"`
	FirstName  string `json:"firstName" db:"first_name"`
	LastName   string `json:"lastName" db:"last_name"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"password" db:"password"`
	Role       string `json:"role" db:"role"`
	Created_at string `json:"created_at" db:"created_at,omitempty"`
	Updated_at string `json:"updated_at" db:"updated_at,omitempty"`
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
	GetUser(id string) (User, error)
	//Get all users
	GetAllUsers() ([]User, error)
	//update user
	UpdateUser(id *User) error
	//delete user
	DeleteUser(id string) error
	//Sign in user
	SignIn(user Auth) string
}

//claims object
type Auth struct {
	Id       string
	Email    string
	Password string
}
