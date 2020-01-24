/* eslint-env browser */
"use strict";

var $ = require("jquery");
const ethers = require("ethers");
const ArbProvider = require("arb-provider-ethers").ArbProvider;
const ArbERC20 = require("arb-provider-ethers").ERC20;
const ArbERC721 = require("arb-provider-ethers").ERC721;

require("bootstrap/dist/css/bootstrap.min.css");

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
    // ETH
    $("#depositETHform").submit(event => {
      this.depositEth();
      event.preventDefault();
    });
    $("#withdrawETHForm").submit(event => {
      this.withdrawETH();
      event.preventDefault();
    });

    // ERC20
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

    // ERC721
    $("#mint721Form").submit(event => {
      this.mintERC721();
      event.preventDefault();
    });
    $("#deposit721Form").submit(event => {
      this.depositERC721();
      event.preventDefault();
    });
    $("#withdraw721Form").submit(event => {
      this.withdrawERC721();
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

      await this.renderETHInfo();
      await this.renderERC20Info();
      await this.renderERC721Info();
    } else {
      $("#accountAddress").html("Loading");
    }

    content.show();
  }

  getTokenBlance(tokenMap, tokenAddress) {
    for (var i = 0; i < tokenMap[0].length; i++) {
      if (tokenAddress == tokenMap[0][i].toString()) {
        return tokenMap[1][i];
      }
    }

    return 0;
  }

  async renderETHInfo() {
    const eth = await this.ethWallet.getBalance();
    $("#ethereumBalance").html(eth.toString());
    console.log("ETH balance: " + eth);

    const vmId = await this.arbProvider.getVmID();
    const inboxManager = await this.arbProvider.globalInboxConn();
    var tokenMap = await inboxManager.getTokenBalances(vmId);
    var balance = this.getTokenBlance(tokenMap, this.ethAddress);
    console.log("ETH vmBalance in GolbalWallet: " + balance);

    tokenMap = await inboxManager.getTokenBalances(this.ethwalletAddress);
    balance = this.getTokenBlance(tokenMap, this.ethAddress);
    console.log("ETH balance in GolbalWallet: " + balance);

    // const arbEth = await this.arbWallet.getBalance();
    // $("#arbEthBalance").html(arbEth.toString());
    // console.log("arbitrum ETH balance: " + arbEth);
  }

  async renderERC20Info() {
    const ethBalance = await this.contracts.EthTestToken.balanceOf(
      this.ethwalletAddress
    );
    $("#ethBalance").html(ethBalance.toString());
    console.log("ERC20 token ethbalance: " + ethBalance);

    const vmId = await this.arbProvider.getVmID();
    const inboxManager = await this.arbProvider.globalInboxConn();
    var tokenMap = await inboxManager.getTokenBalances(vmId);
    var balance = this.getTokenBlance(
      tokenMap,
      this.contracts.EthTestToken.address
    );
    console.log("ERC20 token vmBalance in GolbalWallet: " + balance);

    tokenMap = await inboxManager.getTokenBalances(this.ethwalletAddress);
    balance = this.getTokenBlance(
      tokenMap,
      this.contracts.EthTestToken.address
    );
    console.log("ERC20 token balance in GolbalWallet: " + balance);

    const arbBalance = await this.contracts.ArbTestToken.balanceOf(
      this.arbwalletAddress
    );
    $("#arbBalance").html(arbBalance.toString());
    console.log("ERC20 token arbBalance: " + arbBalance);
  }

  async renderERC721Info() {
    const ethBalance = await this.contracts.EthTestItem.balanceOf(
      this.ethwalletAddress
    );
    $("#eth721Balance").html(ethBalance.toString() + " number of tokens");
    console.log("Number of ERC721 tokens in ethBlanace: " + ethBalance);

    const vmId = await this.arbProvider.getVmID();
    const inboxManager = await this.arbProvider.globalInboxConn();
    var tokenMap = await inboxManager.getNFTTokens(vmId);
    console.log("VM ERC721 token list in GolbalWallet: " + tokenMap[1]);

    tokenMap = await inboxManager.getNFTTokens(this.ethwalletAddress);
    console.log("ERC721 token list in GolbalWallet: " + tokenMap[1]);

    const arbBalance = await this.contracts.ArbTestItem.balanceOf(
      this.arbwalletAddress
    );
    $("#arb721Balance").html(arbBalance.toString() + " number of tokens");
    console.log("Number of ERC721 tokens in arbBalance: " + arbBalance);
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

    const tx = await this.contracts.EthTestItem.mintItem(
      this.ethwalletAddress,
      tokenId
    );

    $("#mint721Form").hide();
    $("#mint721Message").html("ERC 721 Tokens are minting...");
    $("#mint721Message").show();

    await tx.wait();

    console.log("minted");
    $("#mint721Message").hide();
    $("#mint721Form").show();
    this.render();
  }

  async depositEth() {
    let ethDepositValue = parseInt($("#ethDepositValue").val());

    const tx = await this.arbWallet.depositETH(
      this.arbwalletAddress,
      ethDepositValue
    );

    $("#depositETHform").hide();
    $("#depositETHMessage").html("Approving transfer to Arbitrum chain");
    $("#depositETHMessage").show();

    await tx.wait();

    console.log("Finished waiting");
    $("#depositETHMessage").hide();
    $("#depositETHform").show();
  }

  async depositERC721() {
    let tokenID = parseInt($("#deposittokenId").val());
    console.log(tokenID);
    const inboxManager = await this.arbProvider.globalInboxConn();
    const tx = await this.contracts.EthTestItem.approve(
      inboxManager.address,
      tokenID
    );

    $("#deposit721Form").hide();
    $("#deposit721Message").html("Approving transfer to Arbitrum chain");
    $("#deposit721Message").show();

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
    $("#deposit721Message").hide();
    $("#deposit721Form").show();

    await delay(1000);
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

    await delay(1000);
    this.render();
  }

  async withdrawETH() {
    // let val = parseInt($("#withdrawEthAmount").val());
    // const tx = await this.contracts.ArbTestToken.withdraw(
    //   this.ethwalletAddress,
    //   val
    // );
    // $("#withdrawForm").hide();
    // $("#withdrawMessage").html("Withdrawing from EthBridge");
    // $("#withdrawMessage").show();
    // await tx.wait();
    // await delay(5000);
    // const tx2 = await this.arbWallet.withdrawERC20(this.contracts.EthTestToken.address);
    // $("#withdrawMessage").hide();
    // $("#withdrawForm").show();
    // this.render();
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

    await delay(5000);

    const tx2 = await this.arbWallet.withdrawERC20(
      this.contracts.EthTestToken.address
    );

    $("#withdrawMessage").hide();
    $("#withdrawForm").show();

    this.render();
  }

  async withdrawERC721() {
    let val = parseInt($("#withdrawtokenId").val());
    const tx = await this.contracts.ArbTestItem.withdraw(
      this.ethwalletAddress,
      val
    );
    $("#withdraw721Form").hide();
    $("#withdraw721Message").html("Withdrawing from EthBridge");
    $("#withdraw721Message").show();
    await tx.wait();

    await delay(15000);

    const tx2 = await this.arbWallet.withdrawERC721(
      this.contracts.EthTestItem.address,
      val
    );
    $("#withdraw721Message").hide();
    $("#withdraw721Form").show();

    this.render();
  }
}

$(function() {
  $(window).on("load", () => {
    new App();
  });
});
