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
import merkleTree from 'zeppelin-solidity/test/helpers/merkleTree';

import {ArbInt, ArbNone} from './vm-arb-value';
import {Stack, Machine} from './arb/vm-machine';
import {ArbVM, generateAssertionStub} from './vmTrackerWrapper'
import {ArbManager} from './vmManager'
import generateProof from './arb/vm-proof';

const utils = require('ethereumjs-util');
const abi = require('ethereumjs-abi');

var VMTracker = artifacts.require("VMTracker");
var ArbProtocol = artifacts.require("ArbProtocol");
var ArbValue = artifacts.require("ArbValue");
var OneStepProof = artifacts.require("OneStepProof");
var ChallengeManager = artifacts.require("ChallengeManager");

let s1n = new Machine(
	new Stack([new ArbInt(25), new ArbInt(25), new ArbInt(25), new ArbInt(25), new ArbInt(10)]), 
	new Stack([new ArbInt(1), new ArbInt(1), new ArbInt(1), new ArbInt(1), new ArbInt(1), new ArbInt(10)]), 
	new Stack([]),
	new ArbNone(),
	new ArbNone()
);

let s0 = new Machine(
	new Stack([new ArbInt(25), new ArbInt(25), new ArbInt(25), new ArbInt(25)]), 
	new Stack([new ArbInt(1), new ArbInt(1), new ArbInt(1), new ArbInt(1), new ArbInt(1)]), 
	new Stack([]),
	new ArbNone(),
	new ArbNone()
);

let s1 = new Machine(
	new Stack([new ArbInt(25), new ArbInt(25), new ArbInt(25)]), 
	new Stack([new ArbInt(1), new ArbInt(1), new ArbInt(1), new ArbInt(2)]), 
	new Stack([]),
	new ArbNone(),
	new ArbNone()
);

let s2 = new Machine(
	new Stack([new ArbInt(25), new ArbInt(25)]), 
	new Stack([new ArbInt(1), new ArbInt(1), new ArbInt(3)]), 
	new Stack([]),
	new ArbNone(),
	new ArbNone()
);

let s3 = new Machine(
	new Stack([new ArbInt(25)]), 
	new Stack([new ArbInt(1), new ArbInt(4)]), 
	new Stack([]),
	new ArbNone(),
	new ArbNone()
);

let s4 = new Machine(
	new Stack([]), 
	new Stack([new ArbInt(5)]), 
	new Stack([]),
	new ArbNone(),
	new ArbNone()
);

var oneStepProof;
var arbValue;

var manager1;
var manager2;
var manager3;
var manager4;

let vmId = "0x765435";

async function pause() {
	await new Promise(resolve => setTimeout(resolve, 1000));
}
contract('End to end', function(accounts) {
	before(async function() {
	let vmTracker = await VMTracker.deployed();
    let arbMachine = await ArbProtocol.deployed();
    let arbValue = await ArbValue.deployed();
    manager1 = new ArbManager(vmTracker, arbMachine, arbValue, accounts[1]);
    manager2 = new ArbManager(vmTracker, arbMachine, arbValue, accounts[2]);
    manager3 = new ArbManager(vmTracker, arbMachine, arbValue, accounts[3]);
    manager4 = new ArbManager(vmTracker, arbMachine, arbValue, accounts[4]);

    await vmTracker.mintArbsToUser(accounts[1], 100000);
    await vmTracker.mintArbsToUser(accounts[2], 100000);
    await vmTracker.mintArbsToUser(accounts[3], 100000);
    await vmTracker.mintArbsToUser(accounts[4], 100000);

    oneStepProof = await OneStepProof.deployed();
    arbValue = await ArbValue.deployed();
  });

  it("full test", async function() {

  	let assertionN1 =  {
	  "afterHash": s0.hash(),
	  "numSteps": 1,
	  "messages": []
	};

    let assertion =  {
	  "afterHash": s4.hash(),
	  "numSteps": 4,
	  "messages": []
	};

	let assertion2 =  {
	  "afterHash": s2.hash(),
	  "numSteps": 2,
	  "messages": []
	};

	let assertion3 =  {
	  "afterHash": s1.hash(),
	  "numSteps": 1,
	  "messages": []
	};

	let assertion4 =  {
	  "afterHash": s2.hash(),
	  "numSteps": 1,
	  "messages": []
	};

  	let baseMachine = new Machine(
			new Stack([new ArbInt(25), new ArbInt(25)]), 
			new Stack([new ArbInt(1), new ArbInt(1)]), 
			new Stack([]),
			new ArbNone(),
			new ArbNone()
		);
  	let proof = generateProof(baseMachine, 25, [new ArbInt(1), new ArbInt(2)]);

  	let managers = [accounts[1], accounts[2], accounts[3], accounts[4]];
  	let vmConfig = {
      "gracePeriod": 10,
      "escrowRequired": 50000,
      "assertKeys": managers,
      "maxExecutionSteps":100000,
      "challengeVerifier":0
    };
    let vmTracker = await VMTracker.deployed();
    let createHash = await vmTracker.createVMHash(
      vmConfig.gracePeriod,
      vmConfig.escrowRequired,
      vmConfig.maxExecutionSteps,
      s1n.hash(),
      vmConfig.challengeVerifier,
      vmConfig.assertKeys
    );

    let signatures = [
    	utils.fromRpcSig(web3.eth.sign(accounts[1], createHash)),
    	utils.fromRpcSig(web3.eth.sign(accounts[2], createHash)),
    	utils.fromRpcSig(web3.eth.sign(accounts[3], createHash)),
    	utils.fromRpcSig(web3.eth.sign(accounts[4], createHash))
    ];
    console.log(signatures[0]);
  	let challengeManager = await ChallengeManager.deployed();
  	let vm = await manager1.createVm(vmId, vmConfig, s1n.hash(), signatures, challengeManager);
    let vm2 = await manager2.getVm(vmId, challengeManager, managers, challengeManager);
    let vm3 = await manager3.getVm(vmId, challengeManager, managers, challengeManager);
    let vm4 = await manager4.getVm(vmId, challengeManager, managers, challengeManager);
  	
  	await pause();
	await vm2.disputableAssert(await vm.getVmInfo(), assertionN1);
	await advanceToBlock(web3.eth.blockNumber + 11);
	await vm2.confirmAsserted();
	await pause();
	await vm3.disputableAssert(await vm.getVmInfo(), assertion);
	await pause();
	await vm4.initiateChallenge();
	await pause();
	await vm3.bisectAssertion(assertion2);
	await pause();
	await vm4.continueChallenge(0);
	await pause();
	await vm3.bisectAssertion(assertion3);
	await pause();
	await vm4.continueChallenge(1);
	await pause();
	await vm3.oneStepProof(proof);
  });
});
