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

describe('Bridge peripherals end-to-end custom gateway', () => {
  let accounts: SignerWithAddress[]

  let l1RouterTestBridge: Contract
  let l2RouterTestBridge: Contract
  let l1TestBridge: Contract
  let l2TestBridge: Contract

  const maxSubmissionCost = 0
  const maxGas = 1000000000
  const gasPrice = 0

  before(async function () {
    accounts = await ethers.getSigners()

    // l1 side deploy
    const L1RouterTestBridge: ContractFactory = await ethers.getContractFactory(
      'L1GatewayRouter'
    )
    l1RouterTestBridge = await L1RouterTestBridge.deploy()

    const L1TestBridge: ContractFactory = await ethers.getContractFactory(
      'L1CustomGatewayTester'
    )
    l1TestBridge = await L1TestBridge.deploy()

    // l2 side deploy

    const L2TestBridge: ContractFactory = await ethers.getContractFactory(
      'L2CustomGatewayTester'
    )
    l2TestBridge = await L2TestBridge.deploy()

    const L2RouterTestBridge: ContractFactory = await ethers.getContractFactory(
      'L2GatewayRouter'
    )
    l2RouterTestBridge = await L2RouterTestBridge.deploy()

    await l1TestBridge.functions.initialize(
      l2TestBridge.address,
      l1RouterTestBridge.address,
      accounts[0].address, // inbox
      accounts[0].address // owner
    )

    await l2TestBridge.initialize(
      l1TestBridge.address,
      l2RouterTestBridge.address
    )

    await l1RouterTestBridge.functions.initialize(
      accounts[0].address,
      l1TestBridge.address, // defaultGateway
      '0x0000000000000000000000000000000000000000', // no whitelist
      l2RouterTestBridge.address, // counterparty
      accounts[0].address // inbox
    )

    const l2DefaultGateway = await l1TestBridge.counterpartGateway()
    await l2RouterTestBridge.functions.initialize(
      l1RouterTestBridge.address,
      l2DefaultGateway
    )
  })

  it('should deposit tokens', async function () {
    // custom token setup
    const L1CustomToken: ContractFactory = await ethers.getContractFactory(
      'TestCustomTokenL1'
    )
    const l1CustomToken = await L1CustomToken.deploy(l1TestBridge.address)

    const L2Token = await ethers.getContractFactory('TestArbCustomToken')
    const l2Token = await L2Token.deploy(
      l2TestBridge.address,
      l1CustomToken.address
    )

    await l1CustomToken.registerTokenOnL2(l2Token.address, 0, 0, 0)

    // send escrowed tokens to bridge
    const tokenAmount = 100
    await l1CustomToken.mint()
    await l1CustomToken.approve(l1TestBridge.address, tokenAmount)

    const data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    const tx = await l1RouterTestBridge.outboundTransfer(
      l1CustomToken.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data
    )

    const escrowedTokens = await l1CustomToken.balanceOf(l1TestBridge.address)
    assert.equal(escrowedTokens, tokenAmount, 'Tokens not escrowed')

    const l2TokenAddress = await l2RouterTestBridge.calculateL2TokenAddress(
      l1CustomToken.address
    )
    assert.equal(l2TokenAddress, l2Token.address, 'Token Pair not correct')
    const l2Balance = await l2Token.balanceOf(accounts[0].address)
    assert.equal(l2Balance, tokenAmount, 'Tokens not minted')
  })

  it('should withdraw tokens', async function () {
    // custom token setup
    const L1CustomToken: ContractFactory = await ethers.getContractFactory(
      'TestCustomTokenL1'
    )
    const l1CustomToken = await L1CustomToken.deploy(l1TestBridge.address)

    const L2Token = await ethers.getContractFactory('TestArbCustomToken')
    const l2Token = await L2Token.deploy(
      l2TestBridge.address,
      l1CustomToken.address
    )

    await l1CustomToken.registerTokenOnL2(l2Token.address, 0, 0, 0)

    // send escrowed tokens to bridge
    const tokenAmount = 100
    await l1CustomToken.mint()
    await l1CustomToken.approve(l1TestBridge.address, tokenAmount)

    const data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    const tx = await l1RouterTestBridge.outboundTransfer(
      l1CustomToken.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data
    )

    const prevUserBalance = await l1CustomToken.balanceOf(accounts[0].address)

    await l2Token.approve(l2TestBridge.address, tokenAmount)

    await l2TestBridge.functions[
      'outboundTransfer(address,address,uint256,bytes)'
    ](l1CustomToken.address, accounts[0].address, tokenAmount, '0x')

    const postUserBalance = await l1CustomToken.balanceOf(accounts[0].address)

    assert.equal(
      prevUserBalance.toNumber() + tokenAmount,
      postUserBalance.toNumber(),
      'Tokens not escrowed'
    )
  })
  it('should withdraw tokens if no token is deployed', async function () {
    // custom token setup
    const L1CustomToken: ContractFactory = await ethers.getContractFactory(
      'TestCustomTokenL1'
    )
    const l1CustomToken = await L1CustomToken.deploy(l1TestBridge.address)

    // register a non-existent L2 token so we can test the force withdrawal
    await l1CustomToken.registerTokenOnL2(
      '0x0000000000000000000000000000000000000000',
      0,
      0,
      0
    )

    // send escrowed tokens to bridge
    const tokenAmount = 100
    await l1CustomToken.mint()
    await l1CustomToken.approve(l1TestBridge.address, tokenAmount)

    const data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    const prevAllowance = await l1CustomToken.allowance(
      accounts[0].address,
      l1TestBridge.address
    )

    const prevUserBalance = await l1CustomToken.balanceOf(accounts[0].address)

    const tx = await l1RouterTestBridge.outboundTransfer(
      l1CustomToken.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data
    )

    const postUserBalance = await l1CustomToken.balanceOf(accounts[0].address)
    const postAllowance = await l1CustomToken.allowance(
      accounts[0].address,
      l1TestBridge.address
    )

    assert.equal(
      prevUserBalance.toNumber(),
      postUserBalance.toNumber(),
      'Tokens not withdrawn'
    )
    assert.equal(
      prevAllowance.toNumber() - tokenAmount,
      postAllowance.toNumber(),
      'Tokens not spent in allowance'
    )
  })
})
