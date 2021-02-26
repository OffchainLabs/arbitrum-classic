/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

/* eslint-env node, mocha */
import { ethers } from 'hardhat';
import { assert, expect } from 'chai';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/dist/src/signer-with-address';

describe('Bridge peripherals layer 2', () => {
  let accounts: SignerWithAddress[];
  before(async function() {
    accounts = await ethers.getSigners();
  })
  
  it('should mint erc20 tokens correctly', async function () {
    const TestBridge = await ethers.getContractFactory("TestBridge");
    const testBridge = await TestBridge.deploy();
    console.log("Bridge deployed")

    console.log("current pair")
    console.log(await testBridge.getL1Pair())
    console.log("current account")
    console.log(accounts[0].address)
    console.log("current origin")
    console.log(await testBridge.getOrigin())
    console.log("and again origin")
    console.log(await testBridge.getOrigin())

    const l1ERC20 = "0x0000000000000000000000000000000000000000";
    const account = "0x0000000000000000000000000000000000000000";
    const amount = "1";
    const decimals = "18";

    // const tx = await testBridge.mintERC20FromL1(l1ERC20, account, amount, decimals, {
    //   gasLimit: 99999999999999,

    // });
    // console.log(tx)

    // const l2ERC20 = await testBridge.calculateBridgedERC20Address(l1ERC20);
    // console.log("calculated address");
    // console.log(l2ERC20);

    // assert.notEqual(code, "0x", "Signed liquidity contract not deployed");
  })
})
