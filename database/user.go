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

func (dbUser SqlUserService) CreateUser(user *eplaza.User) error {
	//Get connection
	db := dbUser.DB
	//Get uuid values
	id := fmt.Sprintln(uuid.NewString())
	//SQL query
	stmt, err := db.Prepare("INSERT INTO users (Id, firstName, lastName, email, password, role) VALUES (?, ?, ?, ?, ?, ? )")
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	_, err = stmt.Exec(id, user.FirstName, user.LastName, user.Email, user.Password, user.Role)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	return nil
}
func (dbUser SqlUserService) GetUser(id int) eplaza.User {
	var user eplaza.User
	//Get connection
	db := dbUser.DB
	//statement
	stmt, err := db.Prepare("SELECT * FROM users WHERE id=?")
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	err = stmt.QueryRow(id).Scan(user.Id, user.FirstName, user.LastName, user.Email, user.Role, user.Created_at, user.Updated_at)
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	return user
}

func (dbUser SqlUserService) GetAllUsers() []eplaza.User {

	return []eplaza.User{}
}

func (dbUser SqlUserService) UpdateUser(id int) error {

	return nil
}
func (dbUser SqlUserService) DeleteUser(id int) error {

	return nil
}
