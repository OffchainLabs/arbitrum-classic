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

import {RPCManager} from './vmManagerRPC'
import advanceToBlock from 'zeppelin-solidity/test/helpers/advanceToBlock';

const utils = require('ethereumjs-util');
const abi = require('ethereumjs-abi');

var ArbBalanceTracker = artifacts.require("ArbBalanceTracker");
var VMTracker = artifacts.require("VMTracker");
var ArbProtocol = artifacts.require("ArbProtocol");
var ArbValue = artifacts.require("ArbValue");
var VMTrackerLib = artifacts.require("VMTrackerLib");
var ChallengeManager = artifacts.require("ChallengeManager");

var rpcManager1;
var rpcManager2;

function createVm(vmTracker, cMan, vmId, config, vmState) {
  vmTracker.createVMHash(
    config.gracePeriod,
    config.escrowRequired,
    config.maxExecutionSteps,
    vmState,
    config.challengeVerifier,
    config.assertKeys
  ).then(function(createHash) {
    let sigs = config.assertKeys.map(key => utils.fromRpcSig(web3.eth.sign(key, createHash)));
    var sigVs = [];
    var sigRs = [];
    var sigSs = [];
    sigs.forEach(function(sig) {
      sigVs.push(sig.v);
      sigRs.push('0x' + sig.r.toString('hex'));
      sigSs.push('0x' + sig.s.toString('hex'));
    });
    console.log("Creating VM");
    return vmTracker.createVm(
      [
        vmId,
        vmState,
        createHash
      ],
      [
        config["gracePeriod"],
        config["maxExecutionSteps"],
        config["challengeVerifier"]
      ]
      , 
      config["escrowRequired"],
      sigVs,
      sigRs,
      sigSs,
      {from: config.assertKeys[0]}
    );
  });
}

let vmId = "0x223450";

contract('VMTracker', function(accounts) {
  it("test", async function() {
    let challengeManager = await ChallengeManager.deployed();
    let vmTracker = await VMTracker.deployed();

    await vmTracker.mintArbsToUser(accounts[1], 100000);
    await vmTracker.mintArbsToUser(accounts[2], 100000);

    let vmConfig = {
      "gracePeriod": 10,
      "escrowRequired": 50000,
      "assertKeys": [accounts[1], accounts[2]],
      "maxExecutionSteps":10000000,
      "challengeVerifier":0
    };

    rpcManager1 = new RPCManager(vmTracker, challengeManager, vmId, vmConfig, 0, false);
    await new Promise(resolve => setTimeout(resolve, 100));
    rpcManager2 = new RPCManager(vmTracker, challengeManager, vmId, vmConfig, 1, true);

    await new Promise(resolve => setTimeout(resolve, 1000));

    rpcManager1.getVMState(function(response) {
      createVm(vmTracker, challengeManager, vmId, vmConfig, response.result);
    });

    await new Promise(resolve => setTimeout(resolve, 1000));

    rpcManager1.disputableAssert(1000);

    // await new Promise(resolve => setTimeout(resolve, 1000));

    // await advanceToBlock(web3.eth.blockNumber + 14);

    // rpcManager2.disputableAssert(1000);

    // await new Promise(resolve => setTimeout(resolve, 1000));

    // await advanceToBlock(web3.eth.blockNumber + 14);

    // rpcManager1.disputableAssert(1000);

    // await new Promise(resolve => setTimeout(resolve, 1000));

    // await advanceToBlock(web3.eth.blockNumber + 14);

    await new Promise(resolve => setTimeout(resolve, 20000));
  });


});
