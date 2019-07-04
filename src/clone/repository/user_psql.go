package UserRepository

import (
	"database/sql"
	"log"
	"NutFlux/models"
)

type userRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// SignUp will sign the user up and add them into the db
func (u UserRepository) SignUp(db *sql.DB, user models.User) models.User {
	sqlStatement := `
		INSERT INTO users (username, password_digest, email)
		VALUES ($1, $2, $3)
		RETURNING id`

	err := db.QueryRow(sqlStatement, user.Username, user.Password, user.Email).Scan(&user.ID)
	if err != nil {
		panic(err)
	}

	user.Password = ""
	return user
}

// Login will authenticate the user
func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("SELECT * FROM users where email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}