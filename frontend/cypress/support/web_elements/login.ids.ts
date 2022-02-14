class LoginIds {

  get userName() {
    return cy.get('#login_user');
  }
  get password() {
    return cy.get('#login_pass');
  }
  get logIn() {
    return cy.get('#login_button');
  }

}

export const loginIds = new LoginIds();