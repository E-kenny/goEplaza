package database

import (
	"fmt"
	"github.com/E-kenny/eplaza"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"os"
)

type SqlUserService struct {
	DB *sqlx.DB
}

//signing and validating key
var hmacSampleSecret []byte

func (dbUser SqlUserService) SignIn(auth eplaza.Auth) (string, error) {
	var dbDetails eplaza.Auth
	//statement
	err := dbUser.DB.Get(&dbDetails, "SELECT id, password FROM users WHERE email=?", auth.Email)
	if err != nil {
		return "User does not exist", err
	} else if dbDetails.Password != auth.Password {
		return "Wrong password", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    dbDetails.Password,
		"password": dbDetails.Password,
	})

	hmacSampleSecret = []byte(os.Getenv("KEY"))
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		return "Can't sign", err
	}

	return tokenString, nil
}

func Auth(tokenString string)(eplaza.Auth, error){
	var Details eplaza.Auth
	hmacSampleSecret = []byte(os.Getenv("KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Details.Email=claims["email"].(string)
		Details.Password=claims["password"].(string)
		
	} else {
		return Details, err
	}

	return Details, nil
}

func (dbUser SqlUserService) CreateUser(user *eplaza.User) error {

	//Get uuid values
	id := fmt.Sprintln(uuid.NewString())
	//SQL query

	_, err := dbUser.DB.NamedExec(`INSERT INTO users (id, firstName, lastName, email, password, role) VALUES (:id, :first, :last, :email, :pass, :role )`,
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

	//statement
	err := dbUser.DB.Get(&user, "SELECT * FROM users WHERE id=?", id)
	if err != nil {

		return user, err
	}

	return user, nil
}

func (dbUser SqlUserService) GetAllUsers() ([]eplaza.User, error) {
	users := []eplaza.User{}

	err := dbUser.DB.Select(&users, "SELECT * FROM users ORDER BY firstName ASC")
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
