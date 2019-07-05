package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"fmt"

	"../models"
)

var movies []models.Movie

// MovieController struct
type MovieController struct{}

// Index will be used to display the users homepage of movies
func (c MovieController) Index(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies = make([]*Movie, 0)

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
			movie := new(Movie)
			if err := rows.Scan(&movie.Title, &movie.Genre, &movie.ID); err != nil {
				fmt.Println(err)
			}
			movies = append(movies, movie)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(movies)
	}
}