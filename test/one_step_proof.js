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
const utils = require('ethereumjs-util')
var abi = require('ethereumjs-abi')

import {ArbVM} from './vmTrackerWrapper'
import {ArbManager} from './vmManager'
import {ArbInt, ArbNone} from './vm-arb-value';
import {Machine, Stack} from './arb/vm-machine';
import generateProof from './arb/vm-proof';

var OneStepProof = artifacts.require("OneStepProof");
var ArbValue = artifacts.require("ArbValue");
var oneStepProof;
var arbValue;

contract('OneStepProof', function(accounts) {
	before(async function() {
    oneStepProof = await OneStepProof.deployed();
    arbValue = await ArbValue.deployed();
  });

  it("addition proof", async function() {
  	let beforeMachine = new Machine(
  		new Stack([new ArbInt(25)]), 
  		new Stack([new ArbInt(1), new ArbInt(1)]), 
  		new Stack([]),
  		new ArbNone(),
  		new ArbNone()
  	);

  	let afterMachine = new Machine(
  		new Stack([]),
  		new Stack([new ArbInt(2)]),
  		new Stack([]),
  		new ArbNone(),
  		new ArbNone()
  	);

  	let baseMachine = new Machine(
  		new Stack([]), 
  		new Stack([]), 
  		new Stack([]),
  		new ArbNone(),
  		new ArbNone()
  	);

  	let proof = generateProof(baseMachine, 25, [new ArbInt(1), new ArbInt(1)]);
    oneStepProof.validateProof(
    	beforeMachine.hash(),
    	[0, 100000],
    	0,
    	'0x00',
    	afterMachine.hash(),
    	[],
    	proof
    ).then(function(result) {
    	console.log("test1234:", result);
    });
  });
});
