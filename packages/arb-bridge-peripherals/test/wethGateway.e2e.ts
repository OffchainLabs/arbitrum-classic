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
import { ethers, network } from 'hardhat'
import { assert, expect } from 'chai'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import {
  InboxMock,
  L1WethGateway,
  L1GatewayRouter,
  L2WethGateway,
  L2GatewayRouter,
} from '../build/types'
import { processL1ToL2Tx, processL2ToL1Tx } from './testhelper'
import { Contract } from 'ethers'

describe('Bridge peripherals end-to-end weth gateway', () => {
  let accounts: SignerWithAddress[]

  let l1RouterTestBridge: L1GatewayRouter
  let l2RouterTestBridge: L2GatewayRouter
  let l1TestBridge: L1WethGateway
  let l2TestBridge: L2WethGateway
  let l1Weth: Contract
  let l2Weth: Contract
  let inboxMock: InboxMock

  const maxSubmissionCost = 1
  const maxGas = 1000000000
  const gasPrice = 1

  before(async function () {
    accounts = await ethers.getSigners()

    const InboxMock = await ethers.getContractFactory('InboxMock')
    inboxMock = await InboxMock.deploy()

    const TestWETH9 = await ethers.getContractFactory('TestWETH9')

    // l1 side deploy
    const L1RouterTestBridge = await ethers.getContractFactory(
      'L1GatewayRouter'
    )
    l1RouterTestBridge = await L1RouterTestBridge.deploy()

    const L1TestBridge = await ethers.getContractFactory('L1WethGateway')
    l1TestBridge = await L1TestBridge.deploy()

    l1Weth = await TestWETH9.deploy('wethl1', 'wl1')
    l1Weth = await l1Weth.deployed()

    // l2 side deploy
    const L2TestBridge = await ethers.getContractFactory('L2WethGateway')
    l2TestBridge = await L2TestBridge.deploy()

    const L2RouterTestBridge = await ethers.getContractFactory(
      'L2GatewayRouter'
    )
    l2RouterTestBridge = await L2RouterTestBridge.deploy()

    const L2Weth = await ethers.getContractFactory('aeWETH')
    l2Weth = await L2Weth.deploy()
    await l2Weth.deployed()
    // initialize contracts

    await expect(
      l2Weth.initialize(
        'l2weth',
        'l2w',
        18,
        l2TestBridge.address,
        l1Weth.address
      )
    ).to.be.revertedWith('Initializable: contract is already initialized')

    const Proxy = await ethers.getContractFactory('TransparentUpgradeableProxy')
    l2Weth = await Proxy.deploy(l2Weth.address, accounts[1].address, '0x')
    l2Weth = await L2Weth.attach(l2Weth.address)
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
      inboxMock.address, // inbox
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
      inboxMock.address // inbox
    )

    const l2DefaultGateway = await l1TestBridge.counterpartGateway()
    await l2RouterTestBridge.functions.initialize(
      l1RouterTestBridge.address,
      l2DefaultGateway
    )

    const ArbSysMock = await ethers.getContractFactory('ArbSysMock')
    const arbsysmock = await ArbSysMock.deploy()
    await network.provider.send('hardhat_setCode', [
      '0x0000000000000000000000000000000000000064',
      await network.provider.send('eth_getCode', [arbsysmock.address]),
    ])
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

    await processL1ToL2Tx(
      await l1RouterTestBridge.outboundTransfer(
        l1Weth.address,
        accounts[0].address,
        tokenAmount,
        maxGas,
        gasPrice,
        data,
        { value: maxSubmissionCost + maxGas * gasPrice }
      )
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

    await processL1ToL2Tx(
      await l1RouterTestBridge.outboundTransfer(
        l1Weth.address,
        accounts[0].address,
        tokenAmount,
        maxGas,
        gasPrice,
        data,
        { value: maxSubmissionCost + maxGas * gasPrice }
      )
    )

    const prevUserBalance = await l1Weth.balanceOf(accounts[0].address)

    await l2Weth.approve(l2TestBridge.address, tokenAmount)

    await processL2ToL1Tx(
      await l2TestBridge.functions[
        'outboundTransfer(address,address,uint256,bytes)'
      ](l1Weth.address, accounts[0].address, tokenAmount, '0x'),
      inboxMock
    )

    const postUserBalance = await l1Weth.balanceOf(accounts[0].address)

    assert.equal(
      prevUserBalance.toNumber() + tokenAmount,
      postUserBalance.toNumber(),
      'Tokens not escrowed'
    )
  })

  it('should withdraw tokens if no token is deployed', async function () {
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

    const tx = await l1RouterTestBridge.outboundTransfer(
      l1Weth.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data,
      { value: maxSubmissionCost + maxGas * gasPrice }
    )
    const original_code = await accounts[0].provider?.getCode(l2Weth.address)
    // Set l2Weth to invalid code to trigger force withdraw
    await network.provider.send('hardhat_setCode', [l2Weth.address, '0x00'])
    await processL2ToL1Tx((await processL1ToL2Tx(tx))[0], inboxMock)
    // Revert previous setCode
    await network.provider.send('hardhat_setCode', [
      await l2TestBridge.calculateL2TokenAddress(l2Weth.address),
      original_code,
    ])

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
  })
})
