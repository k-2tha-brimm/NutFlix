package userrepository

import (
	"database/sql"
	"log"

	"../models"
)

// UserRepository struct
type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Signup will sign the user up and add them into the db
func (u UserRepository) Signup(db *sql.DB, user models.User) models.User {
	sqlStatement := `
		INSERT INTO users (username, password, email)
		VALUES ($1, $2, $3)
		RETURNING id;`

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
