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
import { Contract, ContractFactory } from 'ethers'
import { AeWETH, L2WethGateway, L2WethGateway__factory } from '../build/types'
import { applyAlias, impersonateAccount } from './testhelper'

describe('Bridge peripherals weth layer 2', () => {
  let accounts: SignerWithAddress[]
  let TestBridge: L2WethGateway__factory
  let testBridge: L2WethGateway
  const l1WethAddr = '0x0000000000000000000000000000000000004351'
  let l2Weth: AeWETH

  before(async function () {
    // constructor(uint256 _gasPrice, uint256 _maxGas, address erc777Template, address erc20Template)
    accounts = await ethers.getSigners()
    TestBridge = await ethers.getContractFactory('L2WethGateway')
    const L2Weth = (await ethers.getContractFactory('aeWETH')) as any
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

    const ArbSysMock = await ethers.getContractFactory('ArbSysMock')
    const arbsysmock = await ArbSysMock.deploy()
    await network.provider.send('hardhat_setCode', [
      '0x0000000000000000000000000000000000000064',
      await network.provider.send('eth_getCode', [arbsysmock.address]),
    ])

    testBridge = testBridge
      .attach(proxy.address)
      .connect(await impersonateAccount(applyAlias(accounts[0].address)))

    await expect(
      l2Weth.initialize('WETH9', 'WETH', 18, testBridge.address, l1WethAddr)
    ).to.be.revertedWith('Initializable: contract is already initialized')
    const wethProxy = await Proxy.deploy(
      l2Weth.address,
      accounts[1].address,
      '0x'
    )
    l2Weth = await L2Weth.attach(wethProxy.address)

    await l2Weth.initialize('WETH9', 'WETH', 18, testBridge.address, l1WethAddr)
    await testBridge.initialize(
      accounts[0].address, // l1 counterpart
      accounts[3].address, // l2 router
      l1WethAddr,
      l2Weth.address
    )
  })

  it('should deposit weth correctly', async function () {
    const sender = '0x6300000000000000000000000000000000000002'
    const amount = '10'
    const dest = sender

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

    assert.equal(balance.toString(), amount, 'Deposit failed')
  })

  it('should burn on withdraw', async function () {
    const sender = applyAlias(accounts[0].address)
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
