/* eslint-env browser */
"use strict";

var $ = require("jquery");
const ethers = require("ethers");
const ArbProvider = require("arb-provider-ethers").ArbProvider;
const ArbERC20 = require("arb-provider-ethers").ERC20;
const ArbERC721 = require("arb-provider-ethers").ERC721;

require("bootstrap/dist/css/bootstrap.min.css");

class App {
  constructor() {
    this.ethProvider = null;
    this.arbProvider = null;
    this.arbWallet = null;
    this.contracts = {};
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

    const contracts = require("../compiled.json");
    this.ethProvider = new ethers.providers.Web3Provider(standardProvider);
    this.arbProvider = new ArbProvider(
      "http://localhost:1235",
      contracts,
      new ethers.providers.Web3Provider(standardProvider)
    );
    return this.initContracts();
  }

  async initContracts() {
    var network = await this.ethProvider.getNetwork();
    var network2 = await this.arbProvider.getNetwork();

    console.log("eth networkId: " + network.chainId);
    console.log("arb networkId: " + network2.chainId);

    const testToken = require("../build/contracts/TestToken.json");
    const testItem = require("../build/contracts/TestItem.json");

    this.testItemAddress =
      testItem.networks[network.chainId.toString()].address;
    this.testTokenAddress =
      testToken.networks[network.chainId.toString()].address;

    console.log("testToken contract addresss: " + this.testTokenAddress);
    console.log("testItem contract addresss: " + this.testItemAddress);

    let ethTestTokenContractRaw = new ethers.Contract(
      this.testTokenAddress,
      testToken.abi,
      this.ethProvider
    );

    let arbTestTokenContractRaw = new ethers.Contract(
      this.testTokenAddress,
      ArbERC20.abi,
      this.arbProvider
    );

    let ethTestItemContractRaw = new ethers.Contract(
      this.testItemAddress,
      testItem.abi,
      this.ethProvider
    );

    let arbTestItemContractRaw = new ethers.Contract(
      this.testItemAddress,
      ArbERC721.abi,
      this.arbProvider
    );

    this.ethWallet = await this.ethProvider.getSigner(0);
    this.arbWallet = await this.arbProvider.getSigner(0);

    this.ethwalletAddress = await this.ethWallet.getAddress();
    this.arbwalletAddress = await this.arbWallet.getAddress();

    console.log("eth wallet address " + this.ethwalletAddress);
    console.log("arb wallet address " + this.arbwalletAddress);

    this.contracts.ArbTestToken = arbTestTokenContractRaw.connect(
      this.arbWallet
    );
    this.contracts.EthTestToken = ethTestTokenContractRaw.connect(
      this.ethWallet
    );

    this.contracts.ArbTestItem = arbTestItemContractRaw.connect(this.arbWallet);
    this.contracts.EthTestItem = ethTestItemContractRaw.connect(this.ethWallet);

    this.listenForEvents();
    this.setupHooks();
    return this.render();
  }

  setupHooks() {
    $("#mintForm").submit(event => {
      this.mint();
      event.preventDefault();
    });
    $("#depositForm").submit(event => {
      this.deposit();
      event.preventDefault();
    });
    $("#withdrawForm").submit(event => {
      this.withdraw();
      event.preventDefault();
    });
  }

  // Listen for events emitted from the contract
  async listenForEvents() {
    const inboxManager = await this.arbProvider.globalInboxConn();
    inboxManager.on(
      "ERC20DepositMessageDelivered",
      (vmid, sender, dest, contract, value) => {
        console.log(
          "deposit ERC20 triggered",
          "vmid address: " + vmid,
          "arb address: " + dest,
          "eth address: " + sender,
          "token address: " + contract,
          value,
          event
        );
        this.render();
      }
    );

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

      const ethBalance = await this.contracts.EthTestToken.balanceOf(
        this.ethwalletAddress
      );

      $("#ethBalance").html(ethBalance.toString());
      console.log("ethbalance: " + ethBalance);

      const vmId = await this.arbProvider.getVmID();
      const inboxManager = await this.arbProvider.globalInboxConn();
      const tx = await inboxManager.getTokenBalances(vmId);

      console.log("vm Balance in GolbalWallet: " + tx[1]);

      const txx = await inboxManager.getTokenBalances(this.ethwalletAddress);

      console.log("arbBalance in GolbalWallet: " + txx[1]);

      const arbBalance = await this.contracts.ArbTestToken.balanceOf(
        this.arbwalletAddress
      );

      $("#arbBalance").html(arbBalance.toString());

      console.log("arbBalance xx : " + arbBalance);

      $("#arbBalance").html(arbBalance.toString());
    } else {
      $("#accountAddress").html("Loading");
    }

    content.show();
  }

  async mint() {
    let val = parseInt($("#mintAmount").val());

    console.log("mint to " + this.ethwalletAddress);

    const tx = await this.contracts.EthTestToken.mint(
      this.ethwalletAddress,
      val
    );

    $("#mintForm").hide();
    $("#mintMessage").html("Tokens are minting...");
    $("#mintMessage").show();

    await tx.wait();

    console.log("minted");
    $("#mintMessage").hide();
    $("#mintForm").show();
    this.render();
  }

  async mintERC721() {
    let tokenId = parseInt($("#tokenId").val());

    console.log("mint to " + this.ethwalletAddress);

    const tx = await this.contracts.EthTestToken.mint(
      this.ethwalletAddress,
      tokenId
    );

    $("#mintForm").hide();
    $("#mintMessage").html("ERC 721 Tokens are minting...");
    $("#mintMessage").show();

    await tx.wait();

    console.log("minted");
    $("#mintMessage").hide();
    $("#mintForm").show();
    this.render();
  }

  async depositEth() {
    let ethDepositValue = parseInt($("#ethDepositValue").val());

    const tx = await this.arbWallet.depositEth(
      this.arbwalletAddress,
      ethDepositValue
    );
  }

  async depositERC721() {
    let tokenID = parseInt($("#depositId").val());
    const inboxManager = await this.arbProvider.globalInboxConn();
    const tx = await this.contracts.EthTestItem.approve(
      inboxManager.address,
      tokenID
    );

    $("#depositForm").hide();
    $("#depositMessage").html("Approving transfer to Arbitrum chain");
    $("#depositMessage").show();

    await tx.wait();
    console.log("approved erc721 deposit from: " + this.ethwalletAddress);

    const tx2 = await this.arbWallet.depositERC721(
      this.arbwalletAddress,
      this.contracts.EthTestItem.address,
      tokenID
    );

    $("#depositMessage").html("Depositing 721 token to Arbitrum chain");

    console.log("Waiting on tx");
    await tx2.wait(0);
    console.log("Finished waiting");

    $("#depositMessage").hide();
    $("#depositForm").show();

    this.render();
  }

  async deposit() {
    let val = parseInt($("#depositAmount").val());
    const inboxManager = await this.arbProvider.globalInboxConn();
    const tx1 = await this.contracts.EthTestToken.approve(
      inboxManager.address,
      val
    );

    $("#depositForm").hide();
    $("#depositMessage").html("Approving transfer to Arbitrum chain");
    $("#depositMessage").show();

    await tx1.wait();
    console.log("erc20 approved from: " + this.ethwalletAddress);

    const tx2 = await this.arbWallet.depositERC20(
      this.arbwalletAddress,
      this.contracts.EthTestToken.address,
      val
    );

    $("#depositMessage").html("Depositing 20 token to Arbitrum chain");

    console.log("Waiting on tx");
    await tx2.wait(0);
    console.log("Finished waiting");

    $("#depositMessage").hide();
    $("#depositForm").show();

    this.render();
  }

  async withdraw() {
    let val = parseInt($("#withdrawAmount").val());
    const tx = await this.contracts.ArbTestToken.withdraw(
      this.ethwalletAddress,
      val
    );
    $("#withdrawForm").hide();
    $("#withdrawMessage").html("Withdrawing from EthBridge");
    $("#withdrawMessage").show();
    await tx.wait();
    $("#withdrawMessage").hide();
    $("#withdrawForm").show();
  }
}

$(function() {
  $(window).on("load", () => {
    new App();
  });
});
