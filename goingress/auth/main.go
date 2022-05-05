package main

import (
	db "auth/db/sqlc"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, os.Getenv("POSTGRES_USER"))
	// db ""
}

func main() {

	http.HandleFunc("/", index)

	log.Println("application started at port :8000")
	a := os.Getenv("POSTGRES_USER")
	b := os.Getenv("POSTGRES_PASSWORD")
	// a:= postgres://{POSTGRES_USER}:{POSTGRES_PASSWORD}@postgres:5432/books
	// postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
	stringdb := "postgres://" + a + ":" + b + "@postgres:5432/books"
	log.Println(stringdb)
	conn, err := sql.Open("postgres", stringdb)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	dball := db.New(conn)

	user, err := dball.CreateUser(context.Background(), db.CreateUserParams{Firstname: "aasas", Lastname: "testr"})
	if err != nil {
		log.Fatal("cannot Create user to db:", err)
	}

	log.Println(user)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
