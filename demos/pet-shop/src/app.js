/* eslint-env browser */
"use strict";

var $ = require("jquery");
const Web3 = require("web3");
const ArbProvider = require("arb-provider-web3");

require("bootstrap/dist/css/bootstrap.min.css");

let App = {
  web3: null,
  contracts: {},

  init: async function() {
    // Load pets.
    $.getJSON("pets.json", function(data) {
      var petsRow = $("#petsRow");
      var petTemplate = $("#petTemplate");

      for (let i = 0; i < data.length; i++) {
        petTemplate.find(".panel-title").text(data[i].name);
        petTemplate.find("img").attr("src", data[i].picture);
        petTemplate.find(".pet-breed").text(data[i].breed);
        petTemplate.find(".pet-age").text(data[i].age);
        petTemplate.find(".pet-location").text(data[i].location);
        petTemplate.find(".btn-adopt").attr("data-id", data[i].id);

        petsRow.append(petTemplate.html());
      }
    });

    return await App.initWeb3();
  },

  initWeb3: async function() {
    // Modern dapp browsers...
    let web3Provider = null;
    if (window.ethereum) {
      web3Provider = window.ethereum;
      try {
        // Request account access
        await window.ethereum.enable();
      } catch (error) {
        // User denied account access...
        console.error("User denied account access");
      }
    }
    // Legacy dapp browsers...
    else if (window.web3) {
      web3Provider = window.web3.currentProvider;
    }
    // If no injected web3 instance is detected, fall back to Ganache
    else {
      web3Provider = new Web3.providers.HttpProvider("http://localhost:7545");
    }

    const contracts = require("../compiled.json");

    let provider = await ArbProvider(
      "http://localhost:1235",
      contracts,
      web3Provider
    );
    App.web3 = new Web3(provider); // eslint-disable-line require-atomic-updates

    return App.initContract();
  },

  initContract: function() {
    let adoption = require("../build/contracts/Adoption.json");
    App.contracts.Adoption = new App.web3.eth.Contract(
      adoption.abi,
      adoption.networks["123456789"].address
    );

    // Use our contract to retrieve and mark the adopted pets
    App.markAdopted();

    return App.bindEvents();
  },

  bindEvents: function() {
    $(document).on("click", ".btn-adopt", App.handleAdopt);
  },

  markAdopted: async function() {
    try {
      let adopters = await App.contracts.Adoption.methods.getAdopters().call();
      for (let i = 0; i < adopters.length; i++) {
        if (adopters[i] !== "0x0000000000000000000000000000000000000000") {
          $(".panel-pet")
            .eq(i)
            .find("button")
            .text("Success")
            .attr("disabled", true);
        }
      }
    } catch (err) {
      console.log(err.message);
    }
  },

  handleAdopt: async function(event) {
    event.preventDefault();
    var petId = parseInt($(event.target).data("id"));

    try {
      let accounts = await App.web3.eth.getAccounts();
      await App.contracts.Adoption.methods
        .adopt(petId)
        .send({ from: accounts[0] });
      await App.markAdopted();
    } catch (err) {
      console.log("Error calling adopt", err.message);
    }
  }
};

$(function() {
  $(window).on("load", function() {
    App.init();
  });
});
