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
  const gasPrice = 0
  const l2Template20 = '0x0000000000000000000000000000000000000020'
  let l2Address: string
  // const l2Address = '0x1100000000000000000000000000000000000011'

  before(async function () {
    accounts = await ethers.getSigners()
    l2Address = accounts[0].address

    TestBridge = await ethers.getContractFactory('L1GatewayTester')
    testBridge = await TestBridge.deploy()

    const Inbox = await ethers.getContractFactory('InboxMock')
    inbox = await Inbox.deploy()

    await testBridge.initialize(
      l2Address,
      accounts[0].address,
      inbox.address,
      '0x0000000000000000000000000000000000000000000000000000000000000001', // cloneable proxy hash
      accounts[0].address // beaconProxyFactory
    )
  })

  it('should escrow depositted tokens', async function () {
    const Token = await ethers.getContractFactory('TestERC20')
    const token = await Token.deploy()
    // send escrowed tokens to bridge
    const tokenAmount = 100
    await token.mint()
    await token.approve(testBridge.address, tokenAmount)

    let data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    // router usually does this encoding part
    data = ethers.utils.defaultAbiCoder.encode(
      ['address', 'bytes'],
      [accounts[0].address, data]
    )

    await testBridge.outboundTransfer(
      token.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data
    )

    const escrowedTokens = await token.balanceOf(testBridge.address)
    assert.equal(escrowedTokens.toNumber(), tokenAmount, 'Tokens not escrowed')
  })

  it('should revert post mint call correctly in outbound', async function () {
    const Token = await ethers.getContractFactory('TestERC20')
    const token = await Token.deploy()
    // send escrowed tokens to bridge
    const tokenAmount = 100
    await token.mint()
    await token.approve(testBridge.address, tokenAmount)

    let data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x01']
    )

    // router usually does this encoding part
    data = ethers.utils.defaultAbiCoder.encode(
      ['address', 'bytes'],
      [accounts[0].address, data]
    )

    await expect(
      testBridge.outboundTransfer(
        token.address,
        accounts[0].address,
        tokenAmount,
        maxGas,
        gasPrice,
        data
      )
    ).to.be.revertedWith('EXTRA_DATA_DISABLED')
  })

  it('should revert on inbound if there is data for post mint call', async function () {
    const Token = await ethers.getContractFactory('TestERC20')
    const token = await Token.deploy()
    // send escrowed tokens to bridge
    const tokenAmount = 100

    await token.mint()
    await token.approve(testBridge.address, tokenAmount)

    let data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x12']
    )

    // router usually does this encoding part
    data = ethers.utils.defaultAbiCoder.encode(
      ['address', 'bytes'],
      [accounts[0].address, data]
    )

    const exitNum = 0
    const withdrawData = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [exitNum, '0x11']
    )

    await expect(
      testBridge.finalizeInboundTransfer(
        token.address,
        accounts[0].address,
        accounts[0].address,
        tokenAmount,
        withdrawData
      )
    ).to.be.revertedWith('')
  })
  it('should withdraw erc20 tokens from L2', async function () {
    const Token = await ethers.getContractFactory('TestERC20')
    const token = await Token.deploy()
    // send escrowed tokens to bridge
    const tokenAmount = 100

    await token.mint()
    await token.approve(testBridge.address, tokenAmount)

    let data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    // router usually does this encoding part
    data = ethers.utils.defaultAbiCoder.encode(
      ['address', 'bytes'],
      [accounts[0].address, data]
    )

    await testBridge.outboundTransfer(
      token.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data
    )

    await inbox.setL2ToL1Sender(l2Address)

    const prevUserBalance = await token.balanceOf(accounts[0].address)

    const exitNum = 0
    const withdrawData = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [exitNum, '0x']
    )

    await testBridge.finalizeInboundTransfer(
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

    const Token = await ethers.getContractFactory('TestERC20')
    const token = await Token.deploy()
    // send escrowed tokens to bridge
    const tokenAmount = 100
    await token.mint()
    await token.approve(l1ERC20Gateway.address, tokenAmount)

    let data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    data = ethers.utils.defaultAbiCoder.encode(
      ['address', 'bytes'],
      [accounts[0].address, data]
    )

    const tx = await l1ERC20Gateway.outboundTransfer(
      token.address,
      accounts[0].address,
      tokenAmount,
      maxGas,
      gasPrice,
      data
    )
    const receipt = await tx.wait()
    // TicketData(uint256)
    const expectedTopic =
      '0x7efacbad201ebbc50ec0ce4b474c54b735a31b1bac996acff50df7de0314e8f9'
    const events = receipt.events

    if (!events) {
      const msg = 'No events in receipt'
      assert(events, msg)
      throw new Error(msg)
    }

    const logs = events
      .filter((curr: any) => curr.topics[0] === expectedTopic)
      .map((curr: any) => inbox.interface.parseLog(curr))

    assert.equal(
      logs[0].args.maxSubmissionCost.toNumber(),
      maxSubmissionCost,
      'Invalid submission cost'
    )

    const escrowedTokens = await token.balanceOf(l1ERC20Gateway.address)
    assert.equal(escrowedTokens.toNumber(), tokenAmount, 'Tokens not escrowed')
  })
})
