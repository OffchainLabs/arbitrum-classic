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
import { ValueTester } from '../build/types/ValueTester'
import { ArbValue } from 'arb-provider-ethers'

chai.use(chaiAsPromised)

const { assert, expect } = chai

async function getMessageData(
  sender: utils.BigNumberish,
  receiver: utils.BigNumberish,
  value: utils.BigNumberish,
  valueTester: ValueTester
): Promise<Uint8Array> {
  const msgType = 1
  const msg = new ArbValue.TupleValue([
    new ArbValue.IntValue(msgType),
    new ArbValue.IntValue(ethers.utils.bigNumberify(sender).toString()),
    new ArbValue.TupleValue([
      new ArbValue.IntValue(ethers.utils.bigNumberify(receiver).toString()),
      new ArbValue.IntValue(value),
    ]),
  ])

  const msgData = ArbValue.marshal(msg)
  const res = await valueTester.deserializeMessageData(msgData, 0)

  assert.isTrue(res['0'], 'did not deserialize first part corrctly')

  const offset = res['1'].toNumber()
  assert.equal(res['2'].toNumber(), 1, 'Incorrect message type, must be ethMsg')
  assert.equal(res['3'], sender, 'Incorrect sender')

  const res2 = await valueTester.getEthMsgData(msgData, offset)
  assert.isTrue(res2['0'], "value didn't deserialize correctly")
  assert.equal(res2['2'], receiver, 'Incorrect receiver')

  assert.equal(res2['3'].toNumber(), value, 'Incorrect value sent')

  return msgData
}

function calcTxHash(
  chain: string,
  to: string,
  sequenceNum: utils.BigNumberish,
  value: utils.BigNumberish,
  messageData: string
): string {
  return ethers.utils.solidityKeccak256(
    ['address', 'address', 'uint256', 'uint256', 'bytes32'],
    [
      chain,
      to,
      sequenceNum,
      value,
      ethers.utils.solidityKeccak256(['bytes'], [messageData]),
    ]
  )
}

async function generateTxData(
  accounts: Signer[],
  chain: string,
  messageCount: utils.BigNumberish
): Promise<string> {
  const txDataTemplate = {
    to: '0xffffffffffffffffffffffffffffffffffffffff',
    sequenceNum: 2000,
    value: 54254535454544,
    messageData: '0x00',
  }

  const transactionsData = []
  for (let i = 0; i < messageCount; i++) {
    transactionsData.push(txDataTemplate)
  }

  let data = '0x'

  for (let i = 0; i < transactionsData.length; i++) {
    const txData = transactionsData[i]

    const txHash = calcTxHash(
      chain,
      txData['to'],
      txData['sequenceNum'],
      txData['value'],
      txData['messageData']
    )
    const signedTxHash = await accounts[0].signMessage(txHash)
    const packedTxData = ethers.utils.solidityPack(
      ['uint16', 'address', 'uint256', 'uint256', 'bytes', 'bytes'],
      [
        (txData['messageData'].length - 2) / 2,
        txData['to'],
        txData['sequenceNum'],
        txData['value'],
        signedTxHash,
        txData['messageData'],
      ]
    )
    data += packedTxData.slice(2)
  }
  return data
}

const chainAddress = ethers.utils.getAddress(
  '0xffffffffffffffffffffffffffffffffffffffff'
)

let accounts: Signer[]
let globalInbox: GlobalInbox
let valueTester: ValueTester

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

    const ValueTester = await ethers.getContractFactory('ValueTester')
    valueTester = (await ValueTester.deploy()) as ValueTester
    await valueTester.deployed()

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
        chainAddress,
        mockERC20.address,
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
        chainAddress,
        mockERC721.address,
        1234
      )

      await expect(
        globalInbox.hasERC721(mockERC721.address, chainAddress, 1234),
        "ERC721 Balance wasn't deposited successfully"
      ).to.eventually.be.true
    })
  }

  it('should make initial call', async () => {
    await globalInbox.sendTransactionMessage(
      chainAddress,
      chainAddress,
      2000,
      54254535454544,
      '0x'
    )
  })

  it('should make second call', async () => {
    await globalInbox.sendTransactionMessage(
      chainAddress,
      chainAddress,
      2000,
      54254535454544,
      '0x'
    )
  })

  it('should make bigger call', async () => {
    await globalInbox.sendTransactionMessage(
      chainAddress,
      chainAddress,
      2000,
      54254535454544,
      // 64 bytes of data
      '0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'
    )
  })

  it('should make a batch call', async () => {
    const messageCount = 500

    // console.log(data);

    const data = await generateTxData(accounts, chainAddress, messageCount)

    const txPromise = globalInbox.deliverTransactionBatch(chainAddress, data)

    await expect(txPromise)
      .to.emit(globalInbox, 'TransactionMessageBatchDelivered')
      .withArgs(chainAddress)

    const tx = await txPromise
    const [chainInput, txDataInput] = ethers.utils.defaultAbiCoder.decode(
      ['address', 'bytes'],
      ethers.utils.hexDataSlice(tx.data, 4)
    )

    assert.equal(
      chainInput.toLowerCase(),
      chainAddress.toLowerCase(),
      'incorrect chain from input'
    )

    assert.equal(
      txDataInput.toLowerCase(),
      data.toLowerCase(),
      'incorrect tx data from input'
    )
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
    ).to.emit(globalInbox, 'EthDepositMessageDelivered')

    await expect(globalInbox.getEthBalance(address4)).to.eventually.equal(100)
    await expect(globalInbox.getEthBalance(currOwner)).to.eventually.equal(0)
    await expect(globalInbox.getEthBalance(originalOwner)).to.eventually.equal(
      0
    )

    const msgData = await getMessageData(
      originalOwner,
      currOwner,
      50,
      valueTester
    )

    await globalInbox
      .connect(accounts[3])
      .sendMessages(msgData, [1], [nodeHash])

    await expect(globalInbox.getEthBalance(address4)).to.eventually.equal(50)
    await expect(globalInbox.getEthBalance(currOwner)).to.eventually.equal(50)
    await expect(globalInbox.getEthBalance(originalOwner)).to.eventually.equal(
      0
    )

    const msgData2 = await getMessageData(
      address4,
      originalOwner,
      50,
      valueTester
    )

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
    ).to.emit(globalInbox, 'EthDepositMessageDelivered')

    const msgData = await getMessageData(
      chainAddress,
      destAddress,
      10,
      valueTester
    )

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
