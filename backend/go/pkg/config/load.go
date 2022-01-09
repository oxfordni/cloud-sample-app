package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func getEsMapping() string {
	return esMapping
}

// Load the configuration from the environment
func Load() Config {
	// Load the configuration from a .env file
	_ = godotenv.Load()

	cfg := Config{}

	// Load the configuration from the environment
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	cfg.ElasticSearch.Mapping = getEsMapping()

	return cfg
}
