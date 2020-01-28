/* eslint-env browser */
"use strict";

var $ = require("jquery");
const ethers = require("ethers");
const ArbProvider = require("arb-provider-ethers").ArbProvider;
const ArbERC20Factory = require("arb-provider-ethers/dist/lib/abi/ArbERC20Factory")
  .ArbERC20Factory;
const ArbERC721Factory = require("arb-provider-ethers/dist/lib/abi/ArbERC721Factory")
  .ArbERC721Factory;
const ArbSysFactory = require("arb-provider-ethers/dist/lib/abi/ArbSysFactory")
  .ArbSysFactory;

require("bootstrap/dist/css/bootstrap.min.css");
require("bootstrap/js/dist/tab.js");
require("bootstrap/js/dist/alert.js");
require("bootstrap/js/dist/util.js");

const delay = ms => new Promise(res => setTimeout(res, ms));

class App {
  constructor() {
    this.ethProvider = null;
    this.arbProvider = null;
    this.arbWallet = null;
    this.contracts = {};
    this.ethAddress = "0x0000000000000000000000000000000000000000";
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

    this.ethProvider = new ethers.providers.Web3Provider(standardProvider);
    this.arbProvider = new ArbProvider(
      "http://localhost:1235",
      new ethers.providers.Web3Provider(standardProvider)
    );
    return this.initContracts();
  }

  async initContracts() {
    var network = await this.ethProvider.getNetwork();

    const testToken = require("../build/contracts/TestToken.json");
    const testItem = require("../build/contracts/TestItem.json");

    this.testItemAddress =
      testItem.networks[network.chainId.toString()].address;
    this.testTokenAddress =
      testToken.networks[network.chainId.toString()].address;

    let ethTestTokenContractRaw = new ethers.Contract(
      this.testTokenAddress,
      testToken.abi,
      this.ethProvider
    );

    let arbTestTokenContractRaw = ArbERC20Factory.connect(
      this.testTokenAddress,
      this.arbProvider
    );

    let ethTestItemContractRaw = new ethers.Contract(
      this.testItemAddress,
      testItem.abi,
      this.ethProvider
    );

    let arbTestItemContractRaw = ArbERC721Factory.connect(
      this.testItemAddress,
      this.arbProvider
    );

    this.ethWallet = await this.ethProvider.getSigner(0);
    this.arbWallet = await this.arbProvider.getSigner(0);

    this.ethwalletAddress = await this.ethWallet.getAddress();
    this.arbwalletAddress = await this.arbWallet.getAddress();

    this.contracts.ArbTestToken = arbTestTokenContractRaw.connect(
      this.arbWallet
    );
    this.contracts.EthTestToken = ethTestTokenContractRaw.connect(
      this.ethWallet
    );

    // this.contracts.ArbSys = arbSysContractRaw.connect(this.arbWallet);
    this.contracts.ArbTestItem = arbTestItemContractRaw.connect(this.arbWallet);
    this.contracts.EthTestItem = ethTestItemContractRaw.connect(this.ethWallet);

    this.listenForEvents();
    this.setupHooks();
    return this.render();
  }

  setupHooks() {
    // ETH
    $("#depositETHForm").submit(event => {
      this.depositEth();
      event.preventDefault();
    });
    $("#withdrawETHForm").submit(event => {
      this.withdrawETH();
      event.preventDefault();
    });
    $("#withdrawLockboxETHForm").submit(event => {
      this.withdrawLockboxETH();
      event.preventDefault();
    });

    // ERC20
    $("#mintERC20Form").submit(event => {
      this.mintERC20();
      event.preventDefault();
    });
    $("#depositERC20Form").submit(event => {
      this.depositERC20();
      event.preventDefault();
    });
    $("#withdrawERC20Form").submit(event => {
      this.withdrawERC20();
      event.preventDefault();
    });
    $("#withdrawLockboxERC20Form").submit(event => {
      this.withdrawLockboxERC20();
      event.preventDefault();
    });

    // ERC721
    $("#mintERC721Form").submit(event => {
      this.mintERC721();
      event.preventDefault();
    });
    $("#depositERC721Form").submit(event => {
      this.depositERC721();
      event.preventDefault();
    });
    $("#withdrawERC721Form").submit(event => {
      this.withdrawERC721();
      event.preventDefault();
    });
    $("#withdrawLockboxERC721Form").submit(event => {
      this.withdrawLockboxERC721();
      event.preventDefault();
    });
  }

  // Listen for events emitted from the contract
  async listenForEvents() {
    const arbRollup = await this.arbProvider.arbRollupConn();
    arbRollup.on("ConfirmedAssertion", () => {
      this.render();
    });

    var accountInterval = setInterval(async () => {
      let address = await this.arbWallet.getAddress();

      if (address != this.account) {
        this.account = address;
        this.render();
      }
    }, 200);
  }

  async render() {
    var content = $("#content");
    if (this.ethwalletAddress) {
      $("#accountAddress").html(this.ethwalletAddress);
      try {
        await this.renderETHInfo();
      } catch (e) {
        console.log("Error rendering eth", e);
      }

      try {
        await this.renderERC20Info();
      } catch (e) {
        console.log("Error rendering erc-20", e);
      }

      try {
        await this.renderERC721Info();
      } catch (e) {
        console.log("Error rendering erc-721", e);
      }
    } else {
      $("#accountAddress").html("Loading");
    }

    content.show();
  }

  async renderETHInfo() {
    const vmId = await this.arbProvider.getVmID();
    const inboxManager = await this.arbProvider.globalInboxConn();

    const eth = await this.ethWallet.getBalance();
    $("#ethereumBalance").html(ethers.utils.formatEther(eth));

    const vmBalance = await inboxManager.getEthBalance(vmId);
    $("#vmEthBalance").html(ethers.utils.formatEther(vmBalance));

    const userWithdrawnBalance = await inboxManager.getEthBalance(
      this.ethwalletAddress
    );
    $("#withdrawnEthBalance").html(
      ethers.utils.formatEther(userWithdrawnBalance)
    );

    const arbEth = await this.arbProvider.getBalance(this.ethwalletAddress);
    $("#arbEthBalance").html(ethers.utils.formatEther(arbEth));
  }

  async renderERC20Info() {
    const vmId = await this.arbProvider.getVmID();
    const inboxManager = await this.arbProvider.globalInboxConn();

    const ethBalance = await this.contracts.EthTestToken.balanceOf(
      this.ethwalletAddress
    );
    $("#ethERC20Balance").html(ethBalance.toString());

    const vmBalance = await inboxManager.getERC20Balance(
      this.contracts.EthTestToken.address,
      vmId
    );
    $("#vmERC20Balance").html(vmBalance.toString());

    const withdrawnBalance = await inboxManager.getERC20Balance(
      this.contracts.EthTestToken.address,
      this.ethwalletAddress
    );
    $("#withdrawnERC20Balance").html(withdrawnBalance.toString());

    const arbBalance = await this.contracts.ArbTestToken.balanceOf(
      this.arbwalletAddress
    );
    $("#arbERC20Balance").html(arbBalance.toString());
  }

  async renderERC721Info() {
    const vmId = await this.arbProvider.getVmID();
    const inboxManager = await this.arbProvider.globalInboxConn();

    const ethBalance = await this.contracts.EthTestItem.tokensOfOwner(
      this.ethwalletAddress
    );
    $("#ethERC721Balance").html("[" + ethBalance.join(", ") + "]");

    var vmBalance = await inboxManager.getERC721Tokens(
      this.contracts.EthTestItem.address,
      vmId
    );
    $("#vmERC721Balance").html("[" + vmBalance.join(", ") + "]");

    var withdrawnBalance = await inboxManager.getERC721Tokens(
      this.contracts.EthTestItem.address,
      this.ethwalletAddress
    );
    $("#withdrawnERC721Balance").html("[" + withdrawnBalance.join(", ") + "]");

    const arbBalance = await this.contracts.ArbTestItem.tokensOfOwner(
      this.arbwalletAddress
    );
    $("#arbERC721Balance").html("[" + arbBalance.join(", ") + "]");
  }

  alertError(element, alert_class, message) {
    $(element).removeClass("alert-primary alert-danger alert-success");
    $(element).addClass(alert_class);
    $(element + "-message").html(message);
    $(element).show();
  }

  alertEthSuccess(message) {
    this.alertError("#ETH-alert", "alert-success", message);
  }

  alertERC20Success(message) {
    this.alertError("#ERC20-alert", "alert-success", message);
  }

  alertERC721Success(message) {
    this.alertError("#ERC721-alert", "alert-success", message);
  }

  clearAlerts() {
    $("#ETH-alert").hide();
    $("#ERC20-alert").hide();
    $("#ERC721-alert").hide();
  }

  handleFailureCommon(kind, e) {
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

    $("#deposit" + kind + "Message").hide();
    $("#deposit" + kind + "Form").show();
    $("#withdraw" + kind + "Message").hide();
    $("#withdraw" + kind + "Form").show();
    $("#withdrawLockbox" + kind + "Message").hide();
    $("#withdrawLockbox" + kind + "Form").show();
    this.alertError(
      "#" + kind + "-alert",
      "alert-danger",
      "Failed making transaction: " + message
    );
    this.render();
  }

  handleEthFailure(e) {
    this.handleFailureCommon("ETH", e);
  }

  handleERC20Failure(e) {
    $("#mintERC20Message").hide();
    $("#mintERC20Form").show();
    this.handleFailureCommon("ERC20", e);
  }

  handleERC721Failure(e) {
    $("#mintERC721Message").hide();
    $("#mintERC721Form").show();
    this.handleFailureCommon("ERC721", e);
  }

  async depositEth() {
    this.clearAlerts();
    let value = ethers.utils.parseEther($("#ethDepositValue").val());
    $("#ethDepositValue").val("");
    $("#depositETHForm").hide();
    $("#depositETHMessage").html("Creating deposit transaction");
    $("#depositETHMessage").show();
    let tx;
    try {
      tx = await this.arbWallet.depositETH(this.arbwalletAddress, value);
    } catch (e) {
      return this.handleEthFailure(e);
    }
    $("#depositETHMessage").html("Depositing into Arbitrum chain");

    await tx.wait();
    $("#depositETHMessage").hide();
    $("#depositETHForm").show();
    this.alertEthSuccess(
      "Successfully deposited " + ethers.utils.formatEther(value) + " ETH"
    );
    this.render();
  }

  async withdrawETH() {
    this.clearAlerts();
    const value = ethers.utils.parseEther($("#withdrawEthAmount").val());
    $("#withdrawEthAmount").val("");
    $("#withdrawETHForm").hide();
    $("#withdrawETHMessage").html("Creating withdrawal transaction");
    $("#withdrawETHMessage").show();
    let tx;
    try {
      tx = await this.arbWallet.withdrawEthFromChain(value);
    } catch (e) {
      return this.handleEthFailure(e);
    }
    $("#withdrawETHForm").hide();
    $("#withdrawETHMessage").html("Withdrawing from EthBridge");
    $("#withdrawETHMessage").show();

    try {
      await tx.wait();
    } catch (e) {
      return this.handleEthFailure(e);
    }
    $("#withdrawETHMessage").hide();
    $("#withdrawETHForm").show();
    this.alertEthSuccess(
      "Successfully withdrew " + ethers.utils.formatEther(value) + " ETH"
    );
    this.render();
  }

  async withdrawLockboxETH() {
    this.clearAlerts();
    $("#withdrawLockboxETHForm").hide();
    $("#withdrawLockboxETHMessage").html("Withdrawing from lockbox");
    $("#withdrawLockboxETHMessage").show();
    const inboxManager = await this.arbWallet.globalInboxConn();
    let tx;
    try {
      tx = await inboxManager.withdrawEth();
    } catch (e) {
      return this.handleEthFailure(e);
    }
    await tx.wait();
    $("#withdrawLockboxETHMessage").hide();
    $("#withdrawLockboxETHForm").show();
    this.alertEthSuccess("Successfully withdrew funds from lockbox");
    this.render();
  }

  async mintERC20() {
    this.clearAlerts();
    let val = parseInt($("#mintERC20Amount").val());
    $("#mintERC20Amount").val("");
    $("#mintERC20Form").hide();
    $("#mintERC20Message").html("Tokens are minting...");
    $("#mintERC20Message").show();
    let tx;
    try {
      tx = await this.contracts.EthTestToken.mint(this.ethwalletAddress, val);
    } catch (e) {
      return this.handleERC20Failure(e);
    }

    await tx.wait();
    $("#mintERC20Message").hide();
    $("#mintERC20Form").show();
    this.alertERC20Success("Successfully minted " + val + " tokens");
    this.render();
  }

  async depositERC20() {
    this.clearAlerts();
    const inboxManager = await this.arbProvider.globalInboxConn();
    let val = parseInt($("#depositERC20Amount").val());
    $("#depositERC20Amount").val("");
    $("#depositERC20Form").hide();
    $("#depositERC20Message").html("Creating approve transfer transaction");
    $("#depositERC20Message").show();
    let tx1;
    try {
      tx1 = await this.contracts.EthTestToken.approve(
        inboxManager.address,
        val
      );
    } catch (e) {
      return this.handleERC20Failure(e);
    }

    $("#depositERC20Message").html("Approving transfer to Arbitrum chain");
    await tx1.wait();

    $("#depositERC20Message").html("Creating deposit transaction");
    let tx2;
    try {
      tx2 = await this.arbWallet.depositERC20(
        this.arbwalletAddress,
        this.contracts.EthTestToken.address,
        val
      );
    } catch (e) {
      return this.handleERC20Failure(e);
    }
    $("#depositERC20Message").html("Depositing tokens in Arbitrum chain");

    await tx2.wait(0);

    $("#depositERC20Message").hide();
    $("#depositERC20Form").show();
    this.alertERC20Success("Successfully deposited " + val + " tokens");
    this.render();
  }

  async withdrawERC20() {
    this.clearAlerts();
    let val = parseInt($("#withdrawERC20Amount").val());
    $("#withdrawERC20Amount").val("");
    $("#withdrawERC20Form").hide();
    $("#withdrawERC20Message").html("Creating withdraw transaction");
    $("#withdrawERC20Message").show();
    let tx;
    try {
      tx = await this.contracts.ArbTestToken.withdraw(
        this.ethwalletAddress,
        val
      );
    } catch (e) {
      return this.handleERC20Failure(e);
    }
    $("#withdrawERC20Message").html("Withdrawing from EthBridge");
    try {
      await tx.wait();
    } catch (e) {
      return this.handleERC20Failure(e);
    }

    $("#withdrawERC20Message").hide();
    $("#withdrawERC20Form").show();
    this.alertERC20Success("Successfully withdrew " + val + " tokens");
    this.render();
  }

  async withdrawLockboxERC20() {
    this.clearAlerts();
    const inboxManager = await this.arbWallet.globalInboxConn();
    $("#withdrawLockboxERC20Form").hide();
    $("#withdrawLockboxERC20Message").html("Approving withdrawal from lockbox");
    $("#withdrawLockboxERC20Message").show();
    let tx;
    try {
      tx = await inboxManager.withdrawERC20(
        this.contracts.EthTestToken.address
      );
    } catch (e) {
      return this.handleERC20Failure(e);
    }
    $("#withdrawLockboxERC20Message").html("Withdrawing from lockbox");
    await tx.wait();
    $("#withdrawLockboxERC20Message").hide();
    $("#withdrawLockboxERC20Form").show();
    this.alertERC20Success("Successfully withdrew from lockbox");
    this.render();
  }

  async mintERC721() {
    this.clearAlerts();
    let tokenId = parseInt($("#mintERC721Amount").val());
    $("#mintERC721Amount").val("");
    $("#mintERC721Form").hide();
    $("#mintERC721Message").html("Creating mint transaction");
    $("#mintERC721Message").show();
    let tx;
    try {
      tx = await this.contracts.EthTestItem.mintItem(
        this.ethwalletAddress,
        tokenId
      );
    } catch (e) {
      return this.handleERC721Failure(e);
    }
    $("#mintERC721Message").html("Token is minting");
    await tx.wait();
    $("#mintERC721Message").hide();
    $("#mintERC721Form").show();
    this.alertERC721Success("Successfully minted token " + tokenId);
    this.render();
  }

  async depositERC721() {
    this.clearAlerts();
    let tokenId = parseInt($("#depositERC721TokenId").val());
    const inboxManager = await this.arbProvider.globalInboxConn();
    $("#depositERC721TokenId").val("");
    $("#depositERC721Form").hide();
    $("#depositERC721Message").html("Creating approving transaction");
    $("#depositERC721Message").show();
    let tx;
    try {
      tx = await this.contracts.EthTestItem.approve(
        inboxManager.address,
        tokenId
      );
    } catch (e) {
      return this.handleERC721Failure(e);
    }

    $("#depositERC721Message").html("Approving transfer to Arbitrum chain");
    await tx.wait();

    $("#depositERC721Message").html("Creating deposit transaction");
    let tx2;
    try {
      tx2 = await this.arbWallet.depositERC721(
        this.arbwalletAddress,
        this.contracts.EthTestItem.address,
        tokenId
      );
    } catch (e) {
      return this.handleERC721Failure(e);
    }

    $("#depositERC721Message").html("Depositing token to Arbitrum chain");
    await tx2.wait(0);

    $("#depositERC721Message").hide();
    $("#depositERC721Form").show();
    this.alertERC721Success("Deposited token " + tokenId);
    this.render();
  }

  async withdrawERC721() {
    this.clearAlerts();
    let tokenId = parseInt($("#withdrawERC721TokenId").val());
    $("#withdrawERC721TokenId").val("");
    $("#withdrawERC721Form").hide();
    $("#withdrawERC721Message").html("Creating withdraw transaction");
    $("#withdrawERC721Message").show();
    let tx;
    try {
      tx = await this.contracts.ArbTestItem.withdraw(
        this.ethwalletAddress,
        tokenId
      );
    } catch (e) {
      return this.handleERC721Failure(e);
    }

    $("#withdrawERC721Message").html("Withdrawing from EthBridge");
    try {
      await tx.wait();
    } catch (e) {
      return this.handleERC721Failure(e);
    }
    $("#withdrawERC721Message").hide();
    $("#withdrawERC721Form").show();
    this.alertERC721Success("Withdrew token " + tokenId);
    this.render();
  }

  async withdrawLockboxERC721() {
    this.clearAlerts();
    const inboxManager = await this.arbWallet.globalInboxConn();
    let tokenId = parseInt($("#withdrawLockboxERC721TokenId").val());
    $("#withdrawLockboxERC721Form").hide();
    $("#withdrawLockboxERC721Message").html("Creating withdraw transaction");
    $("#withdrawLockboxERC721Message").show();

    let tx;
    try {
      tx = await inboxManager.withdrawERC721(
        this.contracts.EthTestItem.address,
        tokenId
      );
    } catch (e) {
      return this.handleERC721Failure(e);
    }

    $("#withdrawLockboxERC721Message").html("Withdrawing from lockbox");
    await tx.wait();
    $("#withdrawLockboxERC721Message").hide();
    $("#withdrawLockboxERC721Form").show();
    this.alertERC721Success(
      "Successfully withdrew token " + tokenId + " from lockbox"
    );
    this.render();
  }
}

$(function() {
  $(window).on("load", () => {
    new App();
  });
});
