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

    // let testTokenAddress2 =
    //   testToken.networks[network2.chainId.toString()].address;
    // this.testTokenAddres = testTokenAddress;

    console.log("eth token contract addresss: " + testTokenAddress);

    let ethTestTokenContractRaw = new ethers.Contract(
      testTokenAddress,
      testToken.abi,
      this.ethProvider
    );

    let arbTestTokenContractRaw = new ethers.Contract(
      testTokenAddress,
      testToken.abi,
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

      console.log("vmId address: " + vmId);
      console.log("arbBalance in GolbalWallet: " + tx[1]);

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

  async deposit() {
    let val = parseInt($("#depositAmount").val());

    const inboxManager = await this.arbProvider.globalInboxConn();
    const tx1 = await this.contracts.EthTestToken.approve(
      inboxManager.address,
      val
    );
    console.log("approved from : " + this.ethwalletAddress);

    const tx2 = await this.arbWallet.depositERC20(
      this.contracts.EthTestToken.address,
      this.arbwalletAddress,
      val
    );

    $("#depositMessage").html("Depositing into EthBridge");
    $("#depositMessage").hide();
    $("#depositForm").show();

    this.render();
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
