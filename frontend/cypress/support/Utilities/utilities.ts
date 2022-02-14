class Utilities {
  public amendURL(url: string, key: string, values: string): string {
    var newUrl = new URL(url as string);
    var valueArr = values.split(",");
    let urlAmend = "";

    switch (key) {
      case "all":
        newUrl.searchParams.set("utm_source", valueArr[0]);
        newUrl.searchParams.set("utm_campaign", valueArr[1]);
        newUrl.searchParams.set("utm_medium", valueArr[2]);
        urlAmend = newUrl.toString();
        break;
      case "utm_source":
        newUrl.searchParams.set("utm_source", valueArr[0]);
        let utm_source_words = /=value2|=value3/gi;
        urlAmend = newUrl.toString().replace(utm_source_words, "");
        break;
      case "utm_campaign":
        newUrl.searchParams.set("utm_campaign", valueArr[0]);
        let utm_campaign_words = /=value1|=value3/gi;
        urlAmend = newUrl.toString().replace(utm_campaign_words, "");
        break;
      case "utm_medium":
        newUrl.searchParams.set("utm_medium", valueArr[0]);
        let utm_medium_words = /=value1|=value2/gi;
        urlAmend = newUrl.toString().replace(utm_medium_words, "");
        break;

      default:
        let none_words = /=value1|=value2|=value3/gi;
        urlAmend = newUrl.toString().replace(none_words, "");
        break;
    }

    return urlAmend;
  }

  public isElementExist(element: string): boolean {
    let isExist: boolean = false;
    cy.wait(5000)
    cy.window({ timeout: 10000 }).then((win) => {
      const identifiedElement = win.document.querySelector(element)
      if (identifiedElement != null)
        isExist = true
    });
    return isExist
  }
}
export const utilities = new Utilities();
