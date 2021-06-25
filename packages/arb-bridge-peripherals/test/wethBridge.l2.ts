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

describe('Bridge peripherals weth layer 2', () => {
  let accounts: SignerWithAddress[]
  let TestBridge: ContractFactory
  let testBridge: Contract
  const l1WethAddr = '0x0000000000000000000000000000000000004351'
  let l2Weth: Contract

  before(async function () {
    // constructor(uint256 _gasPrice, uint256 _maxGas, address erc777Template, address erc20Template)
    accounts = await ethers.getSigners()
    TestBridge = await ethers.getContractFactory('L2WethGateway')
    const L2Weth = await ethers.getContractFactory('aeWETH')
    l2Weth = await L2Weth.deploy()
    testBridge = await TestBridge.deploy()

    const ProxyAdmin = await ethers.getContractFactory('ProxyAdmin')
    const proxyAdmin = await ProxyAdmin.deploy()
    const Proxy = await ethers.getContractFactory('TransparentUpgradeableProxy')
    const proxy = await Proxy.deploy(
      testBridge.address,
      proxyAdmin.address,
      '0x'
    )

    testBridge = testBridge.attach(proxy.address)

    await testBridge.initialize(
      accounts[0].address, // l1 counterpart
      accounts[3].address, // l2 router
      l1WethAddr,
      l2Weth.address
    )

    await l2Weth.initialize('WETH9', 'WETH', 18, testBridge.address, l1WethAddr)
  })

  it('should deposit weth correctly', async function () {
    const sender = '0x6300000000000000000000000000000000000002'
    const amount = '10'
    const dest = sender

    const data = ethers.utils.defaultAbiCoder.encode(
      ['bytes', 'bytes'],
      ['0x', '0x']
    )

    console.log('here')

    const tx = await testBridge.finalizeInboundTransfer(
      l1WethAddr,
      sender,
      dest,
      amount,
      data,
      {
        value: amount,
      }
    )

    const balance = await l2Weth.balanceOf(dest)

    assert.equal(balance.toString(), amount, 'Deposit failed')
  })

  it('should burn on withdraw', async function () {
    const sender = accounts[0].address
    const dest = sender
    const amount = '10'

    const data = ethers.utils.defaultAbiCoder.encode(
      ['bytes', 'bytes'],
      ['0x', '0x']
    )

    const tx = await testBridge.finalizeInboundTransfer(
      l1WethAddr,
      sender,
      dest,
      amount,
      data,
      {
        value: amount,
      }
    )

    const balance = await l2Weth.balanceOf(dest)
    assert.equal(balance.toString(), amount, 'Tokens not minted correctly')

    await testBridge.functions[
      'outboundTransfer(address,address,uint256,bytes)'
    ](l1WethAddr, accounts[1].address, balance, '0x')

    const newBalance = await l2Weth.balanceOf(dest)
    assert.equal(newBalance.toString(), '0', 'Tokens not burnt correctly')
  })
})
