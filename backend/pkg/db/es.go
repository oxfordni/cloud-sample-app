package db

import (
	"context"
	"log"
	"math"
	"time"

	"github.com/joaocarmo/goes/pkg/config"

	"github.com/olivere/elastic/v7"
)

type ElasticSearch struct {
	Config struct {
		IndexName string
		Mapping string
		MaxRetries int
		Server string
	}
	numRetries int
	context *context.Context
	client *elastic.Client
}

func New(config *config.Config) ElasticSearch {
	es := ElasticSearch{}

	es.Config.IndexName = config.ElasticSearch.IndexName
	es.Config.Mapping = config.ElasticSearch.Mapping
	es.Config.MaxRetries = config.ElasticSearch.MaxRetries
	es.Config.Server = config.ElasticSearch.Server

	es.numRetries = 0

	return es
}

func (es ElasticSearch) GetContext() *context.Context {
	return es.context
}

func (es ElasticSearch) GetClient() *elastic.Client {
	return es.client
}

func CreateClient(es ElasticSearch) (*elastic.Client, error) {
	log.Printf("[ElasticSearch] Connecting to %s\n", es.Config.Server)

	// Connect to the default Elasticsearch @ localhost:9200
	client, err := elastic.NewClient(
		elastic.SetURL(es.Config.Server),
		elastic.SetRetrier(elastic.NewBackoffRetrier(&elastic.ExponentialBackoff{})),
	)
	if err != nil {
		// In case the initial attempt fails, we retry a predefined number of times
		if es.numRetries >= es.Config.MaxRetries {
			return nil, err
		}

		es.numRetries++
		retryTime := int64(math.Pow(float64(2), float64(es.numRetries)))

		log.Printf("[ElasticSearch] Couldn't connect - %s\n", err)
		log.Printf("[ElasticSearch] Retrying in %ds...\n", retryTime)

		time.Sleep(time.Second * time.Duration(retryTime))

		return CreateClient(es)
	}

	return client, err
}

func (es ElasticSearch) Start() *ElasticSearch {
	ctx := context.Background()

	// Connect to the default Elasticsearch @ localhost:9200
	client, err := CreateClient(es)
	if err != nil {
		log.Fatalf("[ElasticSearch] Couldn't connect - %s\n", err)
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
