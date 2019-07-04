package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"log"
	"./controllers"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
	
)

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

	r.HandleFunc("/users/show", UsersShow).Methods("GET")
	r.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	port := ":5000"

	fmt.Println("App is listening on port " + port)
	http.ListenAndServe(port, r)

	// defer db.Close()



	// var id int



	// fmt.Println("New record ID is:", id)

}



// UsersShow will be the users profile page
func UsersShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	id := r.FormValue("id")
	fmt.Printf(id)
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	row := db.QueryRow("SELECT * FROM users WHERE id=$1", id)

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
