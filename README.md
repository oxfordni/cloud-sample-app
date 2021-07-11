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

Random movie and series quotes.

```txt
GET /api/v1/movie-quotes
```
