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

import advanceToBlock from 'zeppelin-solidity/test/helpers/advanceToBlock';
import assertRevert from 'zeppelin-solidity/test/helpers/assertRevert';
const eutil = require('ethereumjs-util')

import {ArbVM} from './vmTrackerWrapper'
import {ArbManager} from './vmManager'

var VMTracker = artifacts.require("VMTracker");
var ArbProtocol = artifacts.require("ArbProtocol");
var ArbValue = artifacts.require("ArbValue");

let vmId = "0x123450";

let msg1 = {
  "data": '0x00',
  "amount": 0,
  "destination": "0x00"
};

let msg2 = {
  "data": '0x00',
  "amount": 0,
  "destination": "0x00"
};

let msg3 = {
  "data": "0x20",
  "amount": 1000000,
  "destination": "0x00"
};

let testAssertion1 =  {
  "afterHash": "0x3000000000000000000000000000000000000000000000000000000000000000",
  "numSteps": 100,
  "messages": [msg1, msg2]
};

let testAssertion2 =  {
  "afterHash": "0x2000000000000000000000000000000000000000000000000000000000000000",
  "numSteps": 100,
  "messages": [msg1, msg2]
};

let testAssertion3 =  {
  "afterHash": "0x2000000000000000000000000000000000000000000000000000000000000000",
  "numSteps": 1000000000000,
  "messages": [msg1, msg2]
};

let testAssertion4 =  {
  "afterHash": "0x2000000000000000000000000000000000000000000000000000000000000000",
  "numSteps": 100,
  "messages": [msg1, msg3]
};

var manager;
var manager2;
var vm;
var vm2;

contract('VMTracker', function(accounts) {
  before(async function() {
    let vmTracker = await VMTracker.deployed();
    let arbMachine = await ArbProtocol.deployed();
    let arbValue = await ArbValue.deployed();
    manager = new ArbManager(vmTracker, arbMachine, arbValue, accounts[0]);
    manager2 = new ArbManager(vmTracker, arbMachine, arbValue, accounts[2]);
  });
  it("it should be able to send a message from one manager to another", async function() {
    return manager.sendMessage(accounts[2], 0, '0x00');
  });
  it("it should fail to send a message from one manager to another with too much funds", async function() {
    return assertRevert(manager.sendMessage(accounts[2], 1000, '0x00'));
  });
  it("it should be able to create a VM", async function() {
    let vmStartState = "0x0000000000000000000000000000000000000000000000000000000000000000";
    vm = await manager.createDefaultVm(vmId, [accounts[0], accounts[1]], vmStartState);
    vm2 = await manager2.getVm(vmId);
    var vmState = await vm.getVmInfo();
    assert.equal(vmState["stateHash"], vmStartState);
  });
  it("it should fail a unanimous assertion with too few signatures", async function() {
    let precondition = await vm.generatePrecondition();
    let unanHash = await vm.unanimousAssertHash(precondition, testAssertion1);
    let signature0 = eutil.fromRpcSig(web3.eth.sign(accounts[0], unanHash));
    await assertRevert(vm.unanimousAssert(precondition, testAssertion1, [signature0]));
  });
  it("it should fail a unanimous assertion with wrong signatures", async function() {
    let precondition = await vm.generatePrecondition();
    let unanHash = await vm.unanimousAssertHash(precondition, testAssertion1);
    let signature0 = eutil.fromRpcSig(web3.eth.sign(accounts[0], unanHash));
    let signature1 = eutil.fromRpcSig(web3.eth.sign(accounts[2], unanHash));
    await assertRevert(vm.unanimousAssert(precondition, testAssertion1, [signature0, signature1]));
  });
  it("it should be able to make a unanimous assertion", async function() {
    let precondition = await vm.generatePrecondition();
    let unanHash = await vm.unanimousAssertHash(precondition, testAssertion1);
    let signature0 = eutil.fromRpcSig(web3.eth.sign(accounts[0], unanHash));
    let signature1 = eutil.fromRpcSig(web3.eth.sign(accounts[1], unanHash));
    return vm.unanimousAssert(precondition, testAssertion1, [signature0, signature1]);
  });
  it("it should fail to make a disputable assertion with invalid asserterNum", async function() {
    let precondition = await vm.generatePrecondition();
    await assertRevert(vm2.disputableAssert(precondition, testAssertion2));
  });
  it("it should fail to make a disputable assertion with incorrectAsserter", async function() {
    vm2.asserterNum = 0;
    let precondition = await vm.generatePrecondition();
    await assertRevert(vm2.disputableAssert(precondition, testAssertion2));
    vm2.asserterNum = -1;
  });
  it("it should fail to make a disputable assertion with too high balance", async function() {
    let precondition = await vm.generatePrecondition();
    precondition.balance += 100;
    await assertRevert(vm.disputableAssert(precondition, testAssertion2));
  });
  it("it should fail to make a disputable assertion with incorrect state hash", async function() {
    let precondition = await vm.generatePrecondition();
    precondition.stateHash = "0x7436546";
    await assertRevert(vm.disputableAssert(precondition, testAssertion2));
  });
  it("it should fail to make a disputable assertion with incorrect inbox hash", async function() {
    let precondition = await vm.generatePrecondition();
    precondition.inbox = "0x7436546";
    await assertRevert(vm.disputableAssert(precondition, testAssertion2));
  });
  it("it should fail to make a disputable assertion with incorrect before time", async function() {
    let precondition = await vm.generatePrecondition();
    precondition.beforeTime = 0;
    await assertRevert(vm.disputableAssert(precondition, testAssertion2));
  });
  it("it should fail to make a disputable assertion with incorrect after time", async function() {
    let precondition = await vm.generatePrecondition();
    precondition.afterTime = 1000000000;
    await assertRevert(vm.disputableAssert(precondition, testAssertion2));
  });
  it("it should fail to make a disputable assertion with too many steps", async function() {
    let precondition = await vm.generatePrecondition();
    precondition.afterTime = 1000000000;
    await assertRevert(vm.disputableAssert(precondition, testAssertion3));
  });
  it("it should fail to make a disputable assertion with too little escrow", async function() {
    let precondition = await vm.generatePrecondition();
    await assertRevert(vm._disputableAssert(precondition, testAssertion2, 100));
  });
  it("it should fail to make a disputable assertion with messages sending too much money", async function() {
    let precondition = await vm.generatePrecondition();
    await assertRevert(vm.disputableAssert(precondition, testAssertion4));
  });
  it("it should be able to make a disputable assertion with correct preconditions", async function() {
    let precondition = await vm.generatePrecondition();
    await vm.disputableAssert(precondition, testAssertion2);
  });
  it("it should fail to confirm the disputable assertion before deadline", async function() {
    await advanceToBlock(web3.eth.blockNumber + 5);
    await assertRevert(vm.confirmAsserted());
  });
  it("it should be able to confirm a disputable assertion after deadline", async function() {
    await advanceToBlock(web3.eth.blockNumber + 5);
    await vm.confirmAsserted();
    var vmState = await vm.getVmInfo();
    assert.equal(vmState["stateHash"], testAssertion2.afterHash);
  });
});
