package models

// User model
type User struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password string `json:"password"`
	ID             int    `json:"id"`
}
