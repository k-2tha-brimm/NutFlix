package utils

import (
	"encoding/json"
	"net/http"

	"../models"

	"github.com/gorilla/sessions"
)

// RespondWithError handles errors
func RespondWithError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

// ResponseJSON renders JSON
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

//Store the session data in the cookie
var Store = sessions.NewCookieStore([]byte("secret-password"))

// IsLoggedIn will check if the user has an active session and return bool
func IsLoggedIn(r *http.Request) bool {
	session, _ := Store.Get(r, "session")
	if session.Values["loggedin"] == "true" {
		return true
	}
	return false
}
