/// <reference types="cypress" />

describe('Server Health Display', () => {
  it('displays green when python server alive', () => {
    cy.intercept('http://localhost:3001/py/health', {
      statusCode: 200,
      body: { alive: 'ok' },
    }).as('pyhealth')
    cy.visit('http://localhost:3001')
    cy.wait('@pyhealth')
    cy.wait(3000)
    cy.contains('python service').then(($py) => {
      expect($py).to.have.css('color', 'rgb(82, 196, 26)')
    })
  })
  it('displays red when python server dead', () => {
    cy.intercept('http://localhost:3001/py/health', {
      statusCode: 504,
      body: '504 Gateway Timeout',
    }).as('pyhealth')
    cy.visit('http://localhost:3001')
    cy.wait('@pyhealth')
    cy.wait(3000)
    cy.contains('python service').then(($py) => {
      expect($py).to.have.css('color', 'rgb(255, 77, 79)')
    })
  })
  it('displays green when go server alive', () => {
    cy.intercept('http://localhost:3001/go/health', {
      statusCode: 200,
      body: { alive: 'ok' },
    }).as('gohealth')
    cy.visit('http://localhost:3001')
    cy.wait('@gohealth')
    cy.wait(3000)
    cy.contains('go service').then(($go) => {
      expect($go).to.have.css('color', 'rgb(82, 196, 26)')
    })
  })
  it('displays red when go server dead', () => {
    cy.intercept('http://localhost:3001/go/health', {
      statusCode: 504,
      body: '504 Gateway Timeout',
    }).as('gohealth')
    cy.visit('http://localhost:3001')
    cy.wait('@gohealth')
    cy.wait(3000)
    cy.contains('go service').then(($go) => {
      expect($go).to.have.css('color', 'rgb(255, 77, 79)')
    })
  })
})
