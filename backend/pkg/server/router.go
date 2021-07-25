package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/olivere/elastic/v7"
)

const (
	DEVELOPMENT = true
	DEFAULT_QUOTE_COUNT = 10
	MOVIE_QUOTES = "https://movie-quote-api.herokuapp.com/v1/quote/?format=json"
	ES_SERVER = "http://elasticsearch:9200"
	ES_INDEX_NAME = "quote"
)

type Quote struct {
	Id               string `json:"id,omitempty"`
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

func parseQuoteHits(esHits []*elastic.SearchHit) []Quote {
	var results []Quote

	for _, hit := range esHits {
		result := &Quote{}
		err := json.Unmarshal(hit.Source, result)
		if err != nil {
			log.Fatalln(err)
		}
		result.Id = hit.Id
		results = append(results, *result)
	}

	return results
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

	log.Println("[GET] quote")

	respondWithJSON(w, http.StatusOK, quote)
}

func CreateMovieQuote(ctx *context.Context, esClient *elastic.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var quote Quote
		_ = json.NewDecoder(r.Body).Decode(&quote)

		createdQuote, err := esClient.
			Index().
			Index(ES_INDEX_NAME).
			BodyJson(quote).
			Pretty(DEVELOPMENT).
			Do(*ctx)
		if err != nil {
			log.Fatalln(err)
		}

		quote.Id = createdQuote.Id
		log.Println("[CREATE] quote " + quote.Id)

		respondWithJSON(w, http.StatusOK, quote)

		// Flush to make sure the document was written
		_, err = esClient.Flush().Index(ES_INDEX_NAME).Do(*ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func ReadMovieQuoteAll(ctx *context.Context, esClient *elastic.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var quoteFrom int
		var quoteTo int

		if vars["from"] == "" {
			quoteFrom = 0
		} else {
			quoteFrom, _ = strconv.Atoi(vars["from"])
		}

		if vars["to"] == "" {
			quoteTo = DEFAULT_QUOTE_COUNT
		} else {
			quoteTo, _ = strconv.Atoi(vars["from"])
		}

		// Get all documents
		readQuotes, err := esClient.
			Search().
			Index(ES_INDEX_NAME).
			From(quoteFrom).Size(quoteTo).
			Pretty(DEVELOPMENT).
			Do(*ctx)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("[READ] quotes " + strconv.FormatInt(readQuotes.TotalHits(), 10))

		allQuotes := parseQuoteHits(readQuotes.Hits.Hits)

		respondWithJSON(w, http.StatusOK, allQuotes)
	}
}

func ReadMovieQuote(ctx *context.Context, esClient *elastic.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		quoteID := vars["id"]

		// Get the document with the specified ID
		readQuote, err := esClient.
			Get().
			Index(ES_INDEX_NAME).
			Id(quoteID).
			Pretty(DEVELOPMENT).
			Do(*ctx)
		if err != nil {
			log.Fatalln(err)
		}

		if readQuote.Found {
			log.Println("[READ] quote " + readQuote.Id)

			var quote Quote
			_ = json.Unmarshal(readQuote.Source, &quote)
			quote.Id = readQuote.Id

			respondWithJSON(w, http.StatusOK, quote)
		} else {
			log.Println("[READ] quote not found" + quoteID)
		}
	}
}

func UpdateMovieQuote(ctx *context.Context, esClient *elastic.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		quoteID := vars["id"]

		var quote Quote
		_ = json.NewDecoder(r.Body).Decode(&quote)

		esScript := elastic.
			NewScriptInline("ctx._source = params.newQuote").
			Lang("painless").
			Param("newQuote", quote)

		updatedQuote, err := esClient.
			Update().
			Index(ES_INDEX_NAME).
			Id(quoteID).
			Script(esScript).
			Pretty(DEVELOPMENT).
			Do(*ctx)
		if err != nil {
			log.Fatalln(err)
		}

		quote.Id = updatedQuote.Id
		log.Println("[UPDATE] quote " + updatedQuote.Id)

		respondWithJSON(w, http.StatusOK, quote)
	}
}

func DeleteMovieQuote(ctx *context.Context, esClient *elastic.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		quoteID := vars["id"]

		// Delete the document with the specified ID
		_, err := esClient.
			Delete().
			Index(ES_INDEX_NAME).
			Id(quoteID).
			Pretty(DEVELOPMENT).
			Do(*ctx)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("[DELETE] quote " + quoteID)
	}
}
