# Testing Notes

I wrote my API tests in a Postman collection and added the exported JSON file to my repo in order to perform tests via CLI with Newman.

E2E testing was written in Cypress.

Created a .github/workflow folder with a yml file to run a github action on every push and pull request.

If I had more time, I would: - Fine-tune the github action, where I'm building/testing the application on every push & pull request. - Create more unit tests within the Go/Python backends (my experience is more geared toward JavaScript at this point, so taking a deep dive would take a bit of time.) - Create more validations within the E2E & API tests.

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
# Run the frontend tests while the app is running
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
