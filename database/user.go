package database

import (
	"database/sql"
	"fmt"

	"github.com/E-kenny/eplaza"
	"github.com/google/uuid"
)

type SqlUserService struct {
	DB *sql.DB
}

func (user SqlUserService) CreateUser(u *eplaza.User) error {
	//Get connection
	db := user.DB
	//Get uuid values
	id := fmt.Sprintln(uuid.New().String())
	//created_at:=fmt.Sprintln(uuid.)
	//updated_at:=fmt.Sprintln(uuid.)
	//SQL query
	stmt, err := db.Prepare("INSERT INTO user (Id, firstName, lastName, email, profile , status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	return nil
}
func (user SqlUserService) GetUser(id int) eplaza.User {

	return eplaza.User{}
}
func (user SqlUserService) GetAllUsers() []eplaza.User {

	return []eplaza.User{}
}

func (user SqlUserService) UpdateUser(id int) error {

	return nil
}
func (user SqlUserService) DeleteUser(id int) error {

	return nil
}
