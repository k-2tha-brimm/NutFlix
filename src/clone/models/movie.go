package models

// Movie model
type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}
