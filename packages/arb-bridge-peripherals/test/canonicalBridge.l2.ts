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
import { deploy1820Registry } from '../scripts/utils'

const TOKEN_TYPE_ENUM = {
  ERC20: 0,
  ERC777: 1,
  Custom: 2,
}

describe('Bridge peripherals layer 2', () => {
  let accounts: SignerWithAddress[]
  let TestBridge: ContractFactory
  let testBridge: Contract
  let erc20Proxy: string

  before(async function () {
    // constructor(uint256 _gasPrice, uint256 _maxGas, address erc777Template, address erc20Template)
    accounts = await ethers.getSigners()
    TestBridge = await ethers.getContractFactory('ArbTokenBridge')
    const StandardArbERC20 = await ethers.getContractFactory('StandardArbERC20')
    const StandardArbERC777 = await ethers.getContractFactory(
      'StandardArbERC777'
    )
    const standardArbERC20Logic = await StandardArbERC20.deploy()
    const standardArbERC777Logic = await StandardArbERC777.deploy()

    const UpgradeableBeacon = await ethers.getContractFactory(
      'UpgradeableBeacon'
    )

    const standardArbERC20Proxy = await UpgradeableBeacon.deploy(
      standardArbERC20Logic.address
    )

    const standardArbERC777Proxy = await UpgradeableBeacon.deploy(
      standardArbERC20Logic.address
    )
    erc20Proxy = standardArbERC20Proxy.address
    testBridge = await TestBridge.deploy()
    await testBridge.initialize(
      accounts[0].address,
      standardArbERC777Proxy.address,
      standardArbERC20Proxy.address
    )

    await deploy1820Registry(accounts[0])
  })

  it('should calculate proxy address correctly', async function () {
    const address: string = (await testBridge.functions.templateERC20())[0]
    // OZ's init code not the same as in https://eips.ethereum.org/EIPS/eip-1167
    const proxyBytecode =
      '0x3d602d80600a3d3981f3363d3d373d3d3d363d73' +
      address.substr(2) +
      '5af43d82803e903d91602b57fd5bf3'

    const ClonableBeaconProxy = await ethers.getContractFactory(
      'ClonableBeaconProxy'
    )
    const l1ERC20 = '0x0000000000000000000000000000000000000001'
    const l2ERC20Address = await testBridge.calculateBridgedERC20Address(
      l1ERC20
    )
    const salt = ethers.utils.solidityKeccak256(
      ['address', 'address'],
      [l1ERC20, erc20Proxy]
    )
    const initCodeHash = ethers.utils.keccak256(ClonableBeaconProxy.bytecode)
    const l2AddressExpected = ethers.utils.getCreate2Address(
      testBridge.address,
      salt,
      initCodeHash
    )
    assert.equal(
      l2ERC20Address,
      l2AddressExpected,
      'Address calculated incorrectly'
    )
  })
  it('should mint erc20 tokens correctly', async function () {
    const l1ERC20 = '0x0000000000000000000000000000000000000001'
    const sender = '0x0000000000000000000000000000000000000002'
    const dest = sender
    const amount = '1'
    const decimals = ethers.utils.defaultAbiCoder.encode(['uint8'], ['18'])

    const l2ERC20Address = await testBridge.calculateBridgedERC20Address(
      l1ERC20
    )

    const preTokenCode = await ethers.provider.getCode(l2ERC20Address)
    assert.equal(preTokenCode, '0x', 'Something already deployed to address')

    const tx = await testBridge.mintFromL1(
      l1ERC20,
      sender,
      TOKEN_TYPE_ENUM.ERC20,
      dest,
      amount,
      decimals,
      '0x'
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

  it('should execute post mint call', async function () {
    const l1ERC20 = '0x0000000000000000000000000000000000000305'
    const sender = '0x0000000000000000000000000000000000000003'
    const amount = '1'
    const decimals = ethers.utils.defaultAbiCoder.encode(['uint8'], ['18'])

    const l2ERC20Address = await testBridge.calculateBridgedERC20Address(
      l1ERC20
    )

    const preTokenCode = await ethers.provider.getCode(l2ERC20Address)
    assert.equal(preTokenCode, '0x', 'Something already deployed to address')

    const L2Called = await ethers.getContractFactory('L2Called')
    const l2Called = await L2Called.deploy()
    const dest = l2Called.address
    const num = 5
    const callHookData = ethers.utils.defaultAbiCoder.encode(['uint256'], [num])

    const tx = await testBridge.mintFromL1(
      l1ERC20,
      sender,
      TOKEN_TYPE_ENUM.ERC20,
      dest,
      amount,
      decimals,
      callHookData
    )
    const receipt = await tx.wait()

    // Called(uint256)
    const eventTopic =
      '0xfea238a9794376fb3707dbbabe56f0fad5e4110a7839485387c458f2d1aa5d50'

    const filteredEvents: Array<any> = receipt.events.filter(
      (event: any) => event.topics[0] === eventTopic
    )

    assert.equal(filteredEvents.length, 1, 'Token post mint hook not triggered')

    const actualNum = ethers.BigNumber.from(filteredEvents[0].data)
    assert.equal(
      actualNum.toNumber(),
      num,
      'Token event called in hook emitted wrong num'
    )

    // dest should hold amount and sender 0
    const Erc20 = await ethers.getContractFactory('OZERC20')
    const erc20 = await Erc20.attach(l2ERC20Address)

    assert.equal(
      (await erc20.balanceOf(dest)).toString(),
      amount,
      'L2Called contract not holding coins'
    )
    assert.equal(
      (await erc20.balanceOf(sender)).toString(),
      '0',
      'Sender should not hold any coins'
    )
  })

  it('should revert post mint call correctly', async function () {
    const l1ERC20 = '0x0000000000000000000000000000000000000325'
    const sender = '0x0000000000000000000000000000000000000005'
    const amount = '1'
    const decimals = ethers.utils.defaultAbiCoder.encode(['uint8'], ['18'])

    const l2ERC20Address = await testBridge.calculateBridgedERC20Address(
      l1ERC20
    )

    const preTokenCode = await ethers.provider.getCode(l2ERC20Address)
    assert.equal(preTokenCode, '0x', 'Something already deployed to address')

    const L2Called = await ethers.getContractFactory('L2Called')
    const l2Called = await L2Called.deploy()
    const dest = l2Called.address
    const num = 7
    const callHookData = ethers.utils.defaultAbiCoder.encode(['uint256'], [num])

    const tx = await testBridge.mintFromL1(
      l1ERC20,
      sender,
      TOKEN_TYPE_ENUM.ERC20,
      dest,
      amount,
      decimals,
      callHookData
    )
    const receipt = await tx.wait()

    // MintAndCallTriggered(bool,address,address,uint256,bytes)
    const eventTopic =
      '0xe934ad33409d1a25da34f3e31354e20013f314d227c3d53952d3e130ece06011'

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
    const Erc20 = await ethers.getContractFactory('OZERC20')
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

  it('should revert post mint call if sent to EOA', async function () {
    const l1ERC20 = '0x0000000000000000000000000000000000000326'
    const sender = '0x0000000000000000000000000000000000000005'
    const amount = '1'
    const decimals = ethers.utils.defaultAbiCoder.encode(['uint8'], ['18'])

    const l2ERC20Address = await testBridge.calculateBridgedERC20Address(
      l1ERC20
    )

    const preTokenCode = await ethers.provider.getCode(l2ERC20Address)
    assert.equal(preTokenCode, '0x', 'Something already deployed to address')

    const dest = accounts[1].address

    const tx = await testBridge.mintFromL1(
      l1ERC20,
      sender,
      TOKEN_TYPE_ENUM.ERC20,
      dest,
      amount,
      decimals,
      '0x01'
    )
    const receipt = await tx.wait()

    // MintAndCallTriggered(bool,address,address,uint256,bytes)
    const eventTopic =
      '0xe934ad33409d1a25da34f3e31354e20013f314d227c3d53952d3e130ece06011'

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

  it.skip('should mint erc777 tokens correctly', async function () {
    const l1ERC20 = '0x0000000000000000000000000000000000000001'
    const sender = '0x0000000000000000000000000000000000000002'
    const dest = sender
    const amount = 10
    const decimalVal = 18
    const decimals = ethers.utils.defaultAbiCoder.encode(
      ['uint8'],
      [decimalVal]
    )

    const l2ERC777Address = await testBridge.calculateBridgedERC777Address(
      l1ERC20
    )

    const preTokenCode = await ethers.provider.getCode(l2ERC777Address)
    assert.equal(preTokenCode, '0x', 'Something already deployed to address')

    const tx = await testBridge.mintFromL1(
      l1ERC20,
      sender,
      TOKEN_TYPE_ENUM.ERC777,
      dest,
      amount,
      decimals,
      '0x'
    )

    const postTokenCode = await ethers.provider.getCode(l2ERC777Address)
    assert.notEqual(
      postTokenCode,
      '0x',
      'Token not deployed to correct address'
    )

    const Erc777 = await ethers.getContractFactory('StandardArbERC777')
    const erc777 = await Erc777.attach(l2ERC777Address)

    const balance = await erc777.balanceOf(dest)
    assert.equal(balance.toString(), amount, 'Tokens not minted correctly')
  })

  it.skip('should burn and mint tokens correctly on migrate', async function () {
    const l1ERC20 = '0x0000000000000000000000000000000000000002'
    const sender = accounts[0].address
    const dest = sender
    const amount = '1'
    const decimals = ethers.utils.defaultAbiCoder.encode(['uint8'], ['18'])

    const tx20 = await testBridge.mintFromL1(
      l1ERC20,
      sender,
      TOKEN_TYPE_ENUM.ERC20,
      dest,
      amount,
      decimals,
      '0x'
    )
    const tx777 = await testBridge.mintFromL1(
      l1ERC20,
      sender,
      TOKEN_TYPE_ENUM.ERC777,
      dest,
      amount,
      decimals,
      '0x'
    )

    const l2ERC20Address = await testBridge.calculateBridgedERC20Address(
      l1ERC20
    )
    const l2ERC777Address = await testBridge.calculateBridgedERC777Address(
      l1ERC20
    )

    const Erc20 = await ethers.getContractFactory('StandardArbERC20')
    const erc20 = await Erc20.attach(l2ERC20Address)
    const Erc777 = await ethers.getContractFactory('StandardArbERC777')
    const erc777 = await Erc777.attach(l2ERC777Address)

    // const balance20 = await erc20.balanceOf(account)
    // const balance777 = await erc777.balanceOf(account)

    const migrate = await erc20.migrate(amount, l2ERC777Address, '0x')
    const newBalance777 = await erc777.balanceOf(dest)
    const newBalance20 = await erc20.balanceOf(dest)
    assert.equal(
      newBalance777.toString(),
      '2',
      'Tokens not migrated correctly on mint'
    )
    assert.equal(
      newBalance20.toString(),
      '0',
      'Tokens not migrated correctly on burn'
    )
  })

  it.skip('should fail to migrate from erc20 to non-deployed erc777', async function () {
    const l1ERC20 = '0x0000000000000000000000000000000000000003'
    const sender = accounts[0].address
    const dest = sender
    const amount = '1'
    const decimals = ethers.utils.defaultAbiCoder.encode(['uint8'], ['18'])

    const tx20 = await testBridge.mintFromL1(
      l1ERC20,
      sender,
      TOKEN_TYPE_ENUM.ERC20,
      dest,
      amount,
      decimals,
      '0x'
    )

    const l2ERC20Address = await testBridge.calculateBridgedERC20Address(
      l1ERC20
    )
    const l2ERC777Address = await testBridge.calculateBridgedERC777Address(
      l1ERC20
    )

    const Erc20 = await ethers.getContractFactory('StandardArbERC20')
    const erc20 = await Erc20.attach(l2ERC20Address)
    try {
      const migrate = await erc20.migrate(amount, l2ERC777Address, '0x')
      assert.equal(true, false, 'Migration should have failed')
    } catch (e) {
      assert.equal(e.message, 'execution reverted', 'Migration did not fail')
    }
  })

  it('should burn on withdraw', async function () {
    const l1ERC20 = '0x0000000000000000000000000000000000000001'
    const sender = accounts[0].address
    const dest = sender
    const amount = '10'
    const decimals = ethers.utils.defaultAbiCoder.encode(['uint8'], ['18'])

    const l2ERC777Address = await testBridge.calculateBridgedERC20Address(
      l1ERC20
    )

    const tx = await testBridge.mintFromL1(
      l1ERC20,
      sender,
      TOKEN_TYPE_ENUM.ERC20,
      dest,
      amount,
      decimals,
      '0x'
    )

    const Erc20 = await ethers.getContractFactory('StandardArbERC777')
    const erc20 = await Erc20.attach(l2ERC777Address)

    const balance = await erc20.balanceOf(dest)
    assert.equal(balance.toString(), amount, 'Tokens not minted correctly')

    await erc20.withdraw(accounts[1].address, balance)

    const newBalance = await erc20.balanceOf(dest)
    assert.equal(newBalance.toString(), '0', 'Tokens not minted correctly')
  })
})
