/// <reference types="cypress" />

describe('Server Health Display', () => {
  beforeEach(() => {
    cy.visit('http://localhost:3001')
  })

  it('displays green when python server alive', () => {
    cy.intercept('http://localhost:3001/py/health', {
      statusCode: 200,
      body: { alive: 'ok' },
    })
    cy.wait(3000)
    cy.contains('python service').then(($py) => {
      expect($py).to.have.css('color', 'rgb(82, 196, 26)')
    })
  })
  it('displays red when python server dead', () => {
    cy.intercept('http://localhost:3001/py/health', {
      statusCode: 504,
      body: '504 Gateway Timeout',
    })
    cy.wait(3000)
    cy.contains('python service').then(($py) => {
      expect($py).to.have.css('color', 'rgb(255, 77, 79)')
    })
  })
  it('displays green when go server alive', () => {
    cy.intercept('http://localhost:3001/go/health', {
      statusCode: 200,
      body: { alive: 'ok' },
    })
    cy.wait(3000)
    cy.contains('go service').then(($go) => {
      expect($go).to.have.css('color', 'rgb(82, 196, 26)')
    })
  })
  it('displays red when go server dead', () => {
    cy.intercept('http://localhost:3001/go/health', {
      statusCode: 504,
      body: '504 Gateway Timeout',
    })
    cy.wait(3000)
    cy.contains('go service').then(($go) => {
      expect($go).to.have.css('color', 'rgb(255, 77, 79)')
    })
  })
})
