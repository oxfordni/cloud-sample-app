/// <reference types="cypress" />
// ***********************************************************
// This example plugins/index.js can be used to load plugins
//
// You can change the location of this file or turn off loading
// the plugins file with the 'pluginsFile' configuration option.
//
// You can read more here:
// https://on.cypress.io/plugins-guide
// ***********************************************************

// This function is called when a project is opened or re-opened (e.g. due to
// the project's config changing)

/**
 * @type {Cypress.PluginConfig}
 */
// eslint-disable-next-line no-unused-vars
module.exports = (on, config) => {
  // `on` is used to hook into various events Cypress emits
  // `config` is the resolved Cypress config
};

const cucumber = require("cypress-cucumber-preprocessor").default;
const browserify = require("@cypress/browserify-preprocessor");
const fs = require("fs-extra");
const path = require("path");
const resolve = require('resolve');
require('dotenv').config()

module.exports = (on, config) => {
  const options = {
    ...browserify.defaultOptions,
    typescript: resolve.sync('typescript', { baseDir: config.projectRoot }),
  };
  on("file:preprocessor", cucumber(options));
  require('cypress-mochawesome-reporter/plugin')(on);
  config.env.baseUrl = process.env.baseUrl
  return processConfig(on, config);
};

function processConfig(on, config) {
  const file = config.env.configFile;
  return getConfigurationByFile(file).then(function (file) {
    if (config.env.configFile === "development") {
      if (!process.env.URI_ROOT) {
        throw new Error(
          "URI_ROOT not set - export URI_ROOT=http://yourlocalhost.com"
        );
      }
      // append the URI_ROOT to the baseUrl
      file.baseUrl = file.baseUrl + process.env.URI_ROOT;
    }
    // always return the file object
    return file;
  });
}

function getConfigurationByFile(file) {
  const pathToConfigFile = path.resolve("cypress", "config", `${file}.json`);
  return fs.readJson(pathToConfigFile);
}
