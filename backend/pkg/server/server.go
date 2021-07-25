package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joaocarmo/goes/pkg/config"
	"github.com/joaocarmo/goes/pkg/db"

	"github.com/gorilla/mux"
)

type Server struct {
		Server struct {
			Development bool
			Port int
		}
		App struct {
			Name string
			MaxResults int
			ApiVersion string
		}
}

func New(config *config.Config) Server {
	s := Server{}

	s.Server.Development = config.Server.Development
	s.Server.Port = config.Server.Port
	s.App.Name = config.App.Name
	s.App.MaxResults = config.App.MaxResults
	s.App.ApiVersion = config.App.ApiVersion

	return s
}

func (s Server) Start(es *db.ElasticSearch) error {
	ctx := es.GetContext()
	esClient := es.GetClient()

	router := mux.NewRouter().StrictSlash(true)

	apiPrefix := fmt.Sprintf("/api/%s", s.App.ApiVersion)
	quoteEndpoint := fmt.Sprintf("%s/movie-quote", apiPrefix)

	router.HandleFunc("/", IndexHandler)
	router.HandleFunc(fmt.Sprintf("%s/movie-quotes", apiPrefix), GetMovieQuote).Methods("GET")
	router.HandleFunc(quoteEndpoint, CreateMovieQuote(ctx, esClient)).Methods("POST")
	router.HandleFunc(quoteEndpoint, ReadMovieQuoteAll(ctx, esClient)).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/{id}", quoteEndpoint), ReadMovieQuote(ctx, esClient)).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/{id}", quoteEndpoint), UpdateMovieQuote(ctx, esClient)).Methods("PUT")
	router.HandleFunc(fmt.Sprintf("%s/{id}", quoteEndpoint), DeleteMovieQuote(ctx, esClient)).Methods("DELETE")

	log.Printf("Listening on port %d...\n", s.Server.Port)

	return http.ListenAndServe(fmt.Sprintf(":%d", s.Server.Port), router)
}
