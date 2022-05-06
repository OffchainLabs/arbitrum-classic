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
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import {
  L1CustomGatewayTester,
  L1GatewayRouter,
  L2CustomGatewayTester,
  L2GatewayRouter,
} from '../build/types'

describe('Bridge peripherals end-to-end custom gateway', () => {
  let accounts: SignerWithAddress[]

  let l1RouterTestBridge: L1GatewayRouter
  let l2RouterTestBridge: L2GatewayRouter
  let l1TestBridge: L1CustomGatewayTester
  let l2TestBridge: L2CustomGatewayTester

  const maxSubmissionCost = 1
  const maxGas = 1000000000
  const gasPrice = 1

  before(async function () {
    accounts = await ethers.getSigners()

    // l1 side deploy
    const L1RouterTestBridge = await ethers.getContractFactory(
      'L1GatewayRouterTester'
    )
    l1RouterTestBridge = await L1RouterTestBridge.deploy()

    const L1TestBridge = await ethers.getContractFactory(
      'L1CustomGatewayTester'
    )
    l1TestBridge = await L1TestBridge.deploy()

    // l2 side deploy

    const L2TestBridge = await ethers.getContractFactory(
      'L2CustomGatewayTester'
    )
    l2TestBridge = await L2TestBridge.deploy()

    const L2RouterTestBridge = await ethers.getContractFactory(
      'L2GatewayRouterTester'
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
      ethers.constants.AddressZero, // l1TestBridge.address, // defaultGateway
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
    const L1CustomToken = await ethers.getContractFactory(
      'TestCustomTokenL1'
    )
    const l1CustomToken = await L1CustomToken.deploy(
      l1TestBridge.address,
      l1RouterTestBridge.address
    )

    const L2Token = await ethers.getContractFactory('TestArbCustomToken')
    const l2Token = await L2Token.deploy(
      l2TestBridge.address,
      l1CustomToken.address
    )

    await l1CustomToken.registerTokenOnL2(
      l2Token.address,
      0,
      0,
      0,
      0,
      0,
      0,
      0,
      accounts[0].address
    )

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
      data,
      { value: maxSubmissionCost + maxGas * gasPrice }
    )

    const escrowedTokens = await l1CustomToken.balanceOf(l1TestBridge.address)
    assert.equal(escrowedTokens.toNumber(), tokenAmount, 'Tokens not escrowed')

    const l2TokenAddress = await l2RouterTestBridge.calculateL2TokenAddress(
      l1CustomToken.address
    )
    assert.equal(l2TokenAddress, l2Token.address, 'Token Pair not correct')
    const l2Balance = await l2Token.balanceOf(accounts[0].address)
    assert.equal(l2Balance.toNumber(), tokenAmount, 'Tokens not minted')
  })

  it('should withdraw tokens', async function () {
    // custom token setup
    const L1CustomToken = await ethers.getContractFactory(
      'TestCustomTokenL1'
    )
    const l1CustomToken = await L1CustomToken.deploy(
      l1TestBridge.address,
      l1RouterTestBridge.address
    )

    const L2Token = await ethers.getContractFactory('TestArbCustomToken')
    const l2Token = await L2Token.deploy(
      l2TestBridge.address,
      l1CustomToken.address
    )

    await l1TestBridge.forceRegisterTokenToL2(
      [l1CustomToken.address],
      [l2Token.address],
      0,
      0,
      0
    )
    await l1RouterTestBridge.setGateways(
      [l1CustomToken.address],
      [l1TestBridge.address],
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

    const tx = await l1RouterTestBridge.outboundTransfer(
      l1CustomToken.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data,
      { value: maxSubmissionCost + maxGas * gasPrice }
    )

    const prevUserBalance = await l1CustomToken.balanceOf(accounts[0].address)

    await l2TestBridge.functions[
      'outboundTransfer(address,address,uint256,bytes)'
    ](l1CustomToken.address, accounts[0].address, tokenAmount, '0x')
    await l2TestBridge.triggerTxToL1()

    const postUserBalance = await l1CustomToken.balanceOf(accounts[0].address)

    assert.equal(
      prevUserBalance.toNumber() + tokenAmount,
      postUserBalance.toNumber(),
      'Tokens not escrowed'
    )
  })
  it('should force withdraw tokens if no token is deployed in L2', async function () {
    // custom token setup
    const L1CustomToken = await ethers.getContractFactory(
      'TestCustomTokenL1'
    )
    const l1CustomToken = await L1CustomToken.deploy(
      l1TestBridge.address,
      l1RouterTestBridge.address
    )

    // register a non-existent L2 token so we can test the force withdrawal
    await l1TestBridge.forceRegisterTokenToL2(
      [l1CustomToken.address],
      ['0x0000000000000000000000000000000000000001'],
      0,
      0,
      0
    )
    await l1RouterTestBridge.setGateways(
      [l1CustomToken.address],
      [l1TestBridge.address],
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
      data,
      { value: maxSubmissionCost + maxGas * gasPrice }
    )
    await l2TestBridge.triggerTxToL1()

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

  it('should withdraw tokens when minted in L2', async function () {
    // custom token setup
    const L1CustomToken = await ethers.getContractFactory(
      'MintableTestCustomTokenL1'
    )
    const l1CustomToken = await L1CustomToken.deploy(
      l1TestBridge.address,
      l1RouterTestBridge.address
    )

    const L2Token = await ethers.getContractFactory(
      'MintableTestArbCustomToken'
    )
    const l2Token = await L2Token.deploy(
      l2TestBridge.address,
      l1CustomToken.address
    )

    await l1TestBridge.forceRegisterTokenToL2(
      [l1CustomToken.address],
      [l2Token.address],
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

    await l1RouterTestBridge.outboundTransfer(
      l1CustomToken.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data,
      { value: maxSubmissionCost + maxGas * gasPrice }
    )

    // mint tokens for the user in L2
    await l2Token.userMint(accounts[0].address, tokenAmount)
    const l2Balance = await l2Token.balanceOf(accounts[0].address)

    assert.equal(
      tokenAmount,
      l2Balance.div(2).toNumber(),
      'Wrong user L2 balance'
    )

    const prevUserBalance = await l1CustomToken.balanceOf(accounts[0].address)
    const prevEscrow = await l1CustomToken.balanceOf(l1TestBridge.address)

    // do a small withdrawal that will have enough collateral
    const smallWithdrawal = tokenAmount / 2
    await l2TestBridge.functions[
      'outboundTransfer(address,address,uint256,bytes)'
    ](l1CustomToken.address, accounts[0].address, smallWithdrawal, '0x')

    const midUserBalance = await l1CustomToken.balanceOf(accounts[0].address)
    const midEscrow = await l1CustomToken.balanceOf(l1TestBridge.address)

    assert.equal(
      midUserBalance.toNumber(),
      prevUserBalance.add(smallWithdrawal).toNumber(),
      'Wrong user balance in initial withdrawal'
    )
    assert.equal(
      midEscrow.toNumber(),
      prevEscrow.sub(smallWithdrawal).toNumber(),
      'Wrong escrow balance in initial withdrawal'
    )

    await expect(
      l2TestBridge.functions['outboundTransfer(address,address,uint256,bytes)'](
        l1CustomToken.address,
        accounts[0].address,
        l2Balance.sub(smallWithdrawal),
        '0x'
      )
    )
      .to.emit(l1CustomToken, 'Transfer(address,address,uint256)')
      .withArgs(ethers.constants.AddressZero, l1TestBridge.address, tokenAmount) // this is the mint

    const postUserBalance = await l1CustomToken.balanceOf(accounts[0].address)
    const postEscrow = await l1CustomToken.balanceOf(l1TestBridge.address)

    assert.equal(prevEscrow.toNumber(), tokenAmount)
    assert.equal(postEscrow.toNumber(), 0)

    assert.equal(
      prevUserBalance.add(l2Balance).toNumber(),
      postUserBalance.toNumber(),
      'Tokens not escrowed'
    )
  })
})
