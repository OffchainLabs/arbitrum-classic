/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License. 
 */

const utils = require('ethereumjs-util');
const abi = require('ethereumjs-abi');

import advanceToBlock from 'zeppelin-solidity/test/helpers/advanceToBlock';
import assertRevert from 'zeppelin-solidity/test/helpers/assertRevert';

import {ArbVM} from './vmTrackerWrapper'
import {ArbManager} from './vmManager'
import {RPCManager} from './vmManagerRPC'

var ArbBalanceTracker = artifacts.require("ArbBalanceTracker");
var VMTracker = artifacts.require("VMTracker");
var ArbProtocol = artifacts.require("ArbProtocol");
var ArbValue = artifacts.require("ArbValue");
var VMTrackerLib = artifacts.require("VMTrackerLib");
var ChallengeManager = artifacts.require("ChallengeManager");

let msg1 = {
  "data": "0x00",
  "tokenType":"0x000000000000000000000000000000000000000000",
  "amount": web3.toBigNumber(0),
  "destination": "0x01"
};

let msg2 = {
  "data": "0x00",
  "tokenType":"0x000000000000000000000000000000000000000000",
  "amount": web3.toBigNumber(0),
  "destination": "0x02"
};

let msg3 = {
  "data": "0x00",
  "tokenType":"0x000000000000000000000000000000000000000000",
  "amount": web3.toBigNumber(0),
  "destination": "0x03"
};

let vmId = "0x223450";
let testAssertion =  {
  "afterHash": "0x2000000000000000000000000000000000000000000000000000000000000000",
  "numSteps": 4,
  "messages": [msg1, msg2, msg3]
};

let testAssertion2 =  {
  "afterHash": "0x3000000000000000000000000000000000000000000000000000000000000000",
  "numSteps": 2,
  "messages": [msg1, msg2]
};

let testAssertion4 =  {
  "afterHash": "0x4000000000000000000000000000000000000000000000000000000000000000",
  "numSteps": 2,
  "messages": [msg3]
};

let testAssertion3 =  {
  "afterHash": "0x4000000000000000000000000000000000000000000000000000000000000000",
  "numSteps": 1,
  "messages": [msg1]
};

var challengeManager;

var manager1;
var manager2;
var vm;
var vm2;

var info;
var prec0;
var prec1;
var newPrec;

var vmConfig;
let vmStartState = "0x0000000000000000000000000000000000000000000000000000000000000000";

var rpcManager;

contract('VMTracker', function(accounts) {
  before(async function() {
    challengeManager = await ChallengeManager.deployed();
    let vmTracker = await VMTracker.deployed();
    let arbMachine = await ArbProtocol.deployed();
    let arbValue = await ArbValue.deployed();

    await vmTracker.mintArbsToUser(accounts[0], 100000);
    await vmTracker.mintArbsToUser(accounts[1], 100000);

    vmConfig = {
      "gracePeriod": 10,
      "escrowRequired": 50000,
      "assertKeys": [accounts[0], accounts[1]],
      "maxExecutionSteps":100000,
      "challengeVerifier":0
    };

    let rpcManager = new RPCManager(vmTracker, challengeManager, arbMachine, vmId, vmConfig, 1);

    manager1 = new ArbManager(vmTracker, arbMachine, arbValue, accounts[0]);
    manager2 = new ArbManager(vmTracker, arbMachine, arbValue, accounts[1]);
  });

  it("it should be able to create a vm", async function() {
    let vmTracker = await VMTracker.deployed();
    let cMan = await ChallengeManager.deployed();
    let createHash = await vmTracker.createVMHash(
      vmConfig.gracePeriod,
      vmConfig.escrowRequired,
      vmConfig.maxExecutionSteps,
      vmStartState,
      vmConfig.challengeVerifier,
      vmConfig.assertKeys
    );

    let signature0 = utils.fromRpcSig(web3.eth.sign(accounts[0], createHash));
    let signature1 = utils.fromRpcSig(web3.eth.sign(accounts[1], createHash));
    
    vm = await manager1.createVm(vmId, vmConfig, vmStartState, [signature0, signature1], cMan);
  });
  it("it should be able to get an existing vm ", async function() {
    let cMan = await ChallengeManager.deployed();
    vm2 = await manager2.getVm(vmId, challengeManager, [accounts[0], accounts[1]]);
  });
  it("it should be able to get vm info", async function() {
    info = await vm.getVmInfo();
  })

  // it("it should fail to initiate a challenge not during an assertion", async function() {
  //   await assertRevert(vm2.initiateChallenge());
  // });
  // it("it should fail to bisect a challenge not during an assertion", async function() {
  //   await assertRevert(vm._bisectAssertion(prec0, [testAssertion2, testAssertion4]));
  // });
  // it("it should fail to continue a challenge not during an assertion", async function() {
  //   await assertRevert(vm2.continueChallenge(0));
  // });

  it("it should be able to make a disputable assertion with correct preconditions", async function() {
    prec1 = await vm.getVmInfo();
    await vm.disputableAssert(info, testAssertion);
    await new Promise(resolve => setTimeout(resolve, 2000));
  });


  // it("it should fail to make a disputable assertion during another disputable assertion", async function() {
  //   let precondition = await vm2.getVmInfo();
  //   await assertRevert(vm2.disputableAssert(precondition, testAssertion));
  // });
  // // it("it should fail to bisect a challenge before a challenge is initiated", async function() {
  // //   await assertRevert(vm._bisectAssertion(prec1, [testAssertion2, testAssertion4]));
  // // });
  // // // it("it should fail to continue a challenge before a challenge is initiated", async function() {
  // // //   await assertRevert(vm2.continueChallenge(0));
  // // // });


  // it("it should be able to initiate a challenge", async function() {
  //   await vm2.initiateChallenge();
  // });


  // it("it should fail to make a disputable assertion after a challenge is initiated", async function() {
  //   let precondition = await vm2.getVmInfo();
  //   await assertRevert(vm2.disputableAssert(precondition, testAssertion));
  // });
  // // it("it should fail to initiate a challenge after a challenge is initiated", async function() {
  // //   await assertRevert(vm2.initiateChallenge());
  // // });
  // // // it("it should fail to continue a challenge before a bisection has occured", async function() {
  // // //   await assertRevert(vm2.continueChallenge(0));
  // // // });


  // it("it should be able to bisect a challenge", async function() {
  //   await vm.bisectAssertion(testAssertion2);
  //   await new Promise(resolve => setTimeout(resolve, 1000));
  // });

  // it("it should fail to make a disputable assertion after a bisection", async function() {
  //   let precondition = await vm2.getVmInfo();
  //   await assertRevert(vm2.disputableAssert(precondition, testAssertion));
  // });
  // // it("it should fail to initiate a challenge after a bisection", async function() {
  // //   await assertRevert(vm2.initiateChallenge());
  // // });
  // // it("it should fail to bisect a challenge after a bisection", async function() {
  // //   await assertRevert(vm._bisectAssertion(prec1, [testAssertion2, testAssertion4]));
  // //   await new Promise(resolve => setTimeout(resolve, 500));
  // // });


  // it("it should be able to continue a challenge", async function() {
  //   await vm2.continueChallenge(0);
  //   await new Promise(resolve => setTimeout(resolve, 500));
  // });


  // it("it should fail to make a disputable assertion after bisecting", async function() {
  //   let precondition = await vm2.getVmInfo();
  //   await assertRevert(vm2.disputableAssert(precondition, testAssertion));
  // });
  // // it("it should fail to initiate a challenge after bisecting", async function() {
  // //   await assertRevert(vm2.initiateChallenge());
  // // });
  // // // it("it should fail to continue a challenge before bisecting", async function() {
  // // //   await assertRevert(vm2.continueChallenge(0));
  // // // });


  // it("it should be able to bisect again", async function() {
  //   await vm.bisectAssertion(testAssertion3);
  //   await new Promise(resolve => setTimeout(resolve, 1000));
  // });
  // it("it should be able to continue again", async function() {
  //   await vm2.continueChallenge(0);
  // });
});
