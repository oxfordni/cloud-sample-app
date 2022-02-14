class HomeIds {

    get menus() {
      return cy.get('.menu-tab-sub');
    }
    get create() {
      return cy.get(':nth-child(2) > .sidebar-item-link-basic');
    }

    get themes() {
      return cy.get('#login_theme');
    }
  }
  
  export default HomeIds;