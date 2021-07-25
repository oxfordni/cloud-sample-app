package main

import (
	"log"

	"github.com/joaocarmo/goes/pkg/config"
	"github.com/joaocarmo/goes/pkg/db"
	"github.com/joaocarmo/goes/pkg/server"
)

func main() {
	appConfig := config.Load()

	log.Printf("Configuration loaded: %+v\n", appConfig)

	s := server.New(&appConfig)
	es := db.New(&appConfig)

	s.Start(es.Start())
}
