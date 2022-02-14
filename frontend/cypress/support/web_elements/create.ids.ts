class CreateIds {

  get firstName() {
    return cy.get('#DetailFormfirst_name-input');
  }
  get lastName() {
    return cy.get('#DetailFormlast_name-input');
  }
  get save() {
    return cy.get('#DetailForm_save-label');
  }
  get savedUserName() {
    return cy.get('#_form_header > h3');
  }
}

export default CreateIds;