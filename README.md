# goes

It goes! (go+es)

## Development

```sh
# Start the containers
docker compose up -d

# Stop the containers
docker compose down -v
```

The app will be exposed on the port `3001`.

The ElasticSearch Admin is reachable on port `8080` and the secret is `goes`.

## API

Get a random movie / series quote from [F4R4N's movie-quote][f4r4n-movie-quote].

```txt
GET /api/v1/movie-quotes
```

CRUD for quotes.

```txt
POST /api/v1/movie-quote
GET /api/v1/movie-quote
GET /api/v1/movie-quote/{id}
PUT /api/v1/movie-quote/{id}
DELETE /api/v1/movie-quote/{id}
```

Payload signature.

```json
{
	"quote": "Ask Yourself, Who Writes The Books? Who Chooses What We Remember And What Gets Forgotten",
	"role": "Ethelrida Smutney",
	"show": "Fargo S04",
	"contain_adult_lang": false
}
```

<!-- References -->

[f4r4n-movie-quote]: https://github.com/F4R4N/movie-quote
