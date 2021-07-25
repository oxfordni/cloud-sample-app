package db

import (
	"context"
	"log"

	"github.com/joaocarmo/goes/pkg/config"

	"github.com/olivere/elastic/v7"
)

type ElasticSearch struct {
	Config struct {
		Server string
		IndexName string
		Mapping string
	}
	context *context.Context
	client *elastic.Client
}

func New(config *config.Config) ElasticSearch {
	es := ElasticSearch{}

	es.Config.Server = config.ElasticSearch.Server
	es.Config.IndexName = config.ElasticSearch.IndexName
	es.Config.Mapping = config.ElasticSearch.Mapping

	return es
}

func (es ElasticSearch) GetContext() *context.Context {
	return es.context
}

func (es ElasticSearch) GetClient() *elastic.Client {
	return es.client
}

func (es ElasticSearch) Start() *ElasticSearch {
	ctx := context.Background()

	// Connect to the default Elasticsearch @ localhost:9200
	client, err := elastic.NewClient(
		elastic.SetURL(es.Config.Server),
		elastic.SetRetrier(elastic.NewBackoffRetrier(&elastic.ExponentialBackoff{})),
	)
	if err != nil {
		log.Fatalln(err)
	}

	// Ping the Elasticsearch server
	info, code, err := client.Ping(es.Config.Server).Do(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("[ElasticSearch] Returned with code %d and version %s\n", code, info.Version.Number)

	// Check if the index exists
	exists, err := client.IndexExists(es.Config.IndexName).Do(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	if !exists {
		// Create the index
		createIndex, err := client.
			CreateIndex(es.Config.IndexName).
			BodyString(es.Config.Mapping).
			Do(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		if !createIndex.Acknowledged {
			log.Printf("[INDEX] %s not acknowledged\n", es.Config.IndexName)
		}
		log.Printf("[INDEX] %s created\n", es.Config.IndexName)
	}

	es.context = &ctx
	es.client = client

	return &es
}
