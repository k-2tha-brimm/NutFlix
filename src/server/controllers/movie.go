package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"../utils"
	"github.com/gorilla/mux"

	"../models"
)

// Movie struce import
var movies []models.Movie

// MovieController struct
type MovieController struct{}

// Index will be used to display the users homepage of movies
func (c MovieController) Index(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if utils.IsLoggedIn(r) {

			var movies = make([]models.Movie, 0)
			var movie models.Movie

			if r.Method != "GET" {
				http.Error(w, http.StatusText(405), 405)
				return
			}

			rows, err := db.Query("SELECT * FROM movies")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer rows.Close()

			for rows.Next() {
				newMovie := movie
				if err := rows.Scan(&newMovie.Title, &newMovie.Genre, &newMovie.ID); err != nil {
					fmt.Println(err)
				}
				movies = append(movies, newMovie)
			}

			enableCors(&w)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(movies)
		} else {
			http.Error(w, http.StatusText(403), 403)
		}
	}
}

// Show comment
func (c MovieController) Show(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if utils.IsLoggedIn(r) {
			var movie models.Movie
			params := mux.Vars(r)

			if r.Method != "GET" {
				http.Error(w, http.StatusText(405), 405)
				return
			}

			id := params["id"]

			if id == "" {
				http.Error(w, http.StatusText(400), 400)
				return
			}

			row := db.QueryRow("SELECT * FROM movies WHERE id=$1", id)

			newMovie := movie
			err := row.Scan(&newMovie.Title, &newMovie.Genre, &newMovie.ID)
			if err == sql.ErrNoRows {
				http.NotFound(w, r)
				return
			} else if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}

			enableCors(&w)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newMovie)
		} else {
			http.Error(w, http.StatusText(403), 403)
		}
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
