# Testing Technical Challenge

We built this amazing web app, but we only have a few very basic
[unit tests][unit].

To be successful, we need to write more tests so we can make sure our code is
working as expected. This will help us deliver a better product and our users
will be happy.

## Requirements

Feel free to use any language you like as well as any framework you are familiar
with for this challenge. Just make sure we can follow you, so please write clear
and concise code. Add comments wherever needed to help us understand what you
are doing.

Preferably, we would like you to write tests in a language and framework that is
already familar to the team. You can find our recommendations below.

### What we expect to see

For this challenge, we expect you to:

- Add some more unit tests to our go and python backends and the frontend
  - You can intercept/mock the network requests
- Write some API tests, e.g. test you can:
  - Create a new user
  - Create a new group
  - Create assign a group to a user
  - List users and groups
  - Create a new quote
  - List quotes
  - Update a quote
  - Delete a quote
- Write some automated end-to-end tests, e.g.
  - Test the UI displays a random quote on refresh
  - Test the UI displays the status of the backend services
    - `green` for alive
    - `red` for dead
- Implement a CI/CD pipeline for testing, e.g.
  - Tests run on every pull request (PR)
- Explain your approach to testing

## The tools we use and recommend

### Unit Tests

#### Frontend

Our frontend tests are written using the [Jest][jest] testing framework.

For testing the [React][react] components, we use the
[testing library][testing-library] framework. Even thou we haven't written any
component tests yet, the framework is already installed and set up.

#### Backend

Our Go backend tests written using the native [Go testing framework][go].

Our Python backend tests are written using the native [unittest][py] testing
framework.

### API Tests

We recommend using [Postman][postman] to test the API.

### End-to-end Tests

We recommend using [Cypress][cypress] to test the end-to-end flow of the app.

### CI/CD

You can use [GitHub Actions][gh-actions], [CircleCI][circleci], or other to
automate the testing process.

<!-- References -->

[circleci]: https://circleci.com/
[cypress]: https://www.cypress.io/
[go]: https://golang.org/pkg/testing/
[jest]: https://jestjs.io/
[postman]: https://www.getpostman.com/
[react]: https://reactjs.org/
[testing-library]: https://testing-library.com/
[unit]: ../README.md#tests
[unittest]: https://docs.python.org/3/library/unittest.html
[gh-actions]: https://docs.github.com/en/actions
