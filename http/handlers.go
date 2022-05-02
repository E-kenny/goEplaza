package http

import (
	"encoding/json"
	"net/http"

	"github.com/E-kenny/eplaza"
	"github.com/E-kenny/eplaza/database"
	"github.com/go-chi/chi/v5"
)

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
	conn := database.SqlUserService{
		DB: db,
	}
	defer db.Close()
	//call the SignIn method
	token, err := conn.SignIn(auth)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":` + err.Error() + "}"))

	} else {
		w.Write([]byte(token))
	}

}

func createUser(w http.ResponseWriter, r *http.Request) {

	//User struct
	var user eplaza.User
	//Decode the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":` + err.Error() + "}"))

	}
	//Validate the data
	err = user.Validate()

	//Create database connection
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
		conn := database.SqlUserService{
			DB: db,
		}
		defer db.Close()

		//Create User
		err = conn.CreateUser(&user)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(`{"message":` + err.Error() + "}"))

		} else {
			w.Write([]byte(`{"message": "User created successfully"}`))
		}

	}

}

func getUser(w http.ResponseWriter, r *http.Request) {
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
	conn := database.SqlUserService{
		DB: db,
	}

	//Get A User
	user, err := conn.GetUser(userID)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":` + err.Error() + "}"))
	} else {
		// respond to the client
		json.NewEncoder(w).Encode(user)
	}

}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connection()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":` + err.Error() + "}"))

	}

	conn := database.SqlUserService{
		DB: db,
	}
	users, err := conn.GetAllUsers()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":` + err.Error() + "}"))

	} else {
		json.NewEncoder(w).Encode(users)
	}
}

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
	conn := database.SqlUserService{
		DB: db,
	}

	err = conn.UpdateUser(&user)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":` + err.Error() + "}"))
	} else {
		w.WriteHeader(200)
		w.Write([]byte(`{"message": "User updated successfully"}`))
	}

}

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
	conn := database.SqlUserService{
		DB: db,
	}

	err = conn.DeleteUser(userID)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":` + err.Error() + "}"))
	} else {
		w.WriteHeader(200)
		w.Write([]byte(`{"message": "User deleted successfully"}`))
	}

}
