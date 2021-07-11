package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// "reflect"
	"time"

	"github.com/gorilla/mux"
	"github.com/olivere/elastic"
)

// Configuration
const SERVER_PORT = "3000"
const ES_SERVER = "http://elasticsearch:9200"
const APP_TITLE = "go+es"
const WELCOME_MESSAGE = "Welcome to <span style=\"color: #d67936;\">" + APP_TITLE + "</span> !"
const HOME_HTML = `
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
const API_VERSION = "v1"
const MOVIE_QUOTES = "https://movie-quote-api.herokuapp.com/v1/quote/?format=json"

type Quote struct {
	Quote            string `json:"quote"`
	Role             string `json:"role"`
	Show             string `json:"show"`
	ContainAdultLang bool   `json:"contain_adult_lang"`
}

type ESQuote struct {
	Quote
	Created time.Time             `json:"created,omitempty"`
	Suggest *elastic.SuggestField `json:"suggest_field,omitempty"`
}

const esMapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"tweet":{
			"properties":{
				"quote":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"role":{
					"type":"keyword"
				},
				"show":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, HOME_HTML, APP_TITLE, APP_TITLE, WELCOME_MESSAGE)
}

func GetMovieQuote(w http.ResponseWriter, r *http.Request) {
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

func es() {
	ctx := context.Background()

	// Connect to the default Elasticsearch @ localhost:9200
	client, err := elastic.NewClient(elastic.SetURL(ES_SERVER))
	if err != nil {
		log.Fatalln(err)
	}

	// Ping the Elasticsearch server
	info, code, err := client.Ping(ES_SERVER).Do(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("[ElasticSearch] Returned with code %d and version %s\n", code, info.Version.Number)
}

func app() {
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
	fmt.Println("Listening on port " + SERVER_PORT + "...")

	es()
	app()
}
