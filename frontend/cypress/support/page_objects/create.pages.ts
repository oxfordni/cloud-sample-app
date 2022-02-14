/// <reference types="cypress-xpath" />

import CreateIds from "../web_elements/create.ids";


class CreatePage extends CreateIds {

  public createContact(): string {

    var first = Math.random().toString(36).substring(2);
    var last = Math.random().toString(36).substring(2);

    this.firstName.type(first, { force: true })
    this.lastName.type(last, { force: true })
    this.save.click({ force: true })
    return first + " " + last;
  }

  public userName(): any {
    return this.savedUserName;
  }
}
export const create = new CreatePage();