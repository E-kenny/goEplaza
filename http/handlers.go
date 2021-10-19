package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/E-kenny/eplaza"
	"github.com/E-kenny/eplaza/database"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	//User struct
	var user eplaza.User
	//Decode the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
	//Create database connection
	db, err := database.Connection()
	if err != nil {
		fmt.Println(err.Error())
	}
	//Assign SqlUserService with the db connection
	us := database.SqlUserService{
		DB: db,
	}
	defer db.Close()
	//declare a UserService interface variable
	var conn eplaza.UserService = us

	//Create User
	err = conn.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())

	} else {
		w.Write([]byte("User Created Successfully!"))
	}

}
