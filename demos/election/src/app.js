/* eslint-env browser */
"use strict";

var $ = require("jquery");
const ethers = require("ethers");
const ArbProvider = require("arb-provider-ethers");

require("bootstrap/dist/css/bootstrap.min.css");

let App = {
  provider: null,
  contracts: {},
  account: "0x0",
  hasVoted: false,

  init: function() {
    return App.initWeb3();
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
      web3Provider = new ethers.providers.JsonRpcProvider(
        "http://localhost:7545"
      );
    }

    const contracts = require("../compiled.json");
    App.provider = new ArbProvider(
      "http://localhost:1235",
      contracts,
      new ethers.providers.Web3Provider(web3Provider)
    );
    return App.initContract();
  },

  initContract: async function() {
    var network = await App.provider.getNetwork();
    const election = require("../build/contracts/Election.json");
    let address = election.networks[network.chainId.toString()].address;
    let electionContractRaw = new ethers.Contract(
      address,
      election.abi,
      App.provider
    );
    let wallet = await App.provider.getSigner(0);
    App.contracts.Election = electionContractRaw.connect(wallet); // eslint-disable-line require-atomic-updates
    App.account = await wallet.getAddress(); // eslint-disable-line require-atomic-updates
    App.listenForEvents();
    App.setupHooks();

    return App.render();
  },

  setupHooks: function() {
    $("#voteForm").submit(function(event) {
      App.castVote();
      event.preventDefault();
    });
  },

  // Listen for events emitted from the contract
  listenForEvents: function() {
    App.contracts.Election.on("votedEvent", (index, eventInfo) => {
      console.log("event triggered", event);
      // Reload when a new vote is recorded
      App.render();
    });

    var accountInterval = setInterval(function() {
      App.provider
        .getSigner(0)
        .then(wallet => {
          return wallet.getAddress();
        })
        .then(address => {
          if (address != App.account) {
            console.log("Updated account", address);
            App.account = address;
            App.render();
          }
        });
    }, 200);
  },

  render: async function() {
    var electionInstance;
    var loader = $("#loader");
    var content = $("#content");

    loader.show();
    content.hide();

    $("#accountAddress").html("Your Account: " + App.account);
    let candidatesCount = await App.contracts.Election.candidatesCount();
    var candidateFutures = [];
    for (
      var i = ethers.utils.bigNumberify(1);
      i.lte(candidatesCount);
      i = i.add(1)
    ) {
      candidateFutures.push(App.contracts.Election.candidates(i));
    }
    Promise.all(candidateFutures).then(candidates => {
      var candidatesResults = $("#candidatesResults");
      candidatesResults.empty();

      var candidatesSelect = $("#candidatesSelect");
      candidatesSelect.empty();
      for (var i = 0; i < candidates.length; i++) {
        var candidate = candidates[i];
        console.log("Candidate", i, "is", candidate);
        var id = candidate[0];
        var name = candidate[1];
        var voteCount = candidate[2];

        // Render candidate Result
        var candidateTemplate =
          "<tr><th>" +
          id +
          "</th><td>" +
          name +
          "</td><td>" +
          voteCount +
          "</td></tr>";
        candidatesResults.append(candidateTemplate);

        // Render candidate ballot option
        var candidateOption =
          "<option value='" + id + "' >" + name + "</ option>";
        candidatesSelect.append(candidateOption);
      }
    });

    let hasVoted = await App.contracts.Election.voters(App.account);
    console.log("hasVoted is", hasVoted);
    // Do not allow a user to vote
    if (hasVoted) {
      $("form").hide();
    } else {
      $("form").show();
    }
    loader.hide();
    content.show();
  },

  castVote: async function() {
    var candidateId = $("#candidatesSelect").val();
    await App.contracts.Election.vote(candidateId);
    // Wait for votes to update
    $("#content").hide();
    $("#loader").show();
  }
};

$(function() {
  $(window).on("load", function() {
    App.init();
  });
});
