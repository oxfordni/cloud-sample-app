package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/joaocarmo/goes/pkg/config"
	"github.com/joaocarmo/goes/pkg/db"

	swagger "github.com/davidebianchi/gswagger"
	"github.com/davidebianchi/gswagger/apirouter"
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

	context := *s.es.GetContext()
	muxRouter := mux.NewRouter().StrictSlash(true)
	router, err := swagger.NewRouter(apirouter.NewGorillaMuxRouter(muxRouter), swagger.Options{
		Context: context,
		Openapi: &openapi3.T{
			Info: &openapi3.Info{
				Title:   "goes",
				Version: "1.0.0",
			},
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	apiPrefix := fmt.Sprintf("/api/%s", s.App.ApiVersion)
	quoteEndpoint := fmt.Sprintf("%s/movie-quote", apiPrefix)

	// Serve a basic index page
	muxRouter.HandleFunc("/", IndexHandler)
	// Health check
	router.AddRoute(http.MethodGet, "/health", HealthHandler, HealthHandlerDefinition)
	// API
	router.AddRoute(http.MethodGet, fmt.Sprintf("%s/movie-quotes", apiPrefix), s.GetMovieQuote, MovieQuotesHandlerDefinition)
	router.AddRoute(http.MethodPost, quoteEndpoint, s.CreateMovieQuote, PostQuoteHandlerDefinition)
	router.AddRoute(http.MethodGet, quoteEndpoint, s.ReadMovieQuoteAll, GetAllQuoteHandlerDefinition)
	router.AddRoute(http.MethodGet, fmt.Sprintf("%s/{id}", quoteEndpoint), s.ReadMovieQuote, GetQuoteHandlerDefinition)
	router.AddRoute(http.MethodPut, fmt.Sprintf("%s/{id}", quoteEndpoint), s.UpdateMovieQuote, PutQuoteHandlerDefinition)
	router.AddRoute(http.MethodDelete, fmt.Sprintf("%s/{id}", quoteEndpoint), s.DeleteMovieQuote, DeleteQuoteHandlerDefinition)

	err = router.GenerateAndExposeSwagger()

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Listening on port %d...\n", s.Server.Port)
	log.Printf("OpenAPI specification available at %s\n", swagger.DefaultYAMLDocumentationPath)

	return http.ListenAndServe(fmt.Sprintf(":%d", s.Server.Port), muxRouter)
}
