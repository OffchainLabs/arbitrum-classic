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

    const augur = require("../build/contracts/Augur.json");
    const augurTrading = require("../build/contracts/AugurTrading.json");

    // let chainId = network.chainId.toString();
    const chainId = "123456789";

    console.log("chainId: " + chainId);

    const wallet = this.provider.getSigner(0);
    this.walletAddress = await wallet.getAddress();

    if (chainId in augur.networks) {
      const augurAddress = augur.networks[chainId].address;
      const augurTradingAddress = augurTrading.networks[chainId].address;
      console.log("augurAddress: " + augurAddress);
      console.log("augurTradingAddress: " + augurTradingAddress);
      // let testTokenAddress = testToken.networks[chainId].address;

      const augurContractRaw = new ethers.Contract(
        augurAddress,
        augur.abi,
        this.provider
      );

      const augurTradingContractRaw = new ethers.Contract(
        augurTradingAddress,
        augurTrading.abi,
        this.provider
      );

      this.contracts.Augur = augurContractRaw.connect(wallet);
      this.contracts.AugurTrading = augurTradingContractRaw.connect(wallet);

      this.setupHooks();
    }

    return this.render();
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
