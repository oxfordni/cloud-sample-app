Feature: api tests

    api tests

    Scenario: Create a new user
        When I create a new user
        Then I verify user is created

    Scenario: Get list of users
        When I get list of users
        Then I verify users

    Scenario: Create a new group
        When I create a new group
        Then I verify group is created

    Scenario: Get list of groups
        When I get list of groups
        Then I verify groups

    Scenario: Create a new quote
        When I create a new quote
        Then I verify quote is created

    Scenario: Get list of quotes
        When I get list of quote
        Then I verify quote
    Scenario:  Update a quote
        When I update a quote
        Then I verify quote is updated
        
    Scenario: Delete a quote
        When I delete a quote
        Then I verify quote is deleted