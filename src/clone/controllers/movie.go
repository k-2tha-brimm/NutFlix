package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"fmt"

	"../models"
)

// Movie struce import
var movies []models.Movie

// MovieController struct
type MovieController struct{}

// Index will be used to display the users homepage of movies
func (c MovieController) Index(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		var movies = make([] models.Movie, 0)
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(movies)
	}
}