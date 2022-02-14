Feature: e2e automated tests

    end to end tests

    Background: Navigate to URL
        Given I navigate to URL

    #@focus
    Scenario: Displays a random quote on refresh
        When I get the random quote
        When I refresh the page
        When I get the new random quote
        Then I verify random quotes are different

    Scenario: Displays the status of the backend services
        When I get the status code of "Go Service"
        When I get the service status of "Go Service"
        When I get the status code of "Python Service"
        When I get the service status of "Python Service"

