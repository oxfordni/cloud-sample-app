import cypress from "cypress";
import { Before, Given, Then, When } from "cypress-cucumber-preprocessor/steps";
import { any } from "cypress/types/bluebird";
import { home } from "../page_objects/home.pages";
import { utilities } from "../Utilities/utilities";


var oldQuote: string;
var newQuote: string;
var statusCode: string;
var serviceMessage: string = "";
var randomUserName: string;
var quoteId: string;
var quoteName: string;
var quoteRole: string;
var quoteShow: string;
var quoteLang: string;
var newUserURL: string;

Given(/^I navigate to URL$/, () => {
  cy.visit('');
});


When(/^I get the random quote$/, () => {
  home.getQuotes().then(($btn) => {
    oldQuote = $btn.text();
    cy.log(oldQuote);
  })
});

When(/^I refresh the page$/, () => {
  cy.reload()
});

When(/^I get the new random quote$/, () => {
  home.getQuotes().then(($btn) => {
    newQuote = $btn.text();
    cy.log(newQuote);
  })
});


Then(/^I verify random quotes are different$/, () => {
  expect(oldQuote).to.not.equal(newQuote)
}
);



Then(/^I verify the service status$/, () => {
  if (statusCode == '200') {
    home.getQuotes().then(($btn) => {
      newQuote = $btn.text();
      assert.notInclude(newQuote, "ERROR")
    })
  } else {
    home.getError().then(($btn) => {
      newQuote = $btn.text();
      assert.include(newQuote, "ERROR")
    })
  }
});


When(/^I get the service status of "([^"]*)"$/, (serviceName) => {
  serviceMessage = home.getStatus(statusCode);
  cy.log(serviceMessage)
});


When(/^I get the status code of "([^"]*)"$/, (serviceName) => {
  if (serviceName == "Go Service") {
    cy.intercept('/go/health').as('go')
    cy.visit('/')

    cy.wait('@go', { timeout: 60000 }).then((interception) => {
      statusCode = interception.response?.statusCode.toString()!
      cy.log(statusCode)
    })
  }
  if (serviceName == "Python Service") {

    cy.intercept('/py/health').as('py')
    cy.visit('/')

    cy.wait('@py', { timeout: 60000 }).then((interception) => {
      statusCode = interception.response?.statusCode.toString()!
      cy.log(statusCode)
    })
  }
});


When(/^I get list of users$/, () => {
  cy.request('GET', '/api/v1/users/').as('users').then((response) => {
    expect(response.body).to.not.be.null
  })
});


Then(/^I verify users$/, () => {

  cy.get('@users')
    .its('status')
    .should('equal', 200);

  cy.get('@users')
    .its('headers')
    .its('content-type')
    .should('include', 'application/json');

});


When(/^I get list of groups$/, () => {
  cy.request('GET', '/api/v1/groups/').as('groups').then((response) => {
    expect(response.body).to.not.be.null
  })
});

Then(/^I verify groups$/, () => {
  cy.get('@groups')
    .its('status')
    .should('equal', 200);

  cy.get('@groups')
    .its('headers')
    .its('content-type')
    .should('include', 'application/json');

});


When(/^I create a new user$/, () => {
  randomUserName = Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15);
  cy.log(randomUserName)
  var user = {
    "url": "http://localhost:3001/api/v1/users/1/",
    "username": "" + randomUserName + "",
    "email": "" + randomUserName + "@example.com",
    "groups": []
  }

  cy.request('POST', '/api/v1/users/', user).as('newUser').then((response) => {
    expect(response.body).to.not.be.null
    newUserURL = response.body.url
  })
});


Then(/^I verify user is created$/, () => {
  cy.get('@newUser')
    .its('status')
    .should('equal', 201);
});


When(/^I create a new group$/, () => {
  randomUserName = Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15);
  cy.log(randomUserName)
  var user = {
    "url": "http://localhost:3001/api/v1/groups/1/",
    "name": "" + randomUserName + ""
  }

  cy.request('POST', '/api/v1/groups/', user).as('newGroup').then((response) => {
    expect(response.body).to.not.be.null
  })
});

Then(/^I verify group is created$/, () => {
  cy.get('@newGroup')
    .its('status')
    .should('equal', 201);
});


When(/^I create a new quote$/, () => {
  randomUserName = Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15);
  cy.log(randomUserName)
  var payload = {
    "quote": "Ask Yourself, Who Writes The Books? Who Chooses What We Remember And What Gets Forgotten",
    "role": "" + randomUserName + "",
    "show": "Fargo S04",
    "contain_adult_lang": false
  }

  cy.request('POST', '/api/v1/quotes', payload).as('newQuote').then((response) => {
    expect(response.body).to.not.be.null
    quoteId = response.body.id
    quoteName = response.body.quote
    quoteRole = response.body.role
    quoteShow = response.body.show
    quoteLang = response.body.contain_adult_lang
  })
});

Then(/^I verify quote is created$/, () => {
  cy.get('@newQuote')
    .its('status')
    .should('equal', 200);
});

When(/^I get list of quote$/, () => {
  cy.request('GET', '/api/v1/quotes/' + quoteId).as('quotes').then((response) => {
    expect(response.body).to.not.be.null
  })
});

Then(/^I verify quote$/, () => {
  cy.get('@quotes')
    .its('status')
    .should('equal', 200);

  cy.get('@quotes')
    .its('headers')
    .its('content-type')
    .should('include', 'application/json');

  cy.get('@quotes')
    .its('body')
    .should('include', { id: quoteId })
    .should('include', { quote: quoteName })
    .should('include', { role: quoteRole })
    .should('include', { show: quoteShow })
    .should('include', { contain_adult_lang: quoteLang })
});

When(/^I update a quote$/, () => {
  randomUserName = Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15);
  cy.log(randomUserName)
  var payload = {
    "quote": "Ask Yourself, Who Writes The Books? Who Chooses What We Remember And What Gets Forgotten",
    "role": "" + randomUserName + "",
    "show": "Fargo S04",
    "contain_adult_lang": false
  }

  cy.request('PUT', '/api/v1/quotes/' + quoteId, payload).as('updatedQuote').then((response) => {
    expect(response.body).to.not.be.null
    quoteId = response.body.id
    quoteName = response.body.quote
    quoteRole = response.body.role
    quoteShow = response.body.show
    quoteLang = response.body.contain_adult_lang
  })
});

Then(/^I verify quote is updated$/, () => {
  cy.get('@updatedQuote')
    .its('status')
    .should('equal', 200);
});

When(/^I delete a quote$/, () => {
  cy.request('DELETE', '/api/v1/quotes/' + quoteId).as('deletedQuotes').then((response) => {
    expect(response.body).to.not.be.null
  })
});

Then(/^I verify quote is deleted$/, () => {
  cy.get('@deletedQuotes')
    .its('status')
    .should('equal', 204);
});







