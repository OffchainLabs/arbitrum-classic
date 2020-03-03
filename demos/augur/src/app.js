/* eslint-env browser */
"use strict";

var $ = require("jquery");
const ethers = require("ethers");

require("bootstrap/dist/css/bootstrap.min.css");
require("bootstrap/js/dist/tab.js");
require("bootstrap/js/dist/alert.js");
require("bootstrap/js/dist/util.js");

const delay = ms => new Promise(res => setTimeout(res, ms));

class App {
  constructor() {
    this.provider = null;
    this.contracts = {};
    this.gld_units = null;
    return this.initWeb3();
  }

  async initWeb3() {
    // Modern dapp browsers...
    var standardProvider = null;
    if (window.ethereum) {
      standardProvider = window.ethereum;
      try {
        // Request account access if needed
        await window.ethereum.enable();
      } catch (error) {
        console.log("User denied account access");
      }
    } else if (window.web3) {
      // Legacy dapp browsers...
      standardProvider = window.web3.currentProvider;
    } else {
      // Non-dapp browsers...
      console.log(
        "Non-Ethereum browser detected. You should consider trying MetaMask!"
      );
    }

    this.provider = new ethers.providers.Web3Provider(standardProvider);
    return this.initContracts();
  }

  async initContracts() {
    var network = await this.provider.getNetwork();

    var contract_map = {};

    contract_map["Augur"] = require("../build/contracts/Augur.json");
    contract_map[
      "AugurTrading"
    ] = require("../build/contracts/AugurTrading.json");
    contract_map["GnosisSafe"] = require("../build/contracts/GnosisSafe.json");
    contract_map[
      "GnosisSafeRegistry"
    ] = require("../build/contracts/GnosisSafeRegistry.json");
    contract_map["WarpSync"] = require("../build/contracts/WarpSync.json");
    contract_map[
      "CreateOrder"
    ] = require("../build/contracts/CreateOrder.json");
    contract_map["ProfitLoss"] = require("../build/contracts/ProfitLoss.json");
    contract_map[
      "SimulateTrade"
    ] = require("../build/contracts/SimulateTrade.json");
    contract_map[
      "LegacyReputationToken"
    ] = require("../build/contracts/LegacyReputationToken.json");
    contract_map[
      "BuyParticipationTokens"
    ] = require("../build/contracts/BuyParticipationTokens.json");
    contract_map["ormulas"] = require("../build/contracts/Formulas.json");
    contract_map["HotLoading"] = require("../build/contracts/HotLoading.json");
    contract_map[
      "RedeemStake"
    ] = require("../build/contracts/RedeemStake.json");
    contract_map["RepSymbol"] = require("../build/contracts/RepSymbol.json");
    contract_map["ShareToken"] = require("../build/contracts/ShareToken.json");
    contract_map[
      "AffiliateValidator"
    ] = require("../build/contracts/AffiliateValidator.json");
    contract_map["Affiliates"] = require("../build/contracts/Affiliates.json");
    contract_map[
      "InitialReporter"
    ] = require("../build/contracts/InitialReporter.json");
    contract_map[
      "InitialReporterFactory"
    ] = require("../build/contracts/InitialReporterFactory.json");
    contract_map["Market"] = require("../build/contracts/Market.json");
    contract_map[
      "MarketFactory"
    ] = require("../build/contracts/MarketFactory.json");
    contract_map["OICash"] = require("../build/contracts/OICash.json");
    contract_map[
      "OICashFactory"
    ] = require("../build/contracts/OICashFactory.json");
    contract_map[
      "RepExchange"
    ] = require("../build/contracts/RepExchange.json");
    contract_map[
      "RepExchangeFactory"
    ] = require("../build/contracts/RepExchangeFactory.json");
    contract_map[
      "ReputationToken"
    ] = require("../build/contracts/ReputationToken.json");
    contract_map[
      "ReputationTokenFactory"
    ] = require("../build/contracts/ReputationTokenFactory.json");
    contract_map[
      "TestNetReputationToken"
    ] = require("../build/contracts/TestNetReputationToken.json");
    contract_map[
      "TestNetReputationTokenFactory"
    ] = require("../build/contracts/TestNetReputationTokenFactory.json");
    contract_map["Universe"] = require("../build/contracts/Universe.json");
    contract_map[
      "UniverseFactory"
    ] = require("../build/contracts/UniverseFactory.json");
    contract_map[
      "DisputeWindow"
    ] = require("../build/contracts/DisputeWindow.json");
    contract_map[
      "DisputeWindowFactory"
    ] = require("../build/contracts/DisputeWindowFactory.json");
    contract_map[
      "DisputeCrowdsourcer"
    ] = require("../build/contracts/DisputeCrowdsourcer.json");
    contract_map[
      "DisputeCrowdsourcerFactory"
    ] = require("../build/contracts/DisputeCrowdsourcerFactory.json");

    // let chainId = network.chainId.toString();
    const chainId = "123456789";
    console.log("chainId: " + chainId);

    const wallet = this.provider.getSigner(0);
    this.walletAddress = await wallet.getAddress();

    for (var contract in contract_map) {
      this.contracts[contract] = this.connectContract(
        contract,
        contract_map[contract],
        chainId
      );

      if (contract != "Augur" && contract != "AugurTrading") {
        console.log("connect: " + contract);
        this.registerToAugur(
          contract,
          contract_map[contract].networks[chainId].address
        );
      }
    }

    this.setupHooks();

    return this.render();
  }

  connectContract(contract, contractJson, chainId) {
    const contractAddress = contractJson.networks[chainId].address;
    console.log("contract Address: " + contractAddress);

    const contractRaw = new ethers.Contract(
      contractAddress,
      contractJson.abi,
      this.provider
    );

    const wallet = this.provider.getSigner(0);

    return contractRaw.connect(wallet);
  }

  async registerToAugur(contract, contractAddress) {
    let tx;
    try {
      tx = await this.contracts.Augur.registerContract(
        contract,
        contractAddress
      );
    } catch (e) {
      return this.handleFailure(e);
    }
    console.log(tx);
  }

  setupHooks() {
    $("#mintERC20Form").submit(event => {
      this.mintERC20();
      event.preventDefault();
    });
    $("#mintERC721Form").submit(event => {
      this.mintERC721();
      event.preventDefault();
    });
  }

  async render() {
    var content = $("#content");
    if (this.walletAddress) {
      $("#accountAddress").html(this.walletAddress);
    } else {
      $("#accountAddress").html("Loading");
    }

    if (this.contracts.Augur) {
      $("#nocontracts").hide();
      $("#minting").show();
      $("#augurAddress").html(this.contracts.Augur.address);
      $("#augurTradingAddress").html(this.contracts.AugurTrading.address);
    } else {
      $("#nocontracts").show();
      $("#minting").hide();
    }

    content.show();
  }

  alertError(element, alert_class, message) {
    $(element).removeClass("alert-primary alert-danger alert-success");
    $(element).addClass(alert_class);
    $(element + "-message").html(message);
    $(element).show();
  }

  alertSuccess(message) {
    this.alertError("#alert", "alert-success", message);
  }

  clearAlerts() {
    $("#ETH-alert").hide();
    $("#ERC20-alert").hide();
    $("#ERC721-alert").hide();
  }

  handleFailure(e) {
    let message;
    if (Object.prototype.hasOwnProperty.call(e, "reason")) {
      message = e.reason;
    } else if (
      Object.prototype.hasOwnProperty.call(e, "data") &&
      Object.prototype.hasOwnProperty.call(e, "message")
    ) {
      message = e.data.message;
    } else if (Object.prototype.hasOwnProperty.call(e, "message")) {
      message = e.message;
    } else {
      message = e.data;
    }

    $("#mintERC20Message").hide();
    $("#mintERC20Form").show();
    $("#mintERC721Message").hide();
    $("#mintERC721Form").show();

    this.alertError(
      "#alert",
      "alert-danger",
      "Failed making transaction: " + message
    );
    this.render();
  }

  // async mintERC20() {
  //   this.clearAlerts();
  //   let val = ethers.utils.parseUnits(
  //     $("#mintERC20Amount").val(),
  //     this.gld_units
  //   );
  //   $("#mintERC20Amount").val("");
  //   $("#mintERC20Form").hide();
  //   $("#mintERC20Message").html("Tokens are minting...");
  //   $("#mintERC20Message").show();
  //   let tx;
  //   try {
  //     tx = await this.contracts.TestToken.mint(this.walletAddress, val);
  //   } catch (e) {
  //     return this.handleFailure(e);
  //   }

  //   await tx.wait();
  //   $("#mintERC20Message").hide();
  //   $("#mintERC20Form").show();
  //   this.alertSuccess(
  //     "Successfully minted " +
  //       ethers.utils.formatUnits(val, this.gld_units) +
  //       " tokens"
  //   );
  //   this.render();
  // }

  // async mintERC721() {
  //   this.clearAlerts();
  //   let tokenId = parseInt($("#mintERC721Amount").val());
  //   $("#mintERC721Amount").val("");
  //   $("#mintERC721Form").hide();
  //   $("#mintERC721Message").html("Creating mint transaction");
  //   $("#mintERC721Message").show();
  //   let tx;
  //   try {
  //     tx = await this.contracts.TestItem.mintItem(
  //       this.walletAddress,
  //       tokenId
  //     );
  //   } catch (e) {
  //     return this.handleFailure(e);
  //   }
  //   $("#mintERC721Message").html("Token is minting");
  //   await tx.wait();
  //   $("#mintERC721Message").hide();
  //   $("#mintERC721Form").show();
  //   this.alertSuccess("Successfully minted token " + tokenId);
  //   this.render();
  // }
}

$(function() {
  $(window).on("load", () => {
    new App();
  });
});
