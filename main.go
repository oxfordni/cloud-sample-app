package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var SERVER_PORT string

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func app() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":"+SERVER_PORT, router))
}

func main() {
	SERVER_PORT = "3000"

	fmt.Println("Listening on port " + SERVER_PORT + "...")

	app()
}
