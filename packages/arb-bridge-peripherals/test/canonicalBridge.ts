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

    const l1ERC20 = "0x0000000000000000000000000000000000000001";
    const account = "0x0000000000000000000000000000000000000002";
    const amount = "1";
    const decimals = "18";

    const l2ERC20Address = await testBridge.calculateBridgedERC20Address(l1ERC20);

    const preTokenCode = await ethers.provider.getCode(l2ERC20Address);
    assert.equal(preTokenCode, "0x", "Something already deployed to address");

    const tx = await testBridge.mintERC20FromL1(l1ERC20, account, amount, decimals);

    const postTokenCode = await ethers.provider.getCode(l2ERC20Address);
    assert.notEqual(postTokenCode, "0x", "Token not deployed to correct address");

    const Erc20 = await ethers.getContractFactory("StandardArbERC20");
    const erc20 = await Erc20.attach(l2ERC20Address);

    const balance = await erc20.balanceOf(account);
    assert.equal(balance.toString(), amount, "Tokens not minted correctly");
  })
})
