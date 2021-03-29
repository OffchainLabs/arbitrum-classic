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
import { ethers } from 'hardhat'
import { assert, expect } from 'chai'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/dist/src/signer-with-address'
import { Contract, ContractFactory } from 'ethers'

describe('Bridge peripherals layer 1', () => {
  let accounts: SignerWithAddress[]
  let TestBridge: ContractFactory
  let testBridge: Contract

  let inbox
  let l2Deployer
  const maxSubmissionCost = 0
  const maxGas = 1000000000
  const gasPrice = 0
  const l2Template777 = '0x0000000000000000000000000000000000000777'
  const l2Template20 = '0x0000000000000000000000000000000000000020'
  const l2Address = '0x1100000000000000000000000000000000000011'

  before(async function () {
    accounts = await ethers.getSigners()

    TestBridge = await ethers.getContractFactory('EthERC20Bridge')
    testBridge = await TestBridge.deploy()

    inbox = accounts[0].address
    l2Deployer = accounts[0].address

    await testBridge.initialize(
      inbox,
      l2Deployer,
      maxSubmissionCost,
      maxGas,
      gasPrice,
      l2Template777,
      l2Template20,
      l2Address
    )
  })

  it.skip('should withdraw from L2', async function () {
    assert.equal(true, false, 'Not implemented')
  })

  it.skip('should deposit erc20 token to L2', async function () {
    assert.equal(true, false, 'Not implemented')
  })

  it.skip('should not deposit erc777 token to L2', async function () {
    assert.equal(true, false, 'Not implemented')
  })

  it.skip('should updateTokenInfo', async function () {
    assert.equal(true, false, 'Not implemented')
  })

  it.skip('should deposit custom token', async function () {})
  it.skip('should registerCustomL2Token', async function () {})
  it.skip('should notifyCustomToken', async function () {})
  it.skip('should fastWithdrawalFromL2', async function () {})
})
