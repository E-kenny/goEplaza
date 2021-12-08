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

func (dbUser SqlUserService) GetUser(id string) (eplaza.User, error) {
	var user = eplaza.User{}

	//Get connection
	db := dbUser.DB

	//statement
	err := db.Get(&user, "SELECT * FROM users WHERE id=?", id)
	if err != nil {

		return user, err
	}

	return user, nil
}

func (dbUser SqlUserService) GetAllUsers() ([]eplaza.User, error) {
	users := []eplaza.User{}

	//Get connection
	db := dbUser.DB

	err := db.Select(&users, "SELECT * FROM users ORDER BY firstName ASC")
	if err != nil {
		return nil, err
	}
	fmt.Println(users)

	return users, nil
}

func (dbUser SqlUserService) UpdateUser(user *eplaza.User) error {

	_, err := dbUser.DB.Queryx(`UPDATE users SET firstName = ? , lastName = ? WHERE id=?`, user.FirstName, user.LastName, user.Id)

	if err != nil {
		return err
	}
	return nil
}
func (dbUser SqlUserService) DeleteUser(id string) error {
	_, err := dbUser.DB.Queryx("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
