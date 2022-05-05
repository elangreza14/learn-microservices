package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, os.Getenv("POSTGRES_USER"))

}

func main() {

	http.HandleFunc("/", index)

	log.Println("application started at port :8000")
	log.Println(os.Getenv("POSTGRES_USER"))
	log.Println(os.Getenv("POSTGRES_PASSWORD"))

	log.Fatal(http.ListenAndServe(":8000", nil))

}
