"use strict";

var $ = require("jquery");
const Web3 = require("web3");
const contract = require("truffle-contract");
const ArbProvider = require("arb-provider-web3");

import "bootstrap/dist/css/bootstrap.min.css";

let App = {
  web3Provider: null,
  contracts: {},
  account: "0x0",
  hasVoted: false,

  init: function() {
    return App.initWeb3();
  },

  initWeb3: async function() {
    // Modern dapp browsers...
    if (window.ethereum) {
      App.web3Provider = window.ethereum;
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
      App.web3Provider = window.web3.currentProvider;
    }
    // If no injected web3 instance is detected, fall back to Ganache
    else {
      App.web3Provider = new Web3.providers.HttpProvider(
        "http://localhost:7545"
      );
    }

    const contracts = require("../compiled.json");
    App.web3Provider = await ArbProvider(
      "http://localhost:1235",
      contracts,
      App.web3Provider
    );
    web3 = new Web3(App.web3Provider);

    return App.initContract();
  },

  initContract: function() {
    const election = require("../build/contracts/Election.json");
    // Instantiate a new truffle contract from the artifact
    App.contracts.Election = contract(election);
    // Connect provider to interact with contract
    App.contracts.Election.setProvider(App.web3Provider);

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
    App.contracts.Election.deployed().then(function(instance) {
      // Restart Chrome if you are unable to receive this event
      // This is a known issue with Metamask
      // https://github.com/MetaMask/metamask-extension/issues/2393
      instance
        .votedEvent(
          {},
          {
            fromBlock: 0,
            toBlock: "latest"
          }
        )
        .watch(function(error, event) {
          console.log("event triggered", event);
          // Reload when a new vote is recorded
          App.render();
        });
    });

    var accountInterval = setInterval(function() {
      web3.eth.getAccounts(function(err, accounts) {
        if (err === null && accounts[0] != App.account) {
          console.log("Updated account", accounts[0]);
          App.account = accounts[0];
          App.render();
        }
      });
    }, 100);
  },

  render: function() {
    var electionInstance;
    var loader = $("#loader");
    var content = $("#content");

    loader.show();
    content.hide();

    web3.eth.getAccounts(function(err, accounts) {
      console.log(err, accounts);
    });

    $("#accountAddress").html("Your Account: " + App.account);

    // Load contract data
    App.contracts.Election.deployed()
      .then(function(instance) {
        electionInstance = instance;
        return electionInstance.candidatesCount();
      })
      .then(function(candidatesCount) {
        console.log("Count is", candidatesCount.toString());

        var candidateFutures = [];
        for (
          var i = web3.toBigNumber(1);
          i.lte(candidatesCount);
          i = i.add(1)
        ) {
          candidateFutures.push(electionInstance.candidates(i));
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
        return electionInstance.voters(App.account);
      })
      .then(function(hasVoted) {
        console.log("hasVoted is", hasVoted);
        // Do not allow a user to vote
        if (hasVoted) {
          $("form").hide();
        }
        loader.hide();
        content.show();
      })
      .catch(function(error) {
        console.warn(error);
      });
  },

  castVote: function() {
    var candidateId = $("#candidatesSelect").val();
    App.contracts.Election.deployed()
      .then(function(instance) {
        return instance.vote(candidateId, { from: App.account });
      })
      .then(function(result) {
        // Wait for votes to update
        $("#content").hide();
        $("#loader").show();
      })
      .catch(function(err) {
        console.error(err);
      });
  }
};

$(function() {
  $(window).on("load", function() {
    App.init();
  });
});
