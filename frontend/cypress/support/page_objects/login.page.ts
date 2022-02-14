/// <reference types="cypress-xpath" />
import { loginIds } from "../web_elements/login.ids";

class LoginPage {

  public visit() {
  
    cy.visit('');
  }

  public signIn(user: string,password:string) {
    loginIds.userName.type(user);
    loginIds.password.type(password);
    loginIds.logIn.click();
  }

  public hoverOnElement(elementName : string)
  {
    cy.contains(elementName).realHover()
  }

}
export const login = new LoginPage();