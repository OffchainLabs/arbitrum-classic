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

describe('Buddy bridge layer 1', () => {
  let accounts: SignerWithAddress[]
  let TestBuddy: ContractFactory
  let testBuddy: Contract

  before(async function () {
    accounts = await ethers.getSigners()

    const Mock = await ethers.getContractFactory('Mock')
    const mock = await Mock.deploy('Inbox')

    const inbox = mock.address
    const l2Deployer = '0x0000000000000000000000000000000000000000'

    TestBuddy = await ethers.getContractFactory('TestBuddy')
    testBuddy = await TestBuddy.deploy(inbox, l2Deployer)
  })

  it('should set correct L2 buddy address', async function () {
    assert.equal(
      await testBuddy.l2Buddy(),
      '0x0000000000000000000000000000000000000000',
      'Initial L2 address not 0'
    )
    const maxSubmissionCost = 0
    const maxGas = 999999999999
    const gasPrice = 0
    const deployCode = '0x000000000000000000000000'
    await testBuddy.initiateBuddyDeploy(
      maxSubmissionCost,
      maxGas,
      gasPrice,
      deployCode
    )

    assert.notEqual(
      await testBuddy.l2Buddy(),
      '0x0000000000000000000000000000000000000000',
      'Not implemented'
    )
  })
})
