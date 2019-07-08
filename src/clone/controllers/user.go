package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../models"
	userrepository "../repository"
	"../utils"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var users []models.User

// UserController struct
type UserController struct{}

// Login handles returning users
func (c UserController) Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var jwt models.JWT
		var error models.Error

		json.NewDecoder(r.Body).Decode(&user)

		if user.Email == "" {
			error.Message = "Email is missing."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		if user.Password == "" {
			error.Message = "Password is missing."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		password := user.Password

		userRepo := userrepository.UserRepository{}
		user, err := userRepo.Login(db, user)

		log.Println(err)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "The user does not exist"
				utils.RespondWithError(w, http.StatusBadRequest, error)
				return
			} else {
				log.Fatal(err)
			}
		}

		hashedPassword := user.Password

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

		if err != nil {
			error.Message = "Invalid Password"
			utils.RespondWithError(w, http.StatusUnauthorized, error)
			return
		}

		token, err := utils.GenerateToken(user)

		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		jwt.Token = token

		utils.ResponseJSON(w, jwt)
	}
}

// Signup for new users
func (c UserController) Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var error models.Error

		json.NewDecoder(r.Body).Decode(&user)
		fmt.Printf("%+v\n", user)

		if user.Email == "" {
			error.Message = "Email is missing."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		if user.Password == "" {
			error.Message = "Password is missing."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

		if err != nil {
			log.Fatal(err)
		}

		user.Password = string(hash)

		userRepo := userrepository.UserRepository{}
		user = userRepo.Signup(db, user)

		if err != nil {
			error.Message = "Server error."
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		user.Password = ""

		w.Header().Set("Content-Type", "application/json")
		utils.ResponseJSON(w, user)

		json.NewEncoder(w).Encode(user)
	}
}

// Show to display profile
func (c UserController) Show(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		params := mux.Vars(r)

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		id := params["id"]
		fmt.Printf(id)

		if id == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		row := db.QueryRow("SELECT * FROM users WHERE id=$1", id)

		newUser := user
		err := row.Scan(&newUser.ID, &newUser.Email, &newUser.Username, &newUser.Password)
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		} else if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	}
}
