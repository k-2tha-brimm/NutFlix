package main

import (
	"database/sql"
	"fmt"
	"net/http"
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

	controller := controllers.Controller{}

	// r.HandleFunc("/users/show", controller.Show(db)).Methods("GET")
	r.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	r.HandleFunc("/login", controller.Login(db)).Methods("POST")
	port := ":5000"

	fmt.Println("App is listening on port " + port)
	http.ListenAndServe(port, r)
}
