package database

import (
	"fmt"

	"github.com/E-kenny/eplaza"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SqlUserService struct {
	DB *sqlx.DB
}

func (dbUser SqlUserService) CreateUser(user *eplaza.User) error {
	//Get connection
	db := dbUser.DB
	//Get uuid values
	id := fmt.Sprintln(uuid.NewString())
	//SQL query

	_, err := db.NamedExec(`INSERT INTO users (id, firstName, lastName, email, password, role) VALUES (:id, :first, :last, :email, :pass, :role )`,
		map[string]interface{}{
			"id":    id,
			"first": user.FirstName,
			"last":  user.LastName,
			"email": user.Email,
			"pass":  user.Password,
			"role":  user.Role,
		})

	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}

	return nil
}

func (dbUser SqlUserService) GetUser(id string) eplaza.User {
	var user = eplaza.User{}

	//Get connection
	db := dbUser.DB

	//statement
	err := db.Get(&user, "SELECT * FROM users WHERE id=?", id)
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
