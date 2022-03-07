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

describe('Bridge peripherals layer 1', () => {
  let accounts: SignerWithAddress[]
  let TestBridge: ContractFactory
  let testBridge: Contract

  let inbox: Contract
  const maxSubmissionCost = 1
  const maxGas = 1000000000
  const gasPrice = 3
  const l2Template20 = '0x0000000000000000000000000000000000000020'
  const l2Address = '0x1100000000000000000000000000000000000011'

  before(async function () {
    accounts = await ethers.getSigners()

    TestBridge = await ethers.getContractFactory('L1GatewayRouter')
    testBridge = await TestBridge.deploy()

    const Inbox = await ethers.getContractFactory('InboxMock')
    inbox = await Inbox.deploy()

    await testBridge.initialize(
      accounts[0].address,
      '0x0000000000000000000000000000000000000000', // default gateway
      '0x0000000000000000000000000000000000000000', // whitelist
      l2Address,
      inbox.address
    )
  })

  it('should submit the correct submission cost to the inbox', async function () {
    const L1ERC20Gateway = await ethers.getContractFactory('L1ERC20Gateway')
    const l1ERC20Gateway = await L1ERC20Gateway.deploy()

    await l1ERC20Gateway.initialize(
      l2Address,
      accounts[0].address,
      inbox.address,
      '0x0000000000000000000000000000000000000000000000000000000000000001', // cloneable proxy hash
      accounts[0].address // beaconProxyFactory
    )

    const tx = await testBridge.setDefaultGateway(
      l1ERC20Gateway.address,
      maxGas,
      gasPrice,
      maxSubmissionCost
    )

    const receipt = await tx.wait()
    const expectedTopic =
      '0x7efacbad201ebbc50ec0ce4b474c54b735a31b1bac996acff50df7de0314e8f9'
    const logs = receipt.events
      .filter((curr: any) => curr.topics[0] === expectedTopic)
      .map((curr: any) => inbox.interface.parseLog(curr))

    assert.equal(
      logs[0].args.maxSubmissionCost.toNumber(),
      maxSubmissionCost,
      'Invalid submission cost'
    )
  })

  it('should submit the correct sender to inbox', async function () {
    const L1ERC20Gateway = await ethers.getContractFactory('L1ERC20Gateway')
    const l1ERC20Gateway = await L1ERC20Gateway.deploy()

    await l1ERC20Gateway.initialize(
      l2Address,
      testBridge.address,
      inbox.address,
      '0x0000000000000000000000000000000000000000000000000000000000000001', // cloneable proxy hash
      accounts[0].address // beaconProxyFactory
    )

    await testBridge.setDefaultGateway(
      l1ERC20Gateway.address,
      maxGas,
      gasPrice,
      maxSubmissionCost
    )

    const Token = await ethers.getContractFactory('TestERC20')
    const token = await Token.deploy()
    const tokenAmount = 100
    await token.mint()
    await token.approve(l1ERC20Gateway.address, tokenAmount)

    const data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    const tx = await testBridge.outboundTransfer(
      token.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data,
      {
        value: maxSubmissionCost + maxGas * gasPrice
      }
    )

    const receipt = await tx.wait()
    // TxToL2(address,address,uint256,bytes)
    const expectedTopic =
      '0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0'
    const logs = receipt.events
      .filter((curr: any) => curr.topics[0] === expectedTopic)
      .map((curr: any) => l1ERC20Gateway.interface.parseLog(curr))
    assert.equal(
      logs[0].args._from,
      accounts[0].address,
      'Invalid from address'
    )
  })

  it('should submit the custom refund address to inbox', async function () {
    const L1ERC20Gateway = await ethers.getContractFactory('L1ERC20Gateway')
    const l1ERC20Gateway = await L1ERC20Gateway.deploy()

    await l1ERC20Gateway.initialize(
      l2Address,
      testBridge.address,
      inbox.address,
      '0x0000000000000000000000000000000000000000000000000000000000000001', // cloneable proxy hash
      accounts[0].address // beaconProxyFactory
    )

    await testBridge.setDefaultGateway(
      l1ERC20Gateway.address,
      maxGas,
      gasPrice,
      maxSubmissionCost
    )

    const Token = await ethers.getContractFactory('TestERC20')
    const token = await Token.deploy()
    const tokenAmount = 100
    await token.mint()
    await token.approve(l1ERC20Gateway.address, tokenAmount)

    const data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    const tx = await testBridge.outboundTransferCustomRefund(
      token.address,
      accounts[0].address,
      accounts[1].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data,
      {
        value: maxSubmissionCost + maxGas * gasPrice
      }
    )

    const receipt = await tx.wait()
    // TxToL2(address,address,uint256,bytes)
    const expectedTopic =
      '0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0'
    const logs = receipt.events
      .filter((curr: any) => curr.topics[0] === expectedTopic)
      .map((curr: any) => l1ERC20Gateway.interface.parseLog(curr))
    assert.equal(
      logs[0].args._from,
      accounts[1].address,
      'Invalid from address'
    )
  })
})
