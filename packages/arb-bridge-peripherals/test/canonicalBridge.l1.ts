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

  let inbox: Contract
  const maxSubmissionCost = 0
  const maxGas = 1000000000
  const gasPrice = 0
  const l2Template20 = '0x0000000000000000000000000000000000000020'
  const l2Address = '0x1100000000000000000000000000000000000011'

  before(async function () {
    accounts = await ethers.getSigners()

    TestBridge = await ethers.getContractFactory('EthERC20Bridge')
    testBridge = await TestBridge.deploy()

    const Inbox = await ethers.getContractFactory('InboxMock')
    inbox = await Inbox.deploy()

    await testBridge.initialize(inbox.address, l2Template20, l2Address)
  })

  it('should escrow depositted tokens', async function () {
    const Token = await ethers.getContractFactory('TestERC20')
    const token = await Token.deploy()
    // send escrowed tokens to bridge
    const tokenAmount = 100
    await token.mint()
    await token.approve(testBridge.address, tokenAmount)

    await testBridge.deposit(
      token.address,
      accounts[0].address,
      tokenAmount,
      0,
      1000000,
      0,
      '0x'
    )

    const escrowedTokens = await token.balanceOf(testBridge.address)
    assert.equal(escrowedTokens.toNumber(), tokenAmount, 'Tokens not escrowed')
  })

  it('should withdraw erc20 tokens from L2', async function () {
    const Token = await ethers.getContractFactory('TestERC20')
    const token = await Token.deploy()
    // send escrowed tokens to bridge
    const tokenAmount = 100

    await token.mint()
    await token.approve(testBridge.address, tokenAmount)
    await testBridge.deposit(
      token.address,
      accounts[0].address,
      tokenAmount,
      0,
      1000000,
      0,
      '0x'
    )

    await inbox.setL2ToL1Sender(l2Address)

    const prevUserBalance = await token.balanceOf(accounts[0].address)

    const exitNum = 0
    await testBridge.withdrawFromL2(
      exitNum,
      token.address,
      accounts[0].address,
      tokenAmount
    )

    const postUserBalance = await token.balanceOf(accounts[0].address)

    assert.equal(
      prevUserBalance.toNumber() + tokenAmount,
      postUserBalance.toNumber(),
      'Tokens not escrowed'
    )
  })

  it.skip('should not deposit erc777 token to L2', async function () {
    assert.equal(true, false, 'Not implemented')
  })

  it.skip('should updateTokenInfo 18 decimals', async function () {
    // deploy erc20 with 18 decimals
    const Token = await ethers.getContractFactory('StandardArbERC20')
    const token = await Token.deploy()

    const newDecimals = 18
    const newName = 'Test Token'
    const newSymbol = 'TT'

    await token.initialize(
      accounts[0].address,
      accounts[0].address,
      newDecimals
    )

    await token.updateInfo(newName, newSymbol, newDecimals)

    const tokenType = 0
    const tx = await testBridge.updateTokenInfo(
      token.address,
      tokenType,
      maxSubmissionCost,
      maxGas,
      gasPrice
    )
    const receipt = await tx.wait()

    // event UpdateTokenInfo
    const eventTopic =
      '0x0388926a40418e22c6e6e9024bedafa0f215f76f61b5c2a069dccfc5c4335d9c'
    const events = receipt.events.filter((e: any) => e.topics[0] === eventTopic)

    assert.equal(events.length, 1, 'Expected only one event to be emitted')

    const event = events[0]

    const {
      // seqNum,
      // l1Address,
      name: eventName,
      symbol: eventSymbol,
      decimals: eventDecimals,
    } = event.args

    assert.equal(
      eventName,
      '0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000a5465737420546f6b656e00000000000000000000000000000000000000000000',
      'Incorrect encoded name'
    )
    assert.equal(
      eventSymbol,
      '0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000025454000000000000000000000000000000000000000000000000000000000000',
      'Incorrect encoded symbol'
    )

    assert.equal(
      eventDecimals,
      '0x0000000000000000000000000000000000000000000000000000000000000012',
      'Incorrect encoded symbol'
    )
  })

  it.skip('should updateTokenInfo 6 decimals as uint8', async function () {
    // deploy erc20 with 18 decimals
    const Token = await ethers.getContractFactory('StandardArbERC20')
    const token = await Token.deploy()

    const newDecimals = 6
    const newName = 'Test Token'
    const newSymbol = 'TT'

    await token.initialize(
      accounts[0].address,
      accounts[0].address,
      newDecimals
    )

    await token.updateInfo(newName, newSymbol, newDecimals)

    const tokenType = 0
    const tx = await testBridge.updateTokenInfo(
      token.address,
      tokenType,
      maxSubmissionCost,
      maxGas,
      gasPrice
    )
    const receipt = await tx.wait()

    // event UpdateTokenInfo
    const eventTopic =
      '0x0388926a40418e22c6e6e9024bedafa0f215f76f61b5c2a069dccfc5c4335d9c'
    const events = receipt.events.filter((e: any) => e.topics[0] === eventTopic)

    assert.equal(events.length, 1, 'Expected only one event to be emitted')

    const event = events[0]

    const {
      // seqNum,
      // l1Address,
      name: eventName,
      symbol: eventSymbol,
      decimals: eventDecimals,
    } = event.args

    assert.equal(
      eventName,
      '0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000a5465737420546f6b656e00000000000000000000000000000000000000000000',
      'Incorrect encoded name'
    )
    assert.equal(
      eventSymbol,
      '0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000025454000000000000000000000000000000000000000000000000000000000000',
      'Incorrect encoded symbol'
    )

    assert.equal(
      eventDecimals,
      '0x0000000000000000000000000000000000000000000000000000000000000006',
      'Incorrect encoded symbol'
    )
  })

  it.skip('should updateTokenInfo 6 decimals set as uint 256 and name as bytes32', async function () {
    const Token = await ethers.getContractFactory('TesterERC20Token')
    // this adds padding at the end, not the start!
    const name = ethers.utils.formatBytes32String('0x617262697472756d')
    const symbol = ethers.utils.formatBytes32String('0x617262')
    const decimal = 6

    const token = await Token.deploy(decimal, name, symbol)

    const tokenType = 0
    const tx = await testBridge.updateTokenInfo(
      token.address,
      tokenType,
      maxSubmissionCost,
      maxGas,
      gasPrice
    )
    const receipt = await tx.wait()

    // event UpdateTokenInfo
    const eventTopic =
      '0x0388926a40418e22c6e6e9024bedafa0f215f76f61b5c2a069dccfc5c4335d9c'
    const events = receipt.events.filter((e: any) => e.topics[0] === eventTopic)

    assert.equal(events.length, 1, 'Expected only one event to be emitted')

    const event = events[0]

    const {
      // seqNum,
      // l1Address,
      name: eventName,
      symbol: eventSymbol,
      decimals: eventDecimals,
    } = event.args

    assert.equal(
      eventName,
      '0x3078363137323632363937343732373536640000000000000000000000000000',
      'Incorrect encoded name'
    )
    assert.equal(
      eventSymbol,
      '0x3078363137323632000000000000000000000000000000000000000000000000',
      'Incorrect encoded symbol'
    )

    assert.equal(
      eventDecimals,
      '0x0000000000000000000000000000000000000000000000000000000000000006',
      'Incorrect encoded symbol'
    )
  })

  it.skip('should updateTokenInfo even with token that has no metadata', async function () {
    const Token = await ethers.getContractFactory('TesterERC20TokenNoMetadata')
    const token = await Token.deploy()

    const tokenType = 0
    const tx = await testBridge.updateTokenInfo(
      token.address,
      tokenType,
      maxSubmissionCost,
      maxGas,
      gasPrice
    )
    const receipt = await tx.wait()

    // event UpdateTokenInfo
    const eventTopic =
      '0x0388926a40418e22c6e6e9024bedafa0f215f76f61b5c2a069dccfc5c4335d9c'
    const events = receipt.events.filter((e: any) => e.topics[0] === eventTopic)

    assert.equal(events.length, 1, 'Expected only one event to be emitted')

    const event = events[0]

    const {
      // seqNum,
      // l1Address,
      name: eventName,
      symbol: eventSymbol,
      decimals: eventDecimals,
    } = event.args

    assert.equal(eventName, '0x', 'Incorrect encoded name')
    assert.equal(eventSymbol, '0x', 'Incorrect encoded symbol')

    assert.equal(eventDecimals, '0x', 'Incorrect encoded symbol')
  })

  it.skip('should deposit custom token', async function () {})
  it.skip('should registerCustomL2Token', async function () {})
  it.skip('should notifyCustomToken', async function () {})
  it.skip('should fastWithdrawalFromL2', async function () {})
})
