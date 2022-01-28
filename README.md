# goes

It goes! (go+es)

## Development

If you have [Docker][docker] installed on your machine, starting up the app is
as simple as executing the following commands, after cloning this repo.

```sh
# Build and start the containers
docker compose up -d --build

# Stop the containers, when you're done
docker compose down
```

The app will be exposed on the port `3001`.

## Tests

Basic unit tests are available as a starting point.

### Frontend

```sh
# Run the tests while the app is running
docker exec -it frontend yarn test
```

### Backend

```sh
# Run the go tests while the app is running
docker exec -it backend-go go test -run="" ./pkg/...
```

```sh
# Run the Python tests while the app is running
docker exec -it backend-py python manage.py test
```

## ElasticSearch

The ElasticSearch Admin is available on port `8080` and the secret is `goes`.

## Documentation

More detailed documentation is available [here][docs], including the [API][api]
and the [Testing Technical Challenge][challenge].

<!-- References -->

[api]: ./docs/API.md
[challenge]: ./docs/TESTING_CHALLENGE.md
[docker]: https://www.docker.com/
[docs]: ./docs/README.md
