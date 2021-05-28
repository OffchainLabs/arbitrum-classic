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

    await testBridge.initialize(
      inbox.address,
      l2Template20,
      l2Address,
      accounts[0].address
    )
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
      maxSubmissionCost,
      maxGas,
      gasPrice,
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
      maxSubmissionCost,
      maxGas,
      gasPrice,
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

  it('should process fast withdrawal correctly', async function () {
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

    // parameters used for exit
    const exitNum = 0
    const maxFee = 10
    const liquidityProof = '0x'

    const FastExitMock = await ethers.getContractFactory('FastExitMock')
    const fastExitMock = await FastExitMock.deploy()

    await fastExitMock.setFee(maxFee)

    // send tokens to liquidity provider
    const liquidityProviderBalance = 10000
    await token.transfer(fastExitMock.address, liquidityProviderBalance)

    const prevUserBalance = await token.balanceOf(accounts[0].address)

    // request liquidity from them
    const PassiveFastExitManager = await ethers.getContractFactory(
      'PassiveFastExitManager'
    )
    const passiveFastExitManager = await PassiveFastExitManager.deploy()
    await passiveFastExitManager.setBridge(testBridge.address)

    const data = ethers.utils.defaultAbiCoder.encode(
      ['address', 'uint256', 'uint256', 'address', 'bytes', 'bytes'],
      [
        accounts[0].address,
        exitNum,
        maxFee,
        fastExitMock.address,
        liquidityProof,
        '0x',
      ]
    )

    await testBridge.transferExitAndCall(
      accounts[0].address,
      token.address,
      tokenAmount,
      exitNum,
      passiveFastExitManager.address,
      data
    )

    const postUserBalance = await token.balanceOf(accounts[0].address)

    assert.equal(
      prevUserBalance.toNumber() + tokenAmount - maxFee,
      postUserBalance.toNumber(),
      'Tokens not escrowed'
    )

    await inbox.setL2ToL1Sender(l2Address)

    // withdrawal should now be sent to liquidity provider
    // const prevLPBalance = await token.balanceOf(expensiveFastExitMock[0].address)

    await testBridge.withdrawFromL2(
      exitNum,
      token.address,
      accounts[0].address,
      tokenAmount
    )

    const postLPBalance = await token.balanceOf(fastExitMock.address)

    assert.equal(
      postLPBalance.toNumber(),
      liquidityProviderBalance + maxFee,
      'Liquidity provider balance not as expected'
    )
  })
})
