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
import { SignerWithAddress } from 'hardhat-deploy-ethers/dist/src/signers'
import { Contract, ContractFactory } from 'ethers'

describe('Bridge peripherals end-to-end weth gateway', () => {
  let accounts: SignerWithAddress[]

  let l1RouterTestBridge: Contract
  let l2RouterTestBridge: Contract
  let l1TestBridge: Contract
  let l2TestBridge: Contract
  let l1Weth: Contract
  let l2Weth: Contract

  const maxSubmissionCost = 0
  const maxGas = 1000000000
  const gasPrice = 0

  before(async function () {
    accounts = await ethers.getSigners()

    const TestWETH9: ContractFactory = await ethers.getContractFactory(
      'TestWETH9'
    )

    // l1 side deploy
    const L1RouterTestBridge: ContractFactory = await ethers.getContractFactory(
      'L1GatewayRouter'
    )
    l1RouterTestBridge = await L1RouterTestBridge.deploy()

    const L1TestBridge: ContractFactory = await ethers.getContractFactory(
      'L1WethGatewayTester'
    )
    l1TestBridge = await L1TestBridge.deploy()

    l1Weth = await TestWETH9.deploy('wethl1', 'wl1')
    l1Weth = await l1Weth.deployed()

    // l2 side deploy

    const L2TestBridge: ContractFactory = await ethers.getContractFactory(
      'L2WethGatewayTester'
    )
    l2TestBridge = await L2TestBridge.deploy()

    const L2RouterTestBridge: ContractFactory = await ethers.getContractFactory(
      'L2GatewayRouter'
    )
    l2RouterTestBridge = await L2RouterTestBridge.deploy()

    const L2Weth: ContractFactory = await ethers.getContractFactory('aeWETH')
    l2Weth = await L2Weth.deploy()
    await l2Weth.deployed()
    // initialize contracts

    await l2Weth.initialize(
      'l2weth',
      'l2w',
      18,
      l2TestBridge.address,
      l1Weth.address
    )

    await l1TestBridge.functions.initialize(
      l2TestBridge.address,
      l1RouterTestBridge.address,
      accounts[0].address, // inbox
      l1Weth.address,
      l2Weth.address
    )

    await l2TestBridge.initialize(
      l1TestBridge.address,
      l2RouterTestBridge.address,
      l1Weth.address,
      l2Weth.address
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
    // send escrowed tokens to bridge
    const tokenAmount = 100
    await l1Weth.deposit({ value: tokenAmount })

    const initialDepositTokens = await l1Weth.balanceOf(accounts[0].address)
    assert.equal(initialDepositTokens, tokenAmount, 'Tokens not deposited')

    await l1Weth.approve(l1TestBridge.address, tokenAmount)

    const data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    const expectedRouter = await l1RouterTestBridge.getGateway(l1Weth.address)
    assert.equal(
      expectedRouter,
      l1TestBridge.address,
      'Router not setup correctly'
    )

    const l2ExpectedAddress = await l1RouterTestBridge.calculateL2TokenAddress(
      l1Weth.address
    )
    assert.equal(
      l2ExpectedAddress,
      l2Weth.address,
      'Not expected l2 weth address'
    )

    const tx = await l1RouterTestBridge.outboundTransfer(
      l1Weth.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data
    )

    const escrowedTokens = await l1Weth.balanceOf(l1TestBridge.address)
    assert.equal(escrowedTokens, 0, 'Tokens should not be escrowed')

    const l2TokenAddress = await l2RouterTestBridge.calculateL2TokenAddress(
      l1Weth.address
    )
    assert.equal(l2TokenAddress, l2Weth.address, 'Token Pair not correct')
    const l2Balance = await l2Weth.balanceOf(accounts[0].address)
    assert.equal(l2Balance, tokenAmount, 'Tokens not minted')
  })

  it('should withdraw tokens', async function () {
    const tokenAmount = 100
    await l1Weth.deposit({ value: tokenAmount })

    const initialDepositTokens = await l1Weth.balanceOf(accounts[0].address)
    assert.equal(initialDepositTokens, tokenAmount, 'Tokens not deposited')

    await l1Weth.approve(l1TestBridge.address, tokenAmount)

    const data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    const expectedRouter = await l1RouterTestBridge.getGateway(l1Weth.address)
    assert.equal(
      expectedRouter,
      l1TestBridge.address,
      'Router not setup correctly'
    )

    const l2ExpectedAddress = await l1RouterTestBridge.calculateL2TokenAddress(
      l1Weth.address
    )
    assert.equal(
      l2ExpectedAddress,
      l2Weth.address,
      'Not expected l2 weth address'
    )

    const tx = await l1RouterTestBridge.outboundTransfer(
      l1Weth.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data
    )

    const prevUserBalance = await l1Weth.balanceOf(accounts[0].address)

    await l2Weth.approve(l2TestBridge.address, tokenAmount)

    const withdrawTx = await l2TestBridge.functions[
      'outboundTransfer(address,address,uint256,bytes)'
    ](l1Weth.address, accounts[0].address, tokenAmount, '0x')

    const postUserBalance = await l1Weth.balanceOf(accounts[0].address)

    assert.equal(
      prevUserBalance.toNumber() + tokenAmount,
      postUserBalance.toNumber(),
      'Tokens not escrowed'
    )
  })

  it('should withdraw tokens if no token is deployed', async function () {
    const ZERO_ADDR = '0x0000000000000000000000000000000000000000'
    // set L2 weth to address zero to test if force withdraw is triggered when
    // no contract is deployed
    await l2TestBridge.setL2WethAddress(ZERO_ADDR)

    // send escrowed tokens to bridge
    const tokenAmount = 100
    await l1Weth.deposit({ value: tokenAmount })
    await l1Weth.approve(l1TestBridge.address, tokenAmount)

    const prevUserBalance = await l1Weth.balanceOf(accounts[0].address)
    const prevAllowance = await l1Weth.allowance(
      accounts[0].address,
      l1TestBridge.address
    )

    const data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    await l1RouterTestBridge.outboundTransfer(
      l1Weth.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data
    )

    const postUserBalance = await l1Weth.balanceOf(accounts[0].address)
    const postAllowance = await l1Weth.allowance(
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
    // unset the custom l2 address as to not affect other tests
    await l2TestBridge.setL2WethAddress(l2Weth.address)
  })
})
