package config

// Config represents a GO+ES configuration
type Config struct {
	Server struct {
		Development bool   `default:"true" envconfig:"SERVER_DEVELOPMENT"`
		Port        int    `default:"3000" envconfig:"SERVER_PORT"`
	}
	App struct {
		Name        string `default:"go+es" envconfig:"APP_NAME"`
		MaxResults  int    `default:"10" envconfig:"APP_MAX_RESULTS"`
		ApiVersion  string `default:"v1" envconfig:"APP_API_VERSION"`
	}
	ElasticSearch struct {
		Server      string `default:"http://elasticsearch:9200" envconfig:"ELASTICSEARCH_SERVER"`
		IndexName   string `default:"quote" envconfig:"ELASTICSEARCH_INDEX_NAME"`
		Mapping     string `default:"{}" envconfig:"ELASTICSEARCH_MAPPING"`
		MaxRetries  int    `default:"5" envconfig:"ELASTICSEARCH_MAX_RETRIES"`
	}
	External struct {
		MovieQuotes string `default:"https://movie-quote-api.herokuapp.com/v1/quote/?format=json" envconfig:"EXTERNAL_MOVIE_QUOTES"`
	}
}

// ES Configuration
const (
	esMapping = `
{
	"settings": {
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings": {
		"properties": {
			"quote": {
				"type": "text",
				"store": true,
				"fielddata": true
			},
			"role": {
				"type": "keyword"
			},
			"show": {
				"type": "keyword"
			},
			"created": {
				"type": "date"
			},
			"suggest_field": {
				"type": "completion"
			}
		}
	}
}`
)
