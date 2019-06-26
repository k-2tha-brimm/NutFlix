package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type User struct {
	username       string
	email          string
	passwordDigest string
	id             int
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "clone_development"
)

var db *sql.DB

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
		
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/users/show", UsersShow).Methods("GET")
	port := ":3000"

	fmt.Println("App is listening on port " + port)
	http.ListenAndServe(port, r)

	// defer db.Close()

	// sqlStatement := `
	// 	INSERT INTO users (username, password_digest, email)
	// 	VALUES ($1, $2, $3)
	// 	RETURNING id`

	// var id int

	// err = db.QueryRow(sqlStatement, "kevykev", "12345678", "kev.brimm@gmail.com").Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("New record ID is:", id)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../../index.html")
}

func UsersShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	row := db.QueryRow("SELECT * FROM users WHERE id = $1")

	user := new(User)
	err := row.Scan(&user.username, &user.email, &user.passwordDigest, &user.id)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "%s, %s, %d", user.username, user.email, user.id)

}
