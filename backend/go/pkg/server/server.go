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
		Port        int
	}
	App struct {
		Name       string
		MaxResults int
		ApiVersion string
	}
	External struct {
		MovieQuotes string
	}
	es *db.ElasticSearch
}

func New(config *config.Config) Server {
	s := Server{}

	s.Server.Development = config.Server.Development
	s.Server.Port = config.Server.Port
	s.App.Name = config.App.Name
	s.App.MaxResults = config.App.MaxResults
	s.App.ApiVersion = config.App.ApiVersion
	s.External.MovieQuotes = config.External.MovieQuotes

	return s
}

func (s Server) Start(es *db.ElasticSearch) error {
	s.es = es

	router := mux.NewRouter().StrictSlash(true)

	apiPrefix := fmt.Sprintf("/api/%s", s.App.ApiVersion)
	quoteEndpoint := fmt.Sprintf("%s/movie-quote", apiPrefix)

	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/health", HealthHandler)
	router.HandleFunc(fmt.Sprintf("%s/movie-quotes", apiPrefix), s.GetMovieQuote).Methods("GET")
	router.HandleFunc(quoteEndpoint, s.CreateMovieQuote).Methods("POST")
	router.HandleFunc(quoteEndpoint, s.ReadMovieQuoteAll).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/{id}", quoteEndpoint), s.ReadMovieQuote).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/{id}", quoteEndpoint), s.UpdateMovieQuote).Methods("PUT")
	router.HandleFunc(fmt.Sprintf("%s/{id}", quoteEndpoint), s.DeleteMovieQuote).Methods("DELETE")

	log.Printf("Listening on port %d...\n", s.Server.Port)

	return http.ListenAndServe(fmt.Sprintf(":%d", s.Server.Port), router)
}
