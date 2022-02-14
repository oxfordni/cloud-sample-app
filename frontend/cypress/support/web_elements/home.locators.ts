class HomeLocators {

    get quote() {
        return cy.get('q')
    }
    get errorMessage() {
        return cy.get('.ant-alert-message')
    }

}

export default HomeLocators;