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

    // let arbSysContractRaw = new ethers.Contract(
    //   "100",
    //   ArbSys.abi,
    //   this.arbProvider
    // );

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
    $("#depositETHform").submit(event => {
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

  async depositEth() {
    let ethDepositValue = ethers.utils.parseEther($("#ethDepositValue").val());
    const tx = await this.arbWallet.depositETH(
      this.arbwalletAddress,
      ethDepositValue
    );

    $("#depositETHform").hide();
    $("#depositETHMessage").html("Approving transfer to Arbitrum chain");
    $("#depositETHMessage").show();
    await tx.wait();
    $("#depositETHMessage").hide();
    $("#depositETHform").show();
  }

  async withdrawETH() {
    const ethWithdrawValue = ethers.utils.parseEther(
      $("#withdrawEthAmount").val()
    );
    const tx = await this.arbWallet.withdrawEthFromChain(ethWithdrawValue);
    $("#withdrawETHForm").hide();
    $("#withdrawETHMessage").html("Withdrawing from EthBridge");
    $("#withdrawETHMessage").show();
    await tx.wait();
    $("#withdrawETHMessage").hide();
    $("#withdrawETHForm").show();
    this.render();
  }

  async withdrawLockboxETH() {
    const inboxManager = await this.arbWallet.globalInboxConn();
    const tx = await inboxManager.withdrawEth();
    $("#withdrawLockboxETHForm").hide();
    $("#withdrawLockboxETHMessage").html("Withdrawing from lockbox");
    $("#withdrawLockboxETHMessage").show();
    await tx.wait();
    $("#withdrawLockboxETHMessage").hide();
    $("#withdrawLockboxETHForm").show();
    this.render();
  }

  async mintERC20() {
    let val = parseInt($("#mintERC20Amount").val());
    const tx = await this.contracts.EthTestToken.mint(
      this.ethwalletAddress,
      val
    );

    $("#mintERC20Form").hide();
    $("#mintERC20Message").html("Tokens are minting...");
    $("#mintERC20Message").show();
    await tx.wait();
    $("#mintERC20Message").hide();
    $("#mintERC20Form").show();
    this.render();
  }

  async depositERC20() {
    let val = parseInt($("#depositERC20Amount").val());
    const inboxManager = await this.arbProvider.globalInboxConn();
    const tx1 = await this.contracts.EthTestToken.approve(
      inboxManager.address,
      val
    );

    $("#depositERC20Form").hide();
    $("#depositERC20Message").html("Approving transfer to Arbitrum chain");
    $("#depositERC20Message").show();
    await tx1.wait();

    const tx2 = await this.arbWallet.depositERC20(
      this.arbwalletAddress,
      this.contracts.EthTestToken.address,
      val
    );
    $("#depositERC20Message").html("Depositing 20 token to Arbitrum chain");
    await tx2.wait(0);
    $("#depositERC20Message").hide();
    $("#depositERC20Form").show();
    this.render();
  }

  async withdrawERC20() {
    let val = parseInt($("#withdrawERC20Amount").val());
    const tx = await this.contracts.ArbTestToken.withdraw(
      this.ethwalletAddress,
      val
    );
    $("#withdrawERC20Form").hide();
    $("#withdrawERC20Message").html("Withdrawing from EthBridge");
    $("#withdrawERC20Message").show();
    await tx.wait();
    $("#withdrawERC20Message").hide();
    $("#withdrawERC20Form").show();
    this.render();
  }

  async withdrawLockboxERC20() {
    const inboxManager = await this.arbWallet.globalInboxConn();
    const tx = await inboxManager.withdrawERC20(
      this.contracts.EthTestToken.address
    );
    $("#withdrawLockboxERC20Form").hide();
    $("#withdrawLockboxERC20Message").html("Withdrawing from lockbox");
    $("#withdrawLockboxERC20Message").show();
    await tx.wait();
    $("#withdrawLockboxERC20Message").hide();
    $("#withdrawLockboxERC20Form").show();
    this.render();
  }

  async mintERC721() {
    let tokenId = parseInt($("#tokenId").val());
    const tx = await this.contracts.EthTestItem.mintItem(
      this.ethwalletAddress,
      tokenId
    );

    $("#mintERC721Form").hide();
    $("#mintERC721Message").html("ERC-721 Token is minting...");
    $("#mintERC721Message").show();
    await tx.wait();
    $("#mintERC721Message").hide();
    $("#mintERC721Form").show();
    this.render();
  }

  async depositERC721() {
    let tokenID = parseInt($("#depositERC721TokenId").val());
    const inboxManager = await this.arbProvider.globalInboxConn();
    const tx = await this.contracts.EthTestItem.approve(
      inboxManager.address,
      tokenID
    );

    $("#depositERC721Form").hide();
    $("#depositERC721Message").html("Approving transfer to Arbitrum chain");
    $("#depositERC721Message").show();
    await tx.wait();
    const tx2 = await this.arbWallet.depositERC721(
      this.arbwalletAddress,
      this.contracts.EthTestItem.address,
      tokenID
    );

    $("#depositERC721Message").html("Depositing 721 token to Arbitrum chain");
    await tx2.wait(0);
    $("#depositERC721Message").hide();
    $("#depositERC721Form").show();
    this.render();
  }

  async withdrawERC721() {
    let val = parseInt($("#withdrawERC721TokenId").val());
    const tx = await this.contracts.ArbTestItem.withdraw(
      this.ethwalletAddress,
      val
    );
    $("#withdrawERC721Form").hide();
    $("#withdrawERC721Message").html("Withdrawing from EthBridge");
    $("#withdrawERC721Message").show();
    await tx.wait();
    $("#withdrawERC721Message").hide();
    $("#withdrawERC721Form").show();

    this.render();
  }

  async withdrawLockboxERC721() {
    let val = parseInt($("#withdrawLockboxERC721TokenId").val());
    const inboxManager = await this.arbWallet.globalInboxConn();
    const tx = await inboxManager.withdrawERC721(
      this.contracts.EthTestItem.address,
      val
    );
    $("#withdrawLockboxERC721Form").hide();
    $("#withdrawLockboxERC721Message").html("Withdrawing from lockbox");
    $("#withdrawLockboxERC721Message").show();
    await tx.wait();
    $("#withdrawLockboxERC721Message").hide();
    $("#withdrawLockboxERC721Form").show();
    this.render();
  }
}

$(function() {
  $(window).on("load", () => {
    new App();
  });
});
