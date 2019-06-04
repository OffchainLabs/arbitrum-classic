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

  	let assertionN1 = {"afterHash":"0x40c6c305615581f430a99ec0931f6df9e2c5ab945d6147800ed2f443254e16b9","numSteps":1,"messages":[]};

    let assertion =  {"afterHash":"0x1c98d610b01fc3c5f6bcfb36f5e6253f539a7b30d3256770328e3c4897cea7a8","numSteps":4,"messages":[]};

	let assertion2 = {"afterHash":"0xc511a6068c16661a469d646c37e33e206820ad59b62aaa0ea09c3711938a7afe","numSteps":2,"messages":[]};

	let assertion3 = {"afterHash":"0xc511a6068c16661a469d646c37e33e206820ad59b62aaa0ea09c3711938a7afe","numSteps":1,"messages":[]};

  	let baseMachine = new Machine(
			new Stack([new ArbInt(25), new ArbInt(25)]), 
			new Stack([new ArbInt(1), new ArbInt(1)]), 
			new Stack([]),
			new ArbNone(),
			new ArbNone()
		);
  	let proof = '0xcf54b67df0c9d94c9cc3b7112eaab470c6008664eb4b1151c181fe0ee6733a806e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d6e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d6e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01dbef828e20ab46e7edaad8843d62020644c1435a61602d5c466c7c836f0159b8a1f064665568a44e9fdd573a5e15c4dc98ffe61d68015fd505d26482664a9375d178b0695c9d35cfc587a62666cdf784abe97f120d754d343017771610eb15c1796fbae06d3023893ed66cd3dd18fe2bbe97027d915890066fbfe2f9bc8cad41d000944b3061f6d4ba3fbb5243c9ec56ddca89cbdaac736d57f20387bd760af79905c3def2b0646717227ea7933346ff9bda29e15cab4198cd63ce5b7e51e4947bd0c8c879fbb065a4cb1702e81f5e7ca48c02f4f44f4ecf97d2a6dd2b7ab9333a06364a174f419063c8f6b8ab3511bd17d24f95dc5a185f81f1dbb63579a4dfb73e31e81dcb8a3de062ef4bf8973f61e69f1a4f8abd5a884bdb4db90a82a6d0c71036f40d6ab7b8e570000000000000000000000000000000000000000000000000000000000000000';

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
