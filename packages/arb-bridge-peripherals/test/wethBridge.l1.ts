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
import { ethers, waffle } from 'hardhat'
import { assert, expect } from 'chai'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { Contract, ContractFactory, providers } from 'ethers'
import { TestWETH9 } from '../build/types'

describe('Bridge peripherals layer 1', () => {
  let accounts: SignerWithAddress[]
  let TestBridge: ContractFactory
  let testBridge: Contract

  let inbox: Contract
  const maxSubmissionCost = 1
  const maxGas = 1000000000
  const gasPrice = 0
  let l2Address: string
  let weth: TestWETH9;
  
  before(async function () {
    accounts = await ethers.getSigners()
    l2Address = accounts[1].address

    TestBridge = await ethers.getContractFactory('L1WethGatewayTester')
    testBridge = await TestBridge.deploy()

    const Inbox = await ethers.getContractFactory('InboxMock')
    inbox = await Inbox.deploy()

    const Weth = await ethers.getContractFactory('TestWETH9')
    weth = await Weth.deploy('weth','weth')

    await testBridge.initialize(
      l2Address,
      accounts[0].address,
      inbox.address,
      weth.address, // _l1Weth
      accounts[0].address, // _l2Weth
    )
  })

  it('should escrow deposited weth as eth', async function () {
    // send weth to bridge
    const wethAmount = 100
    await weth.deposit({ value: wethAmount })
    await weth.approve(testBridge.address, wethAmount)

    let data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    // router usually does this encoding part
    data = ethers.utils.defaultAbiCoder.encode(
      ['address', 'bytes'],
      [accounts[0].address, data]
    )
    const escrowPrevBalance = await waffle.provider.getBalance(l2Address)
    await testBridge.outboundTransfer(
      weth.address,
      accounts[0].address,
      wethAmount,
      maxGas,
      gasPrice,
      data
    )
    const escrowedWeth = await weth.balanceOf(testBridge.address)
    assert.equal(escrowedWeth.toNumber(), 0, 'Weth should not be escrowed')
    const escrowedETH = await waffle.provider.getBalance(l2Address)
    assert.equal(escrowedETH.sub(escrowPrevBalance).toNumber(), wethAmount, 'ETH should be escrowed')
  })

  it('should escrow deposited weth as eth (new entrypoint)', async function () {
    // send weth to bridge
    const wethAmount = 100
    await weth.deposit({ value: wethAmount })
    await weth.approve(testBridge.address, wethAmount)

    let data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    // router usually does this encoding part
    data = ethers.utils.defaultAbiCoder.encode(
      ['address', 'bytes'],
      [accounts[0].address, data]
    )
    const escrowPrevBalance = await waffle.provider.getBalance(l2Address)
    await testBridge.outboundTransferCustomRefund(
      weth.address,
      accounts[0].address,
      accounts[0].address,
      wethAmount,
      maxGas,
      gasPrice,
      data
    )
    const escrowedWeth = await weth.balanceOf(testBridge.address)
    assert.equal(escrowedWeth.toNumber(), 0, 'Weth should not be escrowed')
    const escrowedETH = await waffle.provider.getBalance(l2Address)
    assert.equal(escrowedETH.sub(escrowPrevBalance).toNumber(), wethAmount, 'ETH should be escrowed')
  })

  it('should revert post mint call correctly in outbound', async function () {
    // send weth to bridge
    const wethAmount = 100
    await weth.deposit({ value: wethAmount })
    await weth.approve(testBridge.address, wethAmount)

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
        weth.address,
        accounts[0].address,
        wethAmount,
        maxGas,
        gasPrice,
        data
      )
    ).to.be.revertedWith('EXTRA_DATA_DISABLED')
  })

  it('should revert on inbound if there is data for post mint call', async function () {
    // send weth to bridge
    const wethAmount = 100
    await weth.deposit({ value: wethAmount })
    await weth.approve(testBridge.address, wethAmount)

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
        weth.address,
        accounts[0].address,
        accounts[0].address,
        wethAmount,
        withdrawData
      )
    ).to.be.revertedWith('')
  })
  it.skip('should withdraw weth from L2', async function () {
    // send weth to bridge
    const wethAmount = 100
    await weth.deposit({ value: wethAmount })
    await weth.approve(testBridge.address, wethAmount)

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
      weth.address,
      accounts[0].address,
      wethAmount,
      maxGas,
      gasPrice,
      data
    )

    await inbox.setL2ToL1Sender(l2Address)

    const prevUserBalance = await weth.balanceOf(accounts[0].address)

    const exitNum = 0
    const withdrawData = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [exitNum, '0x']
    )

    await testBridge.finalizeInboundTransfer(
      weth.address,
      accounts[0].address,
      accounts[0].address,
      wethAmount,
      withdrawData
    )

    const postUserBalance = await weth.balanceOf(accounts[0].address)

    assert.equal(
      prevUserBalance.toNumber() + wethAmount,
      postUserBalance.toNumber(),
      'Weth not escrowed'
    )
  })

  it('should submit the correct submission cost to the inbox', async function () {
    const L1WethGateway = await ethers.getContractFactory('L1WethGateway')
    const l1WethGateway = await L1WethGateway.deploy()

    await l1WethGateway.initialize(
      l2Address,
      accounts[0].address,
      inbox.address,
      weth.address, // _l1Weth
      accounts[0].address, // _l2Weth
    )

    // send weth to bridge
    const wethAmount = 100
    await weth.deposit({ value: wethAmount })
    await weth.approve(l1WethGateway.address, wethAmount)

    let data = ethers.utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    data = ethers.utils.defaultAbiCoder.encode(
      ['address', 'bytes'],
      [accounts[0].address, data]
    )

    const tx = await l1WethGateway.outboundTransfer(
      weth.address,
      accounts[0].address,
      wethAmount,
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

    const escrowedWeth = await weth.balanceOf(l1WethGateway.address)
    assert.equal(escrowedWeth.toNumber(), 0, 'Weth should not be escrowed')
  })
})
