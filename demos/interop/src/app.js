/* eslint-env browser */
"use strict";

var $ = require("jquery");
const ethers = require("ethers");
const ArbProvider = require("arb-provider-ethers");

require("bootstrap/dist/css/bootstrap.min.css");

class App {
  constructor() {
    this.ethProvider = null;
    this.arbProvider = null;
    this.contracts = {};
    return this.initWeb3();
  }

  async initWeb3() {
    // Modern dapp browsers...
    let web3Provider = null;
    // if (window.ethereum) {
    //   console.log("Using metamask")
    //   web3Provider = new ethers.providers.Web3Provider(window.ethereum);
    //   try {
    //     // Request account access
    //     await window.ethereum.enable();
    //   } catch (error) {
    //     // User denied account access...
    //     console.error("User denied account access");
    //   }
    // }
    // // Legacy dapp browsers...
    // else if (window.web3) {
    //   web3Provider = new ethers.providers.Web3Provider(window.web3.currentProvider);
    // }
    // // If no injected web3 instance is detected, fall back to Ganache
    // else {
    //   web3Provider = new ethers.providers.JsonRpcProvider(
    //     "http://localhost:7546"
    //   );
    // }
    web3Provider = new ethers.providers.JsonRpcProvider(
      "http://localhost:7545"
    );

    const contracts = require("../compiled.json");
    this.ethProvider = web3Provider;
    this.arbProvider = new ArbProvider(
      "http://localhost:1235",
      contracts,
      new ethers.providers.Web3Provider(web3Provider)
    );
    return this.initContracts();
  }

  async initContracts() {
    var network = await this.ethProvider.getNetwork();
    const testToken = require("../build/contracts/TestToken.json");
    let testTokenAddress =
      testToken.networks[network.chainId.toString()].address;

    let ethTestTokenContractRaw = new ethers.Contract(
      testTokenAddress,
      testToken.abi,
      this.ethProvider
    );
    let ethWallet = this.ethProvider.getSigner(0);
    this.contracts.EthTestToken = ethTestTokenContractRaw.connect(ethWallet);

    let arbWallet = this.arbProvider.getSigner(0);
    let arbTestTokenContractRaw = new ethers.Contract(
      testTokenAddress,
      testToken.abi,
      this.arbProvider
    );
    this.contracts.ArbTestToken = arbTestTokenContractRaw.connect(arbWallet);

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
  listenForEvents() {
    var accountInterval = setInterval(async () => {
      // console.log(this.ethProvider.getSigner(0).getAddress())
      let signer = this.ethProvider.getSigner(0);
      let address = await signer.getAddress();
      if (address != this.account) {
        this.account = address;
        this.render();
      }
    }, 200);
  }

  async render() {
    if (this.account) {
      $("#accountAddress").html(this.account);
      const ethBalance = await this.contracts.EthTestToken.balanceOf(
        this.account
      );
      $("#ethBalance").html(ethBalance.toString());
      const arbBalance = await this.contracts.ArbTestToken.balanceOf(
        this.account
      );
      $("#arbBalance").html(arbBalance.toString());
    } else {
      $("#accountAddress").html("Loading");
    }
  }

  async mint() {
    let val = parseInt($("#mintAmount").val());
    const tx = await this.contracts.EthTestToken.mint(this.account, val);
    $("#mintForm").hide();
    $("#mintMessage").html("Tokens are minting...");
    $("#mintMessage").show();
    await tx.wait();
    $("#mintMessage").hide();
    $("#mintForm").show();
    this.render();
  }

  async deposit() {
    let val = parseInt($("#depositAmount").val());
    const vmAddress = await this.arbProvider.getVmID();
    const tx1 = this.contracts.EthTestToken.approve(vmAddress, val);
    $("#depositForm").hide();
    $("#depositMessage").html("Approving transfer for deposit");
    $("#depositMessage").show();
    await tx1.wait();
    const signer = await this.arbProvider.getSigner();
    // Not yet implemented
    const tx2 = signer.depositERC20(this.contracts.EthTestToken.address, val);
    $("#depositMessage").html("Depositing into EthBridge");
    await tx2.wait();
    $("#depositMessage").hide();
    $("#depositForm").show();
  }

  async withdraw() {
    let val = parseInt($("#depositAmount").val());
    const vmAddress = await this.arbProvider.getVmID();
    const signer = await this.arbProvider.getSigner();
    // Not yet implemented
    const tx = signer.withdrawERC20(this.contracts.EthTestToken.address, val);
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
