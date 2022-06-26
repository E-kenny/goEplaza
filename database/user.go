package database

import (
	"fmt"
	"os"

	"github.com/E-kenny/eplaza"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SqlUserService struct {
	DB *sqlx.DB
}

//signing and validating key
var hmacSampleSecret []byte

//Authenticated data
var authDetail eplaza.Auth


var user eplaza.User

func (dbUser SqlUserService) SignIn(auth eplaza.Auth) (string, error) {
	var dbdetail eplaza.Auth
	//statement
	err := dbUser.DB.Get(&dbdetail, "SELECT id, email, password FROM users WHERE email=?", auth.Email)
	if err != nil {
		return "User does not exist", err
	} else if dbdetail.Password != auth.Password {
		return "Wrong password", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       dbdetail.Id,
		"email":    dbdetail.Email,
		"password": dbdetail.Password,
	})

	hmacSampleSecret = []byte(os.Getenv("KEY"))
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		return "Can't sign", err
	}

	return tokenString, nil
}

func (dbUser SqlUserService) Auth(tokenString string) error {
	hmacSampleSecret = []byte(os.Getenv("KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		authDetail.Email = claims["email"].(string)
		authDetail.Password = claims["password"].(string)
		// fmt.Println(authDetail.Email, authDetail.Password)

	} else {
		return err
	}

	// return detail, nil
	return nil
}

func (dbUser SqlUserService) AuthOne(id string) error {

    err := dbUser.DB.Get(&user, "SELECT * FROM users WHERE id=?", id)

	if err != nil {
		fmt.Println(authDetail.Email)
		return err
	}
	
    if user.Email == authDetail.Email && user.Password == authDetail.Password{
		return nil
	}

	panic("You're not authenticated")
	

}

func (dbUser SqlUserService) SignUp(user *eplaza.User) error {

	//Get uuid values
	id := fmt.Sprintln(uuid.NewString())
	//SQL query

	_, err := dbUser.DB.NamedExec(`INSERT INTO users (id, first_name, last_name, email, password, role) VALUES (:id, :first, :last, :email, :pass, :role )`,
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
	//statement
	err := dbUser.DB.Get(&user, "SELECT * FROM users WHERE id=?", id)
	if err != nil {

		return user, err
	}

	return user, nil
}

func (dbUser SqlUserService) GetAllUsers() ([]*eplaza.User, error) {
	users := []*eplaza.User{}

	err := dbUser.DB.Select(&users, "SELECT * FROM users ORDER BY first_name ASC")
	if err != nil {
		return nil, err
	}
	fmt.Println(users)

	return users, nil
}

func (dbUser SqlUserService) UpdateUser(user *eplaza.User) error {

	_, err := dbUser.DB.Queryx(`UPDATE users SET first_name = ? , last_name = ? WHERE id=?`, user.FirstName, user.LastName, user.Id)

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
