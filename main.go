package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var SERVER_PORT string
var DATA []Data

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(DATA)
}

func app() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/all", returnAllArticles)

	log.Fatal(http.ListenAndServe(":"+SERVER_PORT, router))
}

func main() {
	SERVER_PORT = "3000"

	DATA = []Data{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}

	fmt.Println("Listening on port " + SERVER_PORT + "...")

	app()
}
