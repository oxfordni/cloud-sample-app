import HomeIds from "../web_elements/home.ids";
import HomeLocators from "../web_elements/home.locators";


class HomePage extends HomeLocators {

  public getQuotes() {
    return this.quote;
  }
  public getError() {
    return this.errorMessage;
  }
  public getStatus(statusCode: string): string {
    let serviceStatus = "";
    if (statusCode == '200') {
      cy.get('body').find('span').should('have.class', 'ant-typography-success')
      serviceStatus = "Service is Alive"
    } else {
      cy.get('body').find('span').should('have.class', 'ant-typography-danger')
      serviceStatus = "Service is Down"
    }
    return serviceStatus
  }

}
export const home = new HomePage();