# goes

It goes! (go+es)

## Development

```sh
# Build and start the containers
docker compose up -d --build

# Stop the containers
docker compose down
```

The app will be exposed on the port `3001`.

## ElasticSearch

The ElasticSearch Admin is available on port `8080` and the secret is `goes`.

## API

### Quotes

Check the `alive` status.

```txt
GET /go/health
```

Get a random movie / series quote from [F4R4N's movie-quote][f4r4n-movie-quote].

```txt
GET /api/v1/quotes/random
```

CRUD for quotes.

```txt
GET /api/v1/quotes
POST /api/v1/quotes
GET /api/v1/quotes/{id}
PUT /api/v1/quotes/{id}
DELETE /api/v1/quotes/{id}
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

The full API Specification is available on [localhost:3000/documentation/yaml][api-go-spec].

### Users & Groups

If required, the username and password for the `admin` user are:

```txt
username: admin
password: admin
```

Check the `alive` status.

```txt
GET /py/health
```

CRUD for users.

```txt
GET /api/v1/users/
POST /api/v1/users/
GET /api/v1/users/{id}/
PUT /api/v1/users/{id}/
DELETE /api/v1/users/{id}/
```

Payload signature.

```json
{
  "url": "http://localhost:3001/api/v1/users/1/",
  "username": "admin",
  "email": "admin@example.com",
  "groups": []
}
```

CRUD for groups.

```txt
GET /api/v1/groups/
POST /api/v1/groups/
GET /api/v1/groups/{id}/
PUT /api/v1/groups/{id}/
DELETE /api/v1/groups/{id}/
```

Payload signature.

```json
{
  "url": "http://localhost:3001/api/v1/groups/1/",
  "name": "pub"
}
```

The full API Specification is available on [localhost:8000/documentation][api-py-spec].

## Tests

Basic unit tests are available as starting point.

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

<!-- References -->

[api-go-spec]: http://localhost:3000/documentation/yaml
[api-py-spec]: http://localhost:8000/documentation
[f4r4n-movie-quote]: https://github.com/F4R4N/movie-quote
