package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/E-kenny/eplaza"
	"github.com/google/uuid"
)

type SqlUserService struct {
	DB *sql.DB
}

func (dbUser SqlUserService) CreateUser(user *eplaza.User) error {
	//Get connection
	db := dbUser.DB
	//Get uuid values
	id := fmt.Sprintln(uuid.NewString())
	created_at := fmt.Sprintln(time.Now())
	updated_at := fmt.Sprintln(time.Now())
	//SQL query
	stmt, err := db.Prepare("INSERT INTO users (Id, firstName, lastName, email, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	_, err = stmt.Exec(id, user.FirstName, user.LastName, user.Email, user.Role, created_at, updated_at)
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
