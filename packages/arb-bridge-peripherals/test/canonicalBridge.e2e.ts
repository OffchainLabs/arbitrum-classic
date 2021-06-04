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

  let l1RouterTestBridge: Contract
  let l1TestBridge: Contract
  let l2TestBridge: Contract

  const maxSubmissionCost = 0
  const maxGas = 1000000000
  const gasPrice = 0

  before(async function () {
    accounts = await ethers.getSigners()

    // l1 side deploy
    const L1RouterTestBridge: ContractFactory = await ethers.getContractFactory(
      'GatewayRouter'
    )
    l1RouterTestBridge = await L1RouterTestBridge.deploy()

    const L1TestBridge: ContractFactory = await ethers.getContractFactory(
      'L1GatewayTester'
    )
    l1TestBridge = await L1TestBridge.deploy()

    // l2 side deploy

    const StandardArbERC20 = await ethers.getContractFactory('StandardArbERC20')
    const standardArbERC20Logic = await StandardArbERC20.deploy()

    const UpgradeableBeacon = await ethers.getContractFactory(
      'UpgradeableBeacon'
    )
    const beacon = await UpgradeableBeacon.deploy(standardArbERC20Logic.address)

    const L2TestBridge: ContractFactory = await ethers.getContractFactory(
      'L2GatewayTester'
    )
    l2TestBridge = await L2TestBridge.deploy()

    await l1TestBridge.functions['initialize(address,address,address)'](
      l2TestBridge.address,
      l1RouterTestBridge.address,
      accounts[0].address // inbox
    )

    await l2TestBridge.functions['initialize(address,address)'](
      l1TestBridge.address,
      beacon.address
    )

    await l1RouterTestBridge.functions['initialize(address,address,address)'](
      accounts[0].address,
      l1TestBridge.address,
      '0x0000000000000000000000000000000000000000' // no whitelist
    )
  })

  it('should deposit tokens', async function () {
    const Token = await ethers.getContractFactory('TestERC20')
    const token = await Token.deploy()
    // send escrowed tokens to bridge
    const tokenAmount = 100
    await token.mint()
    await token.approve(l1TestBridge.address, tokenAmount)

    const data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    await l1RouterTestBridge.outboundTransfer(
      token.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data
    )

    const escrowedTokens = await token.balanceOf(l1TestBridge.address)
    assert.equal(escrowedTokens.toNumber(), tokenAmount, 'Tokens not escrowed')

    const l2TokenAddress = await l2TestBridge.calculateL2TokenAddress(
      token.address
    )
    const l2Token = await Token.attach(l2TokenAddress)
    const l2Balance = await l2Token.balanceOf(accounts[0].address)
    assert.equal(l2Balance, tokenAmount, 'Tokens not minted')
  })

  it.skip('should withdraw erc20 tokens from L2', async function () {
    const Token = await ethers.getContractFactory('TestERC20')
    const token = await Token.deploy()
    // send escrowed tokens to bridge
    const tokenAmount = 100

    await token.mint()
    await token.approve(l1TestBridge.address, tokenAmount)

    const data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    await l1RouterTestBridge.outboundTransfer(
      token.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data
    )

    const prevUserBalance = await token.balanceOf(accounts[0].address)

    const exitNum = 0
    const withdrawData = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [exitNum, '0x']
    )

    await l1TestBridge.finalizeInboundTransfer(
      token.address,
      accounts[0].address,
      accounts[0].address,
      tokenAmount,
      withdrawData
    )

    const postUserBalance = await token.balanceOf(accounts[0].address)

    assert.equal(
      prevUserBalance.toNumber() + tokenAmount,
      postUserBalance.toNumber(),
      'Tokens not escrowed'
    )
  })
})
