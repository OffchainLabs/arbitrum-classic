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

describe('Bridge peripherals', () => {
  let accounts: SignerWithAddress[];
  before(async function() {
    accounts = await ethers.getSigners();
  })
  
  it('should deploy bridge contracts correctly', async function () {
    const signedLiquidityFactory = await ethers.getContractFactory("SignedLiquidityProvider");
    const bridge = "0x0000000000000000000000000000000000000000";
    const signedLiquidity = await signedLiquidityFactory.deploy(bridge, accounts[0].address);
    const code = await ethers.provider.getCode(signedLiquidity.address);
    assert.notEqual(code, "0x", "Signed liquidity contract not deployed");
  })
})
