package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/E-kenny/eplaza"
	"github.com/E-kenny/eplaza/database"
	"github.com/go-chi/chi/v5"
)

//This handler log in the user
func signIn(w http.ResponseWriter, r *http.Request) {
	//auth struct
	var auth eplaza.Auth
	//Decode the request body
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":` + err.Error() + "}"))
	}
	//Create database connection
	db, err := database.Connection()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":` + err.Error() + "}"))
	}

	//Assign SqlUserService with the db connection
	var userConn eplaza.UserService = database.SqlUserService{
		DB: db,
	}

	defer db.Close()
	//call the SignIn method
	token, err := userConn.SignIn(auth)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":` + err.Error() + ", " + token + "}"))

	} else {
		w.Write([]byte(`{"message":` + token + `}`))
	}

}

//This is the middleware That authorizes users
func auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RequestToken := r.Header.Get("Authorization")
		RequestToken = strings.Split(RequestToken, "Bearer ")[1]
		db, err := database.Connection()
		if err != nil {
			fmt.Sprintln("%w", err)
		}
		//Assign SqlUserService with the db connection
		conn := database.SqlUserService{
			DB: db,
		}

		defer db.Close()

		err = conn.Auth(RequestToken)

		if err != nil {
			log.Fatal(err)
		} else {
			h.ServeHTTP(w, r)
		}

	})
}

//This is the middleware That authorizes users
func authOne(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := database.Connection()
		if err != nil {
			fmt.Sprintln("%w", err)
		}
		//Assign SqlUserService with the db connection
		conn := database.SqlUserService{
			DB: db,
		}

		defer db.Close()
		//extract user id
		userID := chi.URLParam(r, "userID")

		//AuthOne method
		err = conn.AuthOne(userID)

		if err != nil {
			log.Fatal(err)
		} else {
			h.ServeHTTP(w, r)
		}

	})
}

//This handler creates a new user
func signUp(w http.ResponseWriter, r *http.Request) {

	//User structk
	var user eplaza.User
	//Decode the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":` + err.Error() + "}"))

	}
	//Validate the data
	err = user.Validate()

	//Create database connection//authentication
	// Auth(data string)error
	// //authentication
	// AuthOne(data string)error
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":` + err.Error() + "}"))
	} else {
		db, err := database.Connection()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(`{"message":` + err.Error() + "}"))

		}
		//Assign SqlUserService with the db connection
		var userConn eplaza.UserService = database.SqlUserService{
			DB: db,
		}
		defer db.Close()

		//Create User
		err = userConn.SignUp(user)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(`{"message":` + err.Error() + "}"))

		} else {
			w.Write([]byte(`{"message": "User created successfully"}`))
		}

	}

}

//This handler gets a user
func getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	// fetch `"key"` from the request context
	// ctx := r.Context()//authentication
	// key := ctx.Value("key").(string)

	db, err := database.Connection()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":` + err.Error() + "}"))

	}

	//Assign SqlUserService with the db connection
	var userConn eplaza.UserService = database.SqlUserService{
		DB: db,
	}

	//Get A User
	user, err := userConn.GetUser(userID)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":` + err.Error() + "}"))
	} else {
		// respond to the client
		json.NewEncoder(w).Encode(user)
	}

}

//This handler all usersAuthOne(data string)error
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connection()
	if err != nil {
		w.WriteHeader(500) //Assign SqlUserService with the db connection
		w.Write([]byte(`{"message":` + err.Error() + "}"))
	}

	//Assign SqlUserService with the db connection
	var userConn eplaza.UserService = database.SqlUserService{
		DB: db,
	}
	users, err := userConn.GetAllUsers()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":` + err.Error() + "}"))

	} else {
		json.NewEncoder(w).Encode(users)
	}
}

//This handler updates a user
func updateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	// fetch `"key"` from the request context
	// ctx := r.Context()
	// key := ctx.Value("key").(string)

	var user eplaza.User
	//Decode the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":` + err.Error() + "}"))

	}

	user.Id = userID

	//Create database connection
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":` + err.Error() + "}"))
	} else {

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(`{"message":` + err.Error() + "}"))

		}
	}
	db, err := database.Connection()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":` + err.Error() + "}"))

	}

	//Assign SqlUserService with the db connection
	var userConn eplaza.UserService = database.SqlUserService{
		DB: db,
	}

	err = userConn.UpdateUser(user)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":` + err.Error() + "}"))
	} else {
		w.WriteHeader(200)
		w.Write([]byte(`{"message": "User updated successfully"}`))
	}

}

//This handler deletes user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	// fetch `"key"` from the request context
	// ctx := r.Context()
	// key := ctx.Value("key").(string)

	db, err := database.Connection()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":` + err.Error() + "}"))

	}

	//Assign SqlUserService with the db connection
	var userConn eplaza.UserService = database.SqlUserService{
		DB: db,
	}

	err = userConn.DeleteUser(userID)
	if err != nil {
		w.WriteHeader(200)
		w.Write([]byte(`{"message": "User deleted successfully"}`))
	}

}
