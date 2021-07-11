package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var SERVER_PORT string
var WELCOME_MESSAGE string
var APP_TITLE string
var HOME_HTML string
var API_VERSION string
var MOVIE_QUOTES string

type Quote struct {
	Quote            string `json:"quote"`
	Role             string `json:"role"`
	Show             string `json:"show"`
	ContainAdultLang bool   `json:"contain_adult_lang"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	APP_TITLE = "go+es"
	WELCOME_MESSAGE = "Welcome to <span style=\"color: #d67936;\">" + APP_TITLE + "</span> !"
	HOME_HTML = `
<!doctype html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">

	<title>%s</title>
	<meta name="description" content="%s">
</head>

<body>
	<div style="display: flex; justify-content: center;">
		<h1>%s</h1>
	</div>
</body>
</html>
`

	fmt.Fprintf(w, HOME_HTML, APP_TITLE, APP_TITLE, WELCOME_MESSAGE)
}

func GetMovieQuote(w http.ResponseWriter, r *http.Request) {
	MOVIE_QUOTES = "https://movie-quote-api.herokuapp.com/v1/quote/?format=json"

	resp, err := http.Get(MOVIE_QUOTES)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var quote Quote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("[GET] quote")

	respondWithJSON(w, http.StatusOK, quote)
}

func CreateMovieQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[CREATE] quote ")
}

func ReadMovieQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	quoteID := vars["id"]

	fmt.Println("[READ] quote " + quoteID)
}

func UpdateMovieQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	quoteID := vars["id"]

	fmt.Println("[UPDATE] quote " + quoteID)
}

func DeleteMovieQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	quoteID := vars["id"]

	fmt.Println("[DELETE] quote " + quoteID)
}

func app() {
	API_VERSION = "v1"

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/api/"+API_VERSION+"/movie-quotes", GetMovieQuote).Methods("GET")
	router.HandleFunc("/api/"+API_VERSION+"/movie-quote", CreateMovieQuote).Methods("POST")
	router.HandleFunc("/api/"+API_VERSION+"/movie-quote/{id}", ReadMovieQuote).Methods("GET")
	router.HandleFunc("/api/"+API_VERSION+"/movie-quote/{id}", UpdateMovieQuote).Methods("PUT")
	router.HandleFunc("/api/"+API_VERSION+"/movie-quote/{id}", DeleteMovieQuote).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+SERVER_PORT, router))
}

func main() {
	SERVER_PORT = "3000"

	fmt.Println("Listening on port " + SERVER_PORT + "...")

	app()
}
