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
    this.arbWallet = null;
    this.contracts = {};
    return this.initWeb3();
  }

  async initWeb3() {
    // Modern dapp browsers...
    let web3ProviderArb = null;

    if (window.ethereum) {
      web3ProviderArb = window.ethereum;
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
      web3ProviderArb = window.web3.currentProvider;
    }
    // If no injected web3 instance is detected, fall back to Ganache
    else {
      web3ProviderArb = new ethers.providers.JsonRpcProvider(
        "http://localhost:7545"
      );
    }

    let web3Provider = new ethers.providers.JsonRpcProvider(
      "http://localhost:7545"
    );

    const contracts = require("../compiled.json");
    this.ethProvider = web3Provider;
    this.arbProvider = new ArbProvider(
      "http://localhost:1235",
      contracts,
      new ethers.providers.Web3Provider(web3ProviderArb)
    );
    return this.initContracts();
  }

  async initContracts() {
    var network = await this.ethProvider.getNetwork();
    var network2 = await this.arbProvider.getNetwork();

    console.log("eth networkId: " + network.chainId);
    console.log("arb networkId: " + network2.chainId);

    const testToken = require("../build/contracts/TestToken.json");

    let testTokenAddress =
      testToken.networks[network.chainId.toString()].address;
    this.testTokenAddres = testTokenAddress;

    console.log("eth tooken addresss: " + testTokenAddress);

    let ethTestTokenContractRaw = new ethers.Contract(
      testTokenAddress,
      testToken.abi,
      this.ethProvider
    );

    let ethWallet = await this.ethProvider.getSigner(0);
    this.contracts.EthTestToken = ethTestTokenContractRaw.connect(ethWallet);

    let arbWallet = await this.arbProvider.getSigner(0);
    this.arbWallet = arbWallet;

    let ethAddress = await ethWallet.getAddress();
    this.walletAddress = await arbWallet.getAddress();

    console.log("eth address " + ethAddress);
    console.log("arb address " + this.walletAddress);

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
      //console.log(this.ethProvider.getSigner(0).getAddress())
      let signer = this.ethProvider.getSigner(0);
      let address = await signer.getAddress();
      if (address != this.account) {
        this.account = address;
        this.render();
      }

      const arbsigner = await this.arbProvider.getSigner();
      const inboxManager = await arbsigner.globalInboxConn();
      inboxManager.on(
        "DepositERC20MessageDelivered",
        (dest, sender, contract, value) => {
          console.log(
            "deposit ERC20 triggered",
            dest,
            sender,
            contract,
            value,
            event
          );
          this.render();
        }
      );
    }, 200);
  }

  async render() {
    var content = $("#content");
    if (this.walletAddress) {
      $("#accountAddress").html(this.walletAddress);
      console.log(this.account);
      const ethBalance = await this.contracts.EthTestToken.balanceOf(
        this.walletAddress
      );
      $("#ethBalance").html("eth balance: " + ethBalance.toString());
      console.log(ethBalance);
      const arbBalance = await this.contracts.ArbTestToken.balanceOf(
        this.walletAddress
      );
      $("#arbBalance").html("arb balance: " + arbBalance.toString());
      console.log(arbBalance);
    } else {
      $("#accountAddress").html("Loading");
    }

    content.show();
  }

  async mint() {
    let val = parseInt($("#mintAmount").val());
    console.log("mint to " + this.walletAddress);
    const tx = await this.contracts.EthTestToken.mint(this.walletAddress, val);
    $("#mintForm").hide();
    $("#mintMessage").html("Tokens are minting...");
    $("#mintMessage").show();
    await tx.wait();
    console.log("minted");
    $("#mintMessage").hide();
    $("#mintForm").show();
    this.render();
  }

  async deposit() {
    let val = parseInt($("#depositAmount").val());
    let dest = parseInt($("#destination").val());
    const inboxManager = await this.arbProvider.globalInboxConn();
    const tx1 = await this.contracts.EthTestToken.approve(
      inboxManager.address,
      val
    );
    // $("#depositForm").hide();
    // $("#depositMessage").html("Approving transfer for deposit");
    // $("#depositMessage").show();
    // await tx1.wait();
    const signer = await this.arbProvider.getSigner();
    console.log();
    // Not yet implemented
    const tx2 = await signer.depositERC20(
      this.contracts.EthTestToken.address,
      this.walletAddress,
      val
    );
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
