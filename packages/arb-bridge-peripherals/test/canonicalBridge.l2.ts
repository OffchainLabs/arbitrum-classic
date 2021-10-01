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
import { Contract, ContractFactory } from 'ethers'

const encodeTokenInitData = (
  name: string,
  symbol: string,
  decimals: number | string
) => {
  return ethers.utils.defaultAbiCoder.encode(
    ['bytes', 'bytes', 'bytes'],
    [
      ethers.utils.defaultAbiCoder.encode(['string'], [name]),
      ethers.utils.defaultAbiCoder.encode(['string'], [symbol]),
      ethers.utils.defaultAbiCoder.encode(['uint8'], [decimals]),
    ]
  )
}

describe('Bridge peripherals layer 2', () => {
  let accounts: SignerWithAddress[]
  let TestBridge: ContractFactory
  let testBridge: Contract
  let erc20Proxy: string

  before(async function () {
    // constructor(uint256 _gasPrice, uint256 _maxGas, address erc777Template, address erc20Template)
    accounts = await ethers.getSigners()
    TestBridge = await ethers.getContractFactory('L2ERC20Gateway')
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

    await beaconProxyFactory.initialize(beacon.address)

    testBridge = await TestBridge.deploy()
    await testBridge.initialize(
      accounts[0].address,
      accounts[3].address,
      beaconProxyFactory.address
    )
  })

  it('should deploy erc20 tokens correctly', async function () {
    const l1ERC20 = '0x6100000000000000000000000000000000000001'
    const sender = '0x6300000000000000000000000000000000000002'
    const amount = '10'
    const dest = sender

    const name = 'ArbToken'
    const symbol = 'ATKN'
    const decimals = '18'
    const deployData = encodeTokenInitData(name, symbol, decimals)

    // connect to account 3 to query as if gateway router
    const l2ERC20Address = await testBridge
      .connect(accounts[3])
      .calculateL2TokenAddress(l1ERC20)

    const preTokenCode = await ethers.provider.getCode(l2ERC20Address)
    assert.equal(preTokenCode, '0x', 'Something already deployed to address')

    const data = ethers.utils.defaultAbiCoder.encode(
      ['bytes', 'bytes'],
      [deployData, '0x']
    )

    const tx = await testBridge.finalizeInboundTransfer(
      l1ERC20,
      sender,
      dest,
      amount,
      data
    )

    const postTokenCode = await ethers.provider.getCode(l2ERC20Address)
    assert.notEqual(
      postTokenCode,
      '0x',
      'Token not deployed to correct address'
    )

    const Erc20 = await ethers.getContractFactory('StandardArbERC20')
    const erc20 = await Erc20.attach(l2ERC20Address)

    assert.equal(await erc20.name(), name, 'Tokens not named correctly')
    assert.equal(await erc20.symbol(), symbol, 'Tokens symbol set correctly')
    assert.equal(
      (await erc20.decimals()).toString(),
      decimals,
      'Tokens decimals set incorrectly'
    )
  })

  it('should mint erc20 tokens correctly', async function () {
    const l1ERC20 = '0x0000000000000000000000000000000000000001'
    const sender = '0x0000000000000000000000000000000000000002'
    const dest = sender
    const amount = '1'
    const initializeData = encodeTokenInitData('ArbToken', 'ATKN', '18')

    // connect to account 3 to query as if gateway router
    const l2ERC20Address = await testBridge
      .connect(accounts[3])
      .calculateL2TokenAddress(l1ERC20)

    const preTokenCode = await ethers.provider.getCode(l2ERC20Address)
    assert.equal(preTokenCode, '0x', 'Something already deployed to address')

    const data = ethers.utils.defaultAbiCoder.encode(
      ['bytes', 'bytes'],
      [initializeData, '0x']
    )

    const tx = await testBridge.finalizeInboundTransfer(
      l1ERC20,
      sender,
      dest,
      amount,
      data
    )

    const postTokenCode = await ethers.provider.getCode(l2ERC20Address)
    assert.notEqual(
      postTokenCode,
      '0x',
      'Token not deployed to correct address'
    )

    const Erc20 = await ethers.getContractFactory('StandardArbERC20')
    const erc20 = await Erc20.attach(l2ERC20Address)

    const balance = await erc20.balanceOf(dest)
    assert.equal(balance.toString(), amount, 'Tokens not minted correctly')
  })

  it('should revert post mint call correctly in outbound', async function () {
    const l1ERC20 = '0x0000000000000000000000000000000000000325'
    const sender = '0x0000000000000000000000000000000000000005'
    const amount = '1'

    // connect to account 3 to query as if gateway router
    const l2ERC20Address = await testBridge
      .connect(accounts[3])
      .calculateL2TokenAddress(l1ERC20)

    const L2Called = await ethers.getContractFactory('L2Called')
    const l2Called = await L2Called.deploy()
    const dest = l2Called.address
    // 7 is revert()
    const num = 7
    const callHookData = ethers.utils.defaultAbiCoder.encode(['uint256'], [num])

    const data = ethers.utils.defaultAbiCoder.encode(['bytes'], [callHookData])

    await expect(
      testBridge['outboundTransfer(address,address,uint256,bytes)'](
        l1ERC20,
        dest,
        amount,
        data
      )
    ).to.be.revertedWith('EXTRA_DATA_DISABLED')
  })

  it.skip('should reserve gas in post mint call to ensure rest of function can be executed', async function () {
    // test case skipped since post mint call is disabled
    const l1ERC20 = '0x0000000000000000000000000000000000001325'
    const sender = '0x0000000000000000000000000000000000000015'
    const amount = '1'
    const initializeData = encodeTokenInitData('ArbToken', 'ATKN', '18')

    // connect to account 3 to query as if gateway router
    const l2ERC20Address = await testBridge
      .connect(accounts[3])
      .calculateL2TokenAddress(l1ERC20)

    const preTokenCode = await ethers.provider.getCode(l2ERC20Address)
    assert.equal(preTokenCode, '0x', 'Something already deployed to address')

    const L2Called = await ethers.getContractFactory('L2Called')
    const l2Called = await L2Called.deploy()
    const dest = l2Called.address
    // 9 is assert(false)
    const num = 9
    const callHookData = ethers.utils.defaultAbiCoder.encode(['uint256'], [num])

    const data = ethers.utils.defaultAbiCoder.encode(
      ['bytes', 'bytes'],
      [initializeData, callHookData]
    )

    // we need to hardcode this value as you can only send 63/64 of your remaining
    // gas into a call a high gas limit makes the test pass artificially
    const gasLimit = ethers.BigNumber.from(900000)
    const tx = await testBridge.finalizeInboundTransfer(
      l1ERC20,
      sender,
      dest,
      amount,
      data,
      {
        gasLimit,
      }
    )
    const receipt = await tx.wait()

    const gasUsed = receipt.gasUsed
    const diff = gasLimit.sub(gasUsed)
    // expect to save enough gas for subsequent mint, with at most a 10% margin
    assert(
      diff.mul(10).lte(gasLimit),
      'Did not reserve the correct amount of gas'
    )

    // TransferAndCallTriggered(bool,address,address,uint256,bytes)
    const eventTopic =
      '0x11ff8525c5d96036231ee652c108808dee4c40728a6117830a75029298bb7de6'

    const filteredEvents: Array<any> = receipt.events.filter(
      (event: any) => event.topics[0] === eventTopic
    )

    assert.equal(
      filteredEvents.length,
      1,
      'Token post mint hook should have emitted event'
    )

    const success: boolean = filteredEvents[0].args.success
    assert.equal(success, false, 'Token post mint hook should have reverted')

    // dest should hold not hold amount when reverted
    const Erc20 = await ethers.getContractFactory('aeERC20')
    const erc20 = await Erc20.attach(l2ERC20Address)

    assert.equal(
      (await erc20.balanceOf(dest)).toString(),
      '0',
      'L2Called contract should not be holding coins'
    )
    assert.equal(
      (await erc20.balanceOf(sender)).toString(),
      amount,
      'Sender should hold coins'
    )
  })

  it.skip('should revert post mint call if sent to EOA', async function () {
    // test case skipped since post mint call is disabled
    const l1ERC20 = '0x0000000000000000000000000000000000000326'
    const sender = '0x0000000000000000000000000000000000000005'
    const amount = '1'
    const initializeData = encodeTokenInitData('ArbToken', 'ATKN', '18')

    // connect to account 3 to query as if gateway router
    const l2ERC20Address = await testBridge
      .connect(accounts[3])
      .calculateL2TokenAddress(l1ERC20)

    const preTokenCode = await ethers.provider.getCode(l2ERC20Address)
    assert.equal(preTokenCode, '0x', 'Something already deployed to address')

    const dest = accounts[1].address

    const data = ethers.utils.defaultAbiCoder.encode(
      ['bytes', 'bytes'],
      [initializeData, '0x01']
    )

    const tx = await testBridge.finalizeInboundTransfer(
      l1ERC20,
      sender,
      dest,
      amount,
      data
    )
    const receipt = await tx.wait()

    // TransferAndCallTriggered(bool,address,address,uint256,bytes)
    const eventTopic =
      '0x11ff8525c5d96036231ee652c108808dee4c40728a6117830a75029298bb7de6'

    const filteredEvents: Array<any> = receipt.events.filter(
      (event: any) => event.topics[0] === eventTopic
    )

    assert.equal(
      filteredEvents.length,
      1,
      'Token post mint hook should have emitted event'
    )

    const success: boolean = filteredEvents[0].args.success
    assert.equal(success, false, 'Token post mint hook should have reverted')
  })

  it('should burn on withdraw', async function () {
    const l1ERC20 = '0x0000000000000003000000000000000000000001'
    const sender = accounts[0].address
    const dest = sender
    const amount = '10'
    const initializeData = encodeTokenInitData('ArbToken', 'ATKN', '18')

    // connect to account 3 to query as if gateway router
    const l2ERC20Address = await testBridge
      .connect(accounts[3])
      .calculateL2TokenAddress(l1ERC20)

    const data = ethers.utils.defaultAbiCoder.encode(
      ['bytes', 'bytes'],
      [initializeData, '0x']
    )

    const tx = await testBridge.finalizeInboundTransfer(
      l1ERC20,
      sender,
      dest,
      amount,
      data
    )

    const Erc20 = await ethers.getContractFactory('StandardArbERC20')
    const erc20 = await Erc20.attach(l2ERC20Address)

    const balance = await erc20.balanceOf(dest)
    assert.equal(balance.toString(), amount, 'Tokens not minted correctly')

    await testBridge.functions[
      'outboundTransfer(address,address,uint256,bytes)'
    ](l1ERC20, accounts[1].address, balance, '0x')

    const newBalance = await erc20.balanceOf(dest)
    assert.equal(newBalance.toString(), '0', 'Tokens not minted correctly')
  })

  it('should map L1 to L2 addresses correctly', async function () {
    const TestMessenger = await ethers.getContractFactory('AddressMappingTest')
    const testMessenger = await TestMessenger.deploy()

    const testCases = [
      {
        input: '0x1111000000000000000000000000000000001110',
        expectedOutput: '0xffffffffffffffffffffffffffffffffffffffff',
      },
      {
        input: '0x1111000000000000000000000000081759a885c4',
        expectedOutput: '0x0000000000000000000000000000081759a874b3',
      },
    ]

    for (const { input, expectedOutput } of testCases) {
      const res = await testMessenger.getL1AddressTest(input)
      expect(res.toLowerCase()).to.equal(expectedOutput.toLowerCase())
    }
  })
})
