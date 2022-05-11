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
  L1ERC20Gateway,
  L1GatewayRouter,
  L2ERC20Gateway,
  L2GatewayRouter,
} from '../build/types'
import { processL1ToL2Tx, processL2ToL1Tx } from './testhelper'

describe('Bridge peripherals end-to-end', () => {
  let accounts: SignerWithAddress[]

  let l1RouterTestBridge: L1GatewayRouter
  let l2RouterTestBridge: L2GatewayRouter
  let l1TestBridge: L1ERC20Gateway
  let l2TestBridge: L2ERC20Gateway
  let inboxMock: InboxMock

  const maxSubmissionCost = 1
  const maxGas = 1000000000
  const gasPrice = 1

  before(async function () {
    accounts = await ethers.getSigners()

    const InboxMock = await ethers.getContractFactory('InboxMock')
    inboxMock = await InboxMock.deploy()

    // l1 side deploy
    const L1RouterTestBridge = await ethers.getContractFactory(
      'L1GatewayRouter'
    )
    l1RouterTestBridge = await L1RouterTestBridge.deploy()

    const L1TestBridge = await ethers.getContractFactory('L1ERC20Gateway')
    l1TestBridge = await L1TestBridge.deploy()

    // l2 side deploy

    const StandardArbERC20 = await ethers.getContractFactory('StandardArbERC20')
    const standardArbERC20Logic = await StandardArbERC20.deploy()

    const UpgradeableBeacon = await ethers.getContractFactory(
      'UpgradeableBeacon'
    )
    const beacon = await UpgradeableBeacon.deploy(standardArbERC20Logic.address)
    const BeaconProxyFactory = await ethers.getContractFactory(
      'BeaconProxyFactory'
    )
    const beaconProxyFactory = await BeaconProxyFactory.deploy()
    const cloneableProxyHash = await beaconProxyFactory.cloneableProxyHash()

    await beaconProxyFactory.initialize(beacon.address)

    const L2TestBridge = await ethers.getContractFactory('L2ERC20Gateway')
    l2TestBridge = await L2TestBridge.deploy()

    const L2RouterTestBridge = await ethers.getContractFactory(
      'L2GatewayRouter'
    )
    l2RouterTestBridge = await L2RouterTestBridge.deploy()

    await l1TestBridge.functions.initialize(
      l2TestBridge.address,
      l1RouterTestBridge.address,
      inboxMock.address, // inbox
      cloneableProxyHash,
      beaconProxyFactory.address
    )

    await l2TestBridge.initialize(
      l1TestBridge.address,
      l2RouterTestBridge.address,
      beaconProxyFactory.address
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

    assert.equal(
      await l2TestBridge.cloneableProxyHash(),
      await l1TestBridge.cloneableProxyHash(),
      'Wrong cloneable Proxy Hash'
    )

    const ArbSysMock = await ethers.getContractFactory('ArbSysMock')
    const arbsysmock = await ArbSysMock.deploy()
    await network.provider.send('hardhat_setCode', [
      '0x0000000000000000000000000000000000000064',
      await network.provider.send('eth_getCode', [arbsysmock.address]),
    ])
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

    await expect(
      l1TestBridge.outboundTransfer(
        token.address,
        accounts[0].address,
        tokenAmount,
        maxGas,
        gasPrice,
        data
      )
    ).to.be.revertedWith('NOT_FROM_ROUTER')

    await expect(
      l1RouterTestBridge.outboundTransfer(
        token.address,
        accounts[0].address,
        tokenAmount,
        maxGas,
        gasPrice,
        ethers.utils.defaultAbiCoder.encode(['uint256', 'bytes'], [0, '0x'])
      )
    ).to.be.revertedWith('NO_SUBMISSION_COST')

    await expect(
      l1RouterTestBridge.outboundTransfer(
        token.address,
        accounts[0].address,
        tokenAmount,
        maxGas,
        gasPrice,
        data,
        { value: maxSubmissionCost }
      )
    ).to.be.revertedWith('WRONG_ETH_VALUE')

    const tx = await l1RouterTestBridge.outboundTransfer(
      token.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data,
      { value: maxSubmissionCost + maxGas * gasPrice }
    )
    await processL1ToL2Tx(tx)

    const escrowedTokens = await token.balanceOf(l1TestBridge.address)
    assert.equal(escrowedTokens.toNumber(), tokenAmount, 'Tokens not escrowed')

    const l2TokenAddress = await l2RouterTestBridge.calculateL2TokenAddress(
      token.address
    )
    const l2TokenAddressFromL1Router =
      await l1RouterTestBridge.calculateL2TokenAddress(token.address)
    assert.equal(
      l2TokenAddressFromL1Router,
      l2TokenAddress,
      'Wrong address oracle'
    )

    const l2Token = await Token.attach(l2TokenAddress)
    const l2Balance = await l2Token.balanceOf(accounts[0].address)
    assert.equal(l2Balance.toNumber(), tokenAmount, 'Tokens not minted')
  })

  it('should withdraw erc20 tokens from L2 without router', async function () {
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

    const tx = await l1RouterTestBridge.outboundTransfer(
      token.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data,
      { value: maxSubmissionCost + maxGas * gasPrice }
    )
    await processL1ToL2Tx(tx)

    const prevUserBalance = await token.balanceOf(accounts[0].address)

    await processL2ToL1Tx(
      await l2RouterTestBridge.functions[
        'outboundTransfer(address,address,uint256,bytes)'
      ](token.address, accounts[0].address, tokenAmount, '0x'),
      inboxMock
    )
    // await l2TestBridge.triggerTxToL1()

    const postUserBalance = await token.balanceOf(accounts[0].address)

    assert.equal(
      prevUserBalance.toNumber() + tokenAmount,
      postUserBalance.toNumber(),
      'Tokens not escrowed'
    )
  })

  it('should withdraw erc20 tokens from L2 using router', async function () {
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

    const tx = await l1RouterTestBridge.outboundTransfer(
      token.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data,
      { value: maxSubmissionCost + maxGas * gasPrice }
    )
    await processL1ToL2Tx(tx)

    const prevUserBalance = await token.balanceOf(accounts[0].address)

    await processL2ToL1Tx(
      await l2RouterTestBridge.functions[
        'outboundTransfer(address,address,uint256,bytes)'
      ](token.address, accounts[0].address, tokenAmount, '0x'),
      inboxMock
    )
    // await l2TestBridge.triggerTxToL1()

    const postUserBalance = await token.balanceOf(accounts[0].address)

    assert.equal(
      prevUserBalance.toNumber() + tokenAmount,
      postUserBalance.toNumber(),
      'Tokens not escrowed'
    )
  })

  it('should force withdraw correctly if deposit is incorrect', async function () {
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

    const prevUserBalance = await token.balanceOf(accounts[0].address)
    const prevAllowance = await token.allowance(
      accounts[0].address,
      l1TestBridge.address
    )

    // here we set the L2 router to recover in case of a bad BeaconProxyFactory deploy
    // await l2TestBridge.setStubAddressOracleReturn(accounts[0].address)

    const tx = await l1RouterTestBridge.outboundTransfer(
      token.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data,
      { value: maxSubmissionCost + maxGas * gasPrice }
    )
    const original_code = await accounts[0].provider?.getCode(token.address)
    await network.provider.send('hardhat_setCode', [
      await l2TestBridge.calculateL2TokenAddress(token.address),
      '0x69',
    ])
    await processL2ToL1Tx((await processL1ToL2Tx(tx))[0], inboxMock)
    await network.provider.send('hardhat_setCode', [
      await l2TestBridge.calculateL2TokenAddress(token.address),
      original_code,
    ])
    // await l2TestBridge.triggerTxToL1()

    const postUserBalance = await token.balanceOf(accounts[0].address)
    const postAllowance = await token.allowance(
      accounts[0].address,
      l1TestBridge.address
    )

    assert.equal(
      prevUserBalance.toNumber(),
      postUserBalance.toNumber(),
      'Tokens not escrowed'
    )
    assert.equal(
      prevAllowance.toNumber() - tokenAmount,
      postAllowance.toNumber(),
      'Tokens not spent in allowance'
    )

    const l2TokenAddress = await l1RouterTestBridge.calculateL2TokenAddress(
      token.address
    )

    const l2Token = await Token.attach(l2TokenAddress)
    const l2Balance = await l2Token.balanceOf(accounts[0].address)

    assert.equal(l2Balance.toNumber(), 0, 'User has tokens in L2')

    // reset stub return in test case
    // await l2TestBridge.setStubAddressOracleReturn(ethers.constants.AddressZero)
  })

  it('should deposit tokens with bytes32 field correctly', async function () {
    const Token = await ethers.getContractFactory('Bytes32ERC20WithMetadata')
    const token = await Token.deploy()
    // send escrowed tokens to bridge
    const tokenAmount = 100
    await token.mint()
    await token.approve(l1TestBridge.address, tokenAmount)

    const data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    const tx = await l1RouterTestBridge.outboundTransfer(
      token.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data,
      { value: maxSubmissionCost + maxGas * gasPrice }
    )
    await processL1ToL2Tx(tx)

    const l2TokenAddress = await l2RouterTestBridge.calculateL2TokenAddress(
      token.address
    )

    const l2Code = await ethers.provider.getCode(l2TokenAddress)
    assert.notEqual(l2Code, '0x', 'No code at L2 token address')

    const l2Token = await ethers.getContractAt(
      'StandardArbERC20',
      l2TokenAddress
    )

    const name = await l2Token.name()
    const symbol = await l2Token.symbol()

    assert.equal(name, 'Maker')
    assert.equal(symbol, 'MKR')
  })

  it('should not have L2 getters for unavailable fields', async function () {
    const Token = await ethers.getContractFactory('Bytes32ERC20')
    const token = await Token.deploy()
    // send escrowed tokens to bridge
    const tokenAmount = 100
    await token.mint()
    await token.approve(l1TestBridge.address, tokenAmount)

    const data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    const tx = await l1RouterTestBridge.outboundTransfer(
      token.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data,
      { value: maxSubmissionCost + maxGas * gasPrice }
    )
    await processL1ToL2Tx(tx)

    const l2TokenAddress = await l2RouterTestBridge.calculateL2TokenAddress(
      token.address
    )

    const l2Code = await ethers.provider.getCode(l2TokenAddress)
    assert.notEqual(l2Code, '0x', 'No code at L2 token address')

    const l2Token = await ethers.getContractAt(
      'StandardArbERC20',
      l2TokenAddress
    )

    await expect(l2Token.name()).to.be.revertedWith('')
    await expect(l2Token.symbol()).to.be.revertedWith('')
    await expect(l2Token.decimals()).to.be.revertedWith('')
  })
})
