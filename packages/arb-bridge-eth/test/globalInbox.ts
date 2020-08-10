/*
 * Copyright 2020, Offchain Labs, Inc.
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

import { ethers, waffle } from '@nomiclabs/buidler'
import * as bre from '@nomiclabs/buidler'
import { utils, Signer } from 'ethers'
import * as chai from 'chai'
import chaiAsPromised from 'chai-as-promised'
import { deployMockContract } from '@ethereum-waffle/mock-contract'
import { GlobalInbox } from '../build/types/GlobalInbox'
import { ArbValue, Message } from 'arb-provider-ethers'

chai.use(chaiAsPromised)

const { assert, expect } = chai

async function getMessageData(
  sender: string,
  receiver: string,
  value: utils.BigNumberish
): Promise<Uint8Array> {
  const ethMsg = new Message.EthMessage(receiver, value)
  const msg = new Message.OutgoingMessage(ethMsg, sender)
  return ArbValue.marshal(msg.asValue())
}

const chainAddress = ethers.utils.getAddress(
  '0xffffffffffffffffffffffffffffffffffffffff'
)

let accounts: Signer[]
let globalInbox: GlobalInbox

const nodeHash =
  '0x10c9d77c3846591fdfc3f966935819eb7dd71ebb7e71d5d081b880868ca33e4d'
const nodeHash2 =
  '0x20c9d77c3846591fdfc3f966935819eb7dd71ebb7e71d5d081b880868ca33e4d'
const messageIndex = 0
let originalOwner: string
let address2: string
let address3: string
let address4: string

describe('GlobalInbox', async () => {
  before(async () => {
    accounts = await ethers.getSigners()

    const GlobalInbox = await ethers.getContractFactory('GlobalInbox')
    globalInbox = (await GlobalInbox.deploy()) as GlobalInbox
    await globalInbox.deployed()

    originalOwner = await accounts[0].getAddress()
    address2 = await accounts[1].getAddress()
    address3 = await accounts[2].getAddress()
    address4 = await accounts[3].getAddress()
  })

  it('should deposit eth', async () => {
    await expect(
      globalInbox.getEthBalance(chainAddress),
      'Eth balance should start at 0'
    ).to.eventually.equal(0)

    await globalInbox.depositEthMessage(
      chainAddress,
      await accounts[0].getAddress(),
      {
        value: 1000,
      }
    )

    await expect(
      globalInbox.getEthBalance(chainAddress),
      "Eth balance wasn't deposited successfully"
    ).to.eventually.equal(1000)
  })

  // These tests use a waffle mock which depends on buidlerevm
  if (bre.network.name == 'buidlerevm') {
    it('should deposit an ERC20', async () => {
      const [mockCreator] = await waffle.provider.getWallets()
      const IERC20 = await ethers.getContractFactory('IERC20')
      const mockERC20 = await deployMockContract(
        mockCreator,
        IERC20.interface.abi
      )

      await mockERC20.mock.transferFrom.returns(1)

      await expect(
        globalInbox.getERC20Balance(mockERC20.address, chainAddress),
        'ERC20 balance should start at 0'
      ).to.eventually.equal(0)

      await globalInbox.depositERC20Message(
        chainAddress,
        mockERC20.address,
        chainAddress,
        50
      )

      await expect(
        globalInbox.getERC20Balance(mockERC20.address, chainAddress),
        "ERC20 Balance wasn't deposited successfully"
      ).to.eventually.equal(50)
    })

    it('should deposit an ERC721', async () => {
      const [mockCreator] = await waffle.provider.getWallets()
      const IERC721 = await ethers.getContractFactory('IERC721')
      const mockERC721 = await deployMockContract(
        mockCreator,
        IERC721.interface.abi
      )
      await expect(
        globalInbox.hasERC721(mockERC721.address, chainAddress, 1234),
        'ERC721 Balance should start 0'
      ).to.eventually.be.false

      await mockERC721.mock.transferFrom.returns()
      await globalInbox.depositERC721Message(
        chainAddress,
        mockERC721.address,
        chainAddress,
        1234
      )

      await expect(
        globalInbox.hasERC721(mockERC721.address, chainAddress, 1234),
        "ERC721 Balance wasn't deposited successfully"
      ).to.eventually.be.true
    })
  }

  it('should make initial call', async () => {
    const data = '0x' + 'ff'.repeat(100)
    await globalInbox.sendL2Message(chainAddress, data)
  })

  it('should make second call', async () => {
    const data = '0x' + 'ff'.repeat(100)
    await globalInbox.sendL2Message(chainAddress, data)
  })

  it('should make bigger call', async () => {
    const data = '0x' + 'ff'.repeat(1000)
    await globalInbox.sendL2Message(chainAddress, data)
  })

  it('tradeable-exits: initial', async () => {
    await expect(
      globalInbox.getPaymentOwner(originalOwner, nodeHash, messageIndex),
      'current owner must be original owner'
    ).to.eventually.equal(originalOwner)

    await expect(
      globalInbox
        .connect(accounts[0])
        .transferPayment(originalOwner, address2, nodeHash, messageIndex)
    ).to.emit(globalInbox, 'PaymentTransfer')

    await expect(
      globalInbox.getPaymentOwner(originalOwner, nodeHash, messageIndex),
      'current owner must be new owner (address2)'
    ).to.eventually.equal(address2)
  })

  it('tradeable-exits: subsequent transfers', async () => {
    await expect(
      globalInbox.getPaymentOwner(originalOwner, nodeHash, messageIndex),
      'current owner must be address2'
    ).to.eventually.equal(address2)

    await expect(
      globalInbox
        .connect(accounts[0])
        .transferPayment(originalOwner, address2, nodeHash, messageIndex)
    ).to.be.revertedWith('Must be payment owner')

    await expect(
      globalInbox
        .connect(accounts[1])
        .transferPayment(originalOwner, address3, nodeHash, messageIndex)
    ).to.emit(globalInbox, 'PaymentTransfer')

    await expect(
      globalInbox.getPaymentOwner(originalOwner, nodeHash, messageIndex),
      'current owner must be new owner (address3)'
    ).to.eventually.equal(address3)

    await expect(
      globalInbox
        .connect(accounts[1])
        .transferPayment(originalOwner, address2, nodeHash, messageIndex)
    ).to.be.revertedWith('Must be payment owner.')
  })

  it('tradeable-exits: commiting transfers', async () => {
    const currOwner = await globalInbox.getPaymentOwner(
      originalOwner,
      nodeHash,
      messageIndex
    )
    await expect(
      globalInbox
        .connect(accounts[0])
        .depositEthMessage(address4, originalOwner, {
          value: 100,
        })
    ).to.emit(globalInbox, 'MessageDelivered')

    await expect(globalInbox.getEthBalance(address4)).to.eventually.equal(100)
    await expect(globalInbox.getEthBalance(currOwner)).to.eventually.equal(0)
    await expect(globalInbox.getEthBalance(originalOwner)).to.eventually.equal(
      0
    )

    const msgData = await getMessageData(originalOwner, currOwner, 50)

    await globalInbox
      .connect(accounts[3])
      .sendMessages(msgData, [1], [nodeHash])

    await expect(globalInbox.getEthBalance(address4)).to.eventually.equal(50)
    await expect(globalInbox.getEthBalance(currOwner)).to.eventually.equal(50)
    await expect(globalInbox.getEthBalance(originalOwner)).to.eventually.equal(
      0
    )

    const msgData2 = await getMessageData(address4, originalOwner, 50)

    await globalInbox
      .connect(accounts[3])
      .sendMessages(msgData2, [1], [nodeHash2])

    await expect(globalInbox.getEthBalance(address4)).to.eventually.equal(0)
    await expect(globalInbox.getEthBalance(currOwner)).to.eventually.equal(50)
    await expect(globalInbox.getEthBalance(originalOwner)).to.eventually.equal(
      50
    )
  })

  it('tradeable-exits: commiting transfers, different mnsg indexes', async () => {
    const chainAddress = await accounts[4].getAddress()
    const destAddress = await accounts[5].getAddress()
    await expect(
      globalInbox
        .connect(accounts[1])
        .depositEthMessage(chainAddress, destAddress, {
          value: 100,
        })
    ).to.emit(globalInbox, 'MessageDelivered')

    const msgData = await getMessageData(chainAddress, destAddress, 10)

    await globalInbox
      .connect(accounts[4])
      .sendMessages(msgData, [0], [nodeHash2])

    await expect(globalInbox.getEthBalance(chainAddress)).to.eventually.equal(
      100
    )
    await expect(globalInbox.getEthBalance(destAddress)).to.eventually.equal(0)

    await globalInbox
      .connect(accounts[4])
      .sendMessages(msgData, [1], [nodeHash2])

    await expect(globalInbox.getEthBalance(chainAddress)).to.eventually.equal(
      90
    )
    await expect(globalInbox.getEthBalance(destAddress)).to.eventually.equal(10)
  })
})
